-- Deploy hk_backend:add_forex_candles_m5 to pg

BEGIN;

CREATE MATERIALIZED VIEW trading.forex_candles_m5
WITH (timescaledb.continuous) AS
SELECT
    time_bucket('5 min', open_time) as open_time,
    first(open, open_time) as open,
    last(close, open_time) as close,
    max(high) as high, 
    min(low) as low
FROM trading.forex_candles
group by open_time
WITH NO DATA;

COMMIT;
