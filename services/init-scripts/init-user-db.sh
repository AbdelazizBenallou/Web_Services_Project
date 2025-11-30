#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "amine" --dbname "userdb" <<-EOSQL
    CREATE TABLE IF NOT EXISTS users (
        id BIGSERIAL PRIMARY KEY,
        full_name VARCHAR(255) NOT NULL,
        email VARCHAR(255) UNIQUE NOT NULL,
        password_hash VARCHAR(255) NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );
    
    INSERT INTO users (name, email) VALUES 
    ('Amine', 'amine@example.com', '123456780'),
    ('John', 'john@example.com', '12345678')
    ON CONFLICT (email) DO NOTHING;
EOSQL
