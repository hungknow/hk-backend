-- Deploy hk_backend:add_forex_candles to pg
BEGIN;

CREATE TABLE IF NOT EXISTS trading.forex_candles (
    symbol_id SERIAL NOT NULL references trading.symbol_infos(id),
    resolution chart_resolution NOT NULL,
    open_time TIMESTAMPTZ NOT NULL,
    open DECIMAL(20, 6) NOT NULL,
    close DECIMAL(20, 6) NOT NULL,
    high DECIMAL(20, 6) NOT NULL,
    low DECIMAL(20, 6) NOT NULL,
    volume DECIMAL(20, 6) NULL,
    constraint forex_candle_pk primary key (symbol_id, resolution, open_time)
);

SELECT create_hypertable('trading.forex_candles', by_range('open_time'));

CREATE INDEX ix_forex_candles_symbol_resolution_time ON trading.forex_candles (symbol_id, resolution, open_time DESC);

COMMIT;