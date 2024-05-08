CREATE TABLE customers (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    phone_number VARCHAR(20) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL
     created_at timestamp
);



CREATE TABLE drivers (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    phone_number VARCHAR(20) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    license_number VARCHAR(20) UNIQUE NOT NULL
     created_at timestamp
);

CREATE TABLE cars (
    id SERIAL PRIMARY KEY,
    license_plate VARCHAR(20) UNIQUE NOT NULL,
    model VARCHAR(50) NOT NULL,
    year INT NOT NULL,
    driver_id INT REFERENCES drivers(id)
     created_at timestamp
);

CREATE TABLE addresses (
    id SERIAL PRIMARY KEY,
    street VARCHAR(100) NOT NULL,
    city VARCHAR(50) NOT NULL,
    state VARCHAR(50),
    country VARCHAR(50) NOT NULL,
    zip_code VARCHAR(20) NOT NULL
     created_at timestamp
);

CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    customer_id INT REFERENCES customers(id),
    driver_id INT REFERENCES drivers(id),
    car_id INT REFERENCES cars(id),
    order_time TIMESTAMP NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'pending'
);

CREATE TABLE order_addresses (
    order_id INT REFERENCES orders(id),
    address_id INT REFERENCES addresses(id),
    PRIMARY KEY (order_id, address_id)
);

INSERT INTO order_addresses(order_id, address_id)
VALUES(1,1);

INSERT INTO orders(customer_id, driver_id, car_id, order_time)
VALUES('1', '1', '4', now());

INSERT INTO addresses(street, city, state, country, zip_code, created_at)
VALUES('TestStreet', 'TestCity', 'TestState', 'TestCountry', '00000', now());

INSERT INTO cars(license_plate, model, year, driver_id, created_at)
VALUES('HA0980FH', 'audi', '2017', 1, now());

INSERT INTO drivers(first_name, last_name, phone_number, email, license_number, created_at)
VALUES('testFirstName', 'testsecondName', '+38050405204204', 'test2@gmail.com', '1234567890', now();

INSERT INTO customers(first_name, last_name, phone_number, email, created_at)
VALUES('testFirstName', 'testsecondName', '+380403053043', 'test"gmail.com', now();

