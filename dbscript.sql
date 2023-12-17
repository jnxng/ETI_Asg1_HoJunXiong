CREATE USER 'jnxng'@'localhost' IDENTIFIED BY
'asg1';
GRANT ALL ON *.* TO 'jnxng'@'localhost';

CREATE DATABASE IF NOT EXISTS carpooling;

USE carpooling;

CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    mobile VARCHAR(20) NOT NULL,
    email VARCHAR(255) NOT NULL,
    driver_license VARCHAR(20),
    car_plate VARCHAR(20),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS trips (
    id INT AUTO_INCREMENT PRIMARY KEY,
    car_owner_id INT NOT NULL,
    pickup_location VARCHAR(255) NOT NULL,
    alt_pickup_location VARCHAR(255),
    start_time DATETIME NOT NULL,
    destination VARCHAR(255) NOT NULL,
    available_seats INT NOT NULL,
    CONSTRAINT fk_car_owner
        FOREIGN KEY (car_owner_id) REFERENCES users(id)
);

INSERT INTO users (first_name, last_name, mobile, email) VALUES
    ('Wriothesley', 'Chan', '1234567890', 'Wrio@example.com');

INSERT INTO trips (car_owner_id, pickup_location, start_time, destination, available_seats) VALUES
    (1, 'Office', '2023-12-20 08:00:00', 'Home', 3);
