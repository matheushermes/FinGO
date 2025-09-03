CREATE TABLE IF NOT EXISTS price_alerts (
    id SERIAL PRIMARY KEY,
    user_email VARCHAR(255) NOT NULL,
    crypto_id INT NOT NULL REFERENCES cryptos(id) ON DELETE CASCADE,
    symbol VARCHAR(20) NOT NULL,
    percentage_change NUMERIC(5,2) NOT NULL,
    direction VARCHAR(10) NOT NULL CHECK (direction IN ('above', 'below')),
    target_price NUMERIC(20,8) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    triggered BOOLEAN NOT NULL DEFAULT FALSE
);
