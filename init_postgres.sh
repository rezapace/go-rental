#!/bin/bash
set -e

# Initialize PostgreSQL data directory if it's empty
if [ -z "$(ls -A "$PGDATA")" ]; then
    echo "Initializing PostgreSQL data directory..."
    initdb -D "$PGDATA"
    
    # Modify postgresql.conf to allow connections from all addresses
    echo "listen_addresses='*'" >> "$PGDATA/postgresql.conf"
    
    # Modify pg_hba.conf to allow connections from all addresses
    echo "host all all 0.0.0.0/0 md5" >> "$PGDATA/pg_hba.conf"
fi

# Start PostgreSQL
pg_ctl -D "$PGDATA" -o "-c config_file=/etc/postgresql/postgresql.conf" start

# Wait for PostgreSQL to start
pg_ctl -D "$PGDATA" -w start

echo "PostgreSQL started successfully"

# Create user and database if they don't exist
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" <<-EOSQL
    DO
    \$do\$
    BEGIN
        IF NOT EXISTS (SELECT FROM pg_catalog.pg_roles WHERE rolname = '$POSTGRES_USER') THEN
            CREATE USER $POSTGRES_USER WITH PASSWORD '$POSTGRES_PASSWORD';
        END IF;
    END
    \$do\$;

    SELECT 'CREATE DATABASE $POSTGRES_DB'
    WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = '$POSTGRES_DB')\gexec

    GRANT ALL PRIVILEGES ON DATABASE $POSTGRES_DB TO $POSTGRES_USER;
EOSQL

echo "Running initialization scripts..."
for f in /docker-entrypoint-initdb.d/*; do
    case "$f" in
        *.sql)    echo "$0: running $f"; psql -U "$POSTGRES_USER" -d "$POSTGRES_DB" -f "$f"; echo ;;
        *.sql.gz) echo "$0: running $f"; gunzip -c "$f" | psql -U "$POSTGRES_USER" -d "$POSTGRES_DB"; echo ;;
        *)        echo "$0: ignoring $f" ;;
    esac
done

echo "Database initialization completed."

# Keep PostgreSQL running
tail -f /dev/null