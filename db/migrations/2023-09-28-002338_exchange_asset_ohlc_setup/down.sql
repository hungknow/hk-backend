-- This file should undo anything in `up.sql`


DROP TABLE  "economy".ohlc_data;
DROP TABLE "economy".asset_info;
DROP TABLE "economy".exchange_info;

DROP SCHEMA "economy";