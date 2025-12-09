-- Create orders table if not exists
CREATE TABLE IF NOT EXISTS orders (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    product_id INT NOT NULL,
    quantity INT NOT NULL CHECK (quantity > 0),
    total_price DECIMAL(10, 2) NOT NULL CHECK (total_price >= 0),
    status VARCHAR(50) DEFAULT 'pending' CHECK (status IN ('pending', 'processing', 'completed', 'cancelled')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes if not exist (PostgreSQL 9.5+)
CREATE INDEX IF NOT EXISTS idx_orders_user_id ON orders(user_id);
CREATE INDEX IF NOT EXISTS idx_orders_product_id ON orders(product_id);
CREATE INDEX IF NOT EXISTS idx_orders_status ON orders(status);

-- Insert sample data for testing
INSERT INTO orders (user_id, product_id, quantity, total_price, status) VALUES
(1, 1, 2, 299.98, 'completed'),
(1, 2, 1, 89.99, 'pending'),
(2, 3, 3, 179.97, 'processing'),
(3, 1, 1, 149.99, 'completed')
ON CONFLICT DO NOTHING;  -- لتجنب تكرار البيانات إذا تم تنفيذ السكريبت أكثر من مرة
