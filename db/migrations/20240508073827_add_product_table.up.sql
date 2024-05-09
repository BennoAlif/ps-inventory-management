CREATE TABLE products (
  id SERIAL PRIMARY KEY,
  name VARCHAR(30) NOT NULL CHECK (LENGTH(name) >= 1),
  sku VARCHAR(30) NOT NULL CHECK (LENGTH(sku) >= 1),
  category VARCHAR(255) NOT NULL CHECK (category IN ('Clothing', 'Accessories', 'Footwear', 'Beverages')),
  image_url VARCHAR(255) NOT NULL,
  notes TEXT NOT NULL CHECK (LENGTH(notes) >= 1 AND LENGTH(notes) <= 200),
  price INTEGER NOT NULL CHECK (price >= 1),
  stock INTEGER NOT NULL CHECK (stock >= 0 AND stock <= 100000),
  location VARCHAR(200) NOT NULL CHECK (LENGTH(location) >= 1),
  is_available BOOLEAN NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);