BEGIN;

-- Insert dummy users
INSERT INTO users (name, email, number, password, role, balance)
VALUES 
('John Doe', 'john.doe@example.com', '1234567890', 'hashedpassword1', 'Admin', 1000.00),
('Jane Smith', 'jane.smith@example.com', '0987654321', 'hashedpassword2', 'Sewa', 500.00),
('Alice Johnson', 'alice.johnson@example.com', '1122334455', 'hashedpassword3', 'Penyewa', 300.00),
('Bob Brown', 'bob.brown@example.com', '5566778899', 'hashedpassword4', 'Penyewa', 200.00);

-- Insert dummy locations
INSERT INTO lokasi (user_id, address, city, state, postal_code, country, latitude, longitude)
VALUES 
(1, '123 Main St', 'Anytown', 'Anystate', '12345', 'USA', 40.712776, -74.005974),
(2, '456 Oak St', 'Othertown', 'Otherstate', '67890', 'USA', 34.052235, -118.243683),
(3, '789 Pine St', 'Sometown', 'Somestate', '11223', 'USA', 41.878113, -87.629799),
(4, '101 Maple St', 'Anycity', 'Anystate', '44556', 'USA', 37.774929, -122.419418);

-- Insert dummy products
INSERT INTO produk (owner_id, location_id, name_produk, image, license_plate, price_per_day, description)
VALUES 
(1, 1, 'Car A', 'image_url_1', 'ABC123', 50.00, 'A nice car to rent.'),
(2, 2, 'Car B', 'image_url_2', 'DEF456', 60.00, 'A comfortable car for long trips.'),
(3, 3, 'Car C', 'image_url_3', 'GHI789', 70.00, 'A stylish car for city driving.'),
(4, 4, 'Car D', 'image_url_4', 'JKL012', 80.00, 'A spacious car for family trips.');

-- Insert dummy transactions
INSERT INTO transaksi (user_id, produk_id, tipe_transaksi, jumlah)
VALUES 
(1, 1, 'topup', 100.00),
(2, 2, 'pembelian', 60.00),
(3, 3, 'pencairan', 70.00),
(4, 4, 'topup', 80.00);

-- Insert dummy rentals
INSERT INTO penyewaan (penyewa_id, produk_id, start_date, end_date, total_price, status)
VALUES 
(3, 1, '2024-06-01 10:00:00+00', '2024-06-05 10:00:00+00', 200.00, 'ongoing'),
(4, 2, '2024-06-02 10:00:00+00', '2024-06-06 10:00:00+00', 240.00, 'completed'),
(3, 3, '2024-06-03 10:00:00+00', '2024-06-07 10:00:00+00', 280.00, 'cancelled'),
(4, 4, '2024-06-04 10:00:00+00', '2024-06-08 10:00:00+00', 320.00, 'ongoing');

COMMIT;
