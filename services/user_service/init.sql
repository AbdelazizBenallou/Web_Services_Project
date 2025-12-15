CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    email TEXT NOT NULL UNIQUE,
    full_name TEXT NOT NULL,
    password TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS profiles (
    user_id BIGINT PRIMARY KEY,
    first_name TEXT,
    last_name TEXT,
    birth_date DATE,
    address TEXT
);

CREATE TABLE IF NOT EXISTS roles (
    user_id BIGINT,
    role TEXT NOT NULL,
    UNIQUE (user_id, role)
);

