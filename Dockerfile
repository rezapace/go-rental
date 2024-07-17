# Stage 1: Build Go application
FROM golang:1.22.4 AS builder
WORKDIR /build
COPY . .
COPY ./.env /build/.env
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server ./cmd/server/main.go

# Stage 2: Final image with PostgreSQL and Go application
FROM ubuntu:22.04
WORKDIR /app

# Set non-interactive frontend for apt-get
ARG DEBIAN_FRONTEND=noninteractive

# Install dependencies
RUN apt-get update && \
    apt-get install -y tzdata postgresql-14 postgresql-contrib-14 supervisor curl && \
    ln -fs /usr/share/zoneinfo/Etc/UTC /etc/localtime && \
    dpkg-reconfigure --frontend noninteractive tzdata && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

# Ensure PostgreSQL binaries are in PATH
ENV PATH="/usr/lib/postgresql/14/bin:$PATH"

# Set environment variables for PostgreSQL
ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=mysecretpassword
ENV POSTGRES_DB=rental
ENV PGDATA=/var/lib/postgresql/data

# Setup PostgreSQL data directory
RUN mkdir -p /var/lib/postgresql/data && \
    chown -R postgres:postgres /var/lib/postgresql && \
    chmod 700 /var/lib/postgresql/data

# Copy Go application from builder stage
COPY --from=builder /build/server ./server
COPY --from=builder /build/.env /app/.env

# Copy initialization scripts and supervisord configuration
COPY ./initdb /docker-entrypoint-initdb.d/
COPY ./cmd/server/supervisord.conf /etc/supervisor/conf.d/supervisord.conf
COPY ./init_postgres.sh /app/init_postgres.sh

# Ensure init_postgres.sh is executable
RUN chmod +x /app/init_postgres.sh

# Create necessary directories and set permissions
RUN mkdir -p /var/log/supervisor /tmp /var/run && \
    chown -R postgres:postgres /var/log/supervisor /tmp /var/run /etc/postgresql /docker-entrypoint-initdb.d /app/init_postgres.sh

# Expose ports
EXPOSE 5432 8080

# Start supervisord
CMD ["/usr/bin/supervisord", "-n", "-c", "/etc/supervisor/conf.d/supervisord.conf"]
