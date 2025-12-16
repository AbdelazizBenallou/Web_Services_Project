CREATE TABLE IF NOT EXISTS categories (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS products (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    category_id BIGINT NOT NULL,
    price NUMERIC(10,2) NOT NULL CHECK (price > 0),
    CONSTRAINT fk_category FOREIGN KEY (category_id)
        REFERENCES categories(id)
);

CREATE TABLE IF NOT EXISTS stock (
    product_id BIGINT PRIMARY KEY,
    quantity INT NOT NULL CHECK (quantity >= 0),
    CONSTRAINT fk_product FOREIGN KEY (product_id)
        REFERENCES products(id)
        ON DELETE CASCADE
);

