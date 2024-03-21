-- Revert hk_backend:add_chart_resolution_enum from pg

BEGIN;

DROP TYPE IF EXISTS chart_resolution;

COMMIT;
