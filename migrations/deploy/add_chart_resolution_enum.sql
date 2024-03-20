-- Deploy hk_backend:add_chart_resolution_enum to pg

BEGIN;

CREATE TYPE chart_resolution as ENUM('S1', 'M1', 'M5', 'M15', 'M30', 'H1', 'H4', 'D1', 'W1');

COMMIT;
