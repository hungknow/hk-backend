-- Revert hk_backend:symbol_info from pg

BEGIN;

DROP TABLE trading.symbol_infos;
DROP TABLE trading.exchanges;
DROP TABLE trading.assets;

COMMIT;
