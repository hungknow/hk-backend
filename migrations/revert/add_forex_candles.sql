-- Revert hk_backend:add_forex_candles from pg

BEGIN;

DROP INDEX ix_forex_candles_symbol_resolution_time;
DROP TABLE trading.forex_candles;

COMMIT;
