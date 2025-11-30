#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "product" --dbname "productdb" <<-EOSQL
    CREATE TABLE IF NOT EXISTS products (
        id VARCHAR(36) PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        description TEXT,
        price DECIMAL(10,2) NOT NULL,
        stock INTEGER NOT NULL DEFAULT 0,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );
    
    INSERT INTO products (id, name, description, price, stock) VALUES 
    ('prod-1', 'Laptop', 'Gaming Laptop', 999.99, 10),
    ('prod-2', 'Mouse', 'Wireless Mouse', 29.99, 50)
    ON CONFLICT (id) DO NOTHING;
EOSQL
