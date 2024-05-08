CREATE TABLE products (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    sku VARCHAR(255) NOT NULL,
    category VARCHAR(255) NOT NULL,
    image_url VARCHAR(255),
    stock INTEGER NOT NULL,
    notes TEXT,
    price DECIMAL(10,2) NOT NULL,
    location VARCHAR(255),
    is_available BOOLEAN NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
