-- Revert hk_backend:add_trading_schema from pg

BEGIN;

DROP SCHEMA trading;

COMMIT;
