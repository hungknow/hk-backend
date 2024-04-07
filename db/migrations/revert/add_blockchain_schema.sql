-- Revert hk_backend:add_blockchain_schema from pg

BEGIN;

DROP SCHEMA blockchain;

COMMIT;
