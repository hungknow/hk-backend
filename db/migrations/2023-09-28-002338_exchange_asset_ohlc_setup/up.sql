-- Create table to store the list of provider ---
CREATE SCHEMA economy;

CREATE TABLE IF NOT EXISTS "economy".exchange_info(
    id TEXT PRIMARY KEY,
    "name" TEXT
);

-- Table to store the symbol for OHLC -- 
CREATE TABLE IF NOT EXISTS "economy".asset_info(
    id int PRIMARY KEY,
    symbol_name TEXT,
    symbol_description TEXT,
    symbol_type TEXT,
    exchange_id TEXT,
    country TEXT,
    pip_size decimal,
    timezone int8
);

-- Table to store the ohlc value per symbol --

CREATE TABLE IF NOT EXISTS "economy".ohlc_data(
    asset_info_id TEXT,
    created_time timestamp with time zone,
    timeframe int,
    open decimal,
    high decimal,
    low decimal,
    close decimal,
    volume decimal,
    primary key (asset_info_id, created_time)
);

INSERT INTO "economy".exchange_info(id, name)
    VALUES('HD', 'Historical Data');

INSERT INTO "economy".asset_info(id, symbol_name, symbol_description, symbol_type, exchange_id, country, pip_size, timezone)
    VALUES(1, 'HD:XAUUSD', 'Gold/U.S. Dollar for testing', 'commodity_cfd', 'HD', '', 0.01, -4);
