-- Deploy hk_backend:symbol_info to pg

BEGIN;

CREATE TABLE IF NOT EXISTS trading.assets {
    id SERIAL PRIMARY KEY,
    symbol TEXT NOT NULL UNIQUE,
    name TEXT NOT NULL,
};

CREATE TABLE IF NOT EXISTS trading.exchanges {
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
};

CREATE TABLE IF NOT EXISTS trading.symbol_infos (
    id SERIAL PRIMARY KEY,
    symbol TEXT NOT NULL UNIQUE,
    name TEXT NOT NULL,
    base_currency TEXT NOT NULL references trading.assets(symbol),
    quote_currency TEXT NOT NULL references trading.assets(symbol),
    exchange_id INTEGER NOT NULL references trading.exchanges(id),
);

INSERT INTO trading.assets(symbol, name)
    ("XAU", "Gold"),
    ("USD", "US Dollar");

INSERT INTO trading.exchanges(id, name)
    (1, "MOCK");

INSERT INTO trading.symbol_infos(symbol, name, base_currency, quote_currency, exchange_id)
    ("MOCK:XAUUSD", "XAU/USD", "XAU", "USD", 1);

COMMIT;
