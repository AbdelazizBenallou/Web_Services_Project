-- 001_init.up.sql
CREATE TABLE IF NOT EXISTS users (
    user_id       BIGSERIAL    PRIMARY KEY,
    full_name     TEXT         NOT NULL,
    email         TEXT         NOT NULL UNIQUE,
    password_hash TEXT         NOT NULL,
    created_at    TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

-- Optional: Add index on email for faster lookups
CREATE INDEX IF NOT EXISTS idx_users_email ON users (email);
