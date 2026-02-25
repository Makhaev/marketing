-- USERS
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    phone VARCHAR(20) UNIQUE NOT NULL,
    role VARCHAR(20) NOT NULL DEFAULT 'user',
    created_at TIMESTAMP DEFAULT NOW()
);

-- STORES
CREATE TABLE stores (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    address TEXT,
    owner_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT NOW()
);

-- PRODUCTS
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    category VARCHAR(100)
);

-- STORE_PRODUCTS
CREATE TABLE store_products (
    id SERIAL PRIMARY KEY,
    store_id INTEGER REFERENCES stores(id) ON DELETE CASCADE,
    product_id INTEGER REFERENCES products(id) ON DELETE CASCADE,
    price NUMERIC(10,2) NOT NULL,
    is_promo BOOLEAN DEFAULT FALSE,
    updated_at TIMESTAMP DEFAULT NOW(),
    image_url TEXT
);

-- индекс на цену
CREATE INDEX idx_product_price 
ON store_products(product_id, price);

-- уникальное ограничение
ALTER TABLE store_products
ADD CONSTRAINT unique_store_product UNIQUE (store_id, product_id);