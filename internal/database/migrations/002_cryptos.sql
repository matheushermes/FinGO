CREATE TABLE IF NOT EXISTS cryptos (
    id SERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    name VARCHAR(100) NOT NULL,                  -- Nome da cripto (ex: Bitcoin)
    symbol VARCHAR(10) NOT NULL,                 -- Símbolo (ex: BTC)
    amount NUMERIC(20, 8) NOT NULL,              -- Quantidade comprada
    purchase_price_usd NUMERIC(20, 2) NOT NULL,  -- Valor total gasto na compra (em dólares)
    purchase_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_sold BOOLEAN DEFAULT FALSE,                -- Marcar se o ativo já foi vendido
    notes TEXT,                                   -- Campo opcional para observações
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);