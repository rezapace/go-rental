   BEGIN;

    -- Use ENUM types for predefined lists
    CREATE TYPE role_type AS ENUM ('Admin', 'Sewa', 'Penyewa');
    CREATE TYPE transaksi_type AS ENUM ('topup', 'pencairan', 'pembelian');
    CREATE TYPE status_type AS ENUM ('ongoing', 'completed', 'cancelled');

    -- Table for storing user details
    CREATE TABLE users (
        id SERIAL PRIMARY KEY,
        name VARCHAR(100) NOT NULL,
        email VARCHAR(100) NOT NULL UNIQUE,
        number VARCHAR(20) NOT NULL UNIQUE,
        password TEXT NOT NULL, -- Store hashed passwords
        role role_type NOT NULL,
        balance DECIMAL(10, 2) DEFAULT 0.00,
        created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at TIMESTAMP WITH TIME ZONE NULL
    );

    -- Table for storing locations associated with users
    CREATE TABLE lokasi (
        id SERIAL PRIMARY KEY,
        user_id INT NOT NULL,
        address VARCHAR(255) NOT NULL,
        city VARCHAR(100) NOT NULL,
        state VARCHAR(100) NOT NULL,
        postal_code VARCHAR(20) NOT NULL,
        country VARCHAR(100) NOT NULL,
        latitude DECIMAL(9, 6),
        longitude DECIMAL(9, 6),
        created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
        FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
    );

    -- Indexes for optimizing search operations
    CREATE INDEX idx_users_email ON users(email);
    CREATE INDEX idx_lokasi_city ON lokasi(city);
    CREATE INDEX idx_lokasi_state ON lokasi(state);

    -- Table for storing products
    CREATE TABLE produk (
        id SERIAL PRIMARY KEY,
        owner_id INT NOT NULL,
        location_id INT NOT NULL,
        name_produk VARCHAR(100) NOT NULL,
        image TEXT,
        license_plate VARCHAR(50) NOT NULL UNIQUE,
        price_per_day DECIMAL(10, 2) NOT NULL,
        description TEXT,
        created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at TIMESTAMP WITH TIME ZONE NULL,
        FOREIGN KEY (owner_id) REFERENCES users(id) ON DELETE CASCADE,
        FOREIGN KEY (location_id) REFERENCES lokasi(id) ON DELETE CASCADE
    );

    -- Table for transactions
    CREATE TABLE transaksi (
        id SERIAL PRIMARY KEY,
        user_id INT NOT NULL,
        produk_id INT,
        tipe_transaksi transaksi_type NOT NULL,
        jumlah DECIMAL(10, 2) NOT NULL,
        created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
        FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
        FOREIGN KEY (produk_id) REFERENCES produk(id) ON DELETE SET NULL
    );

    -- Table for rentals
    CREATE TABLE penyewaan (
        id SERIAL PRIMARY KEY,
        penyewa_id INT NOT NULL,
        produk_id INT NOT NULL,
        start_date TIMESTAMP WITH TIME ZONE NOT NULL,
        end_date TIMESTAMP WITH TIME ZONE NOT NULL,
        total_price DECIMAL(10, 2) NOT NULL,
        status status_type NOT NULL,
        created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
        FOREIGN KEY (penyewa_id) REFERENCES users(id) ON DELETE CASCADE,
        FOREIGN KEY (produk_id) REFERENCES produk(id) ON DELETE CASCADE
    );


   COMMIT;

