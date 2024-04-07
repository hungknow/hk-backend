-- Deploy hk_backend:add_evm_compatible_transactions to pg

BEGIN;

CREATE TABLE IF NOT EXISTS blockchain.evm_transactions (
    hash BYTEA PRIMARY KEY,
    chain_id INTEGER NOT NULL,
    block_number INTEGER NOT NULL,
    block_hash BYTEA NOT NULL,
    from_address_hash BYTEA NOT NULL,
    to_address_hash BYTEA NOT NULL,
    gas_price NUMERIC NOT NULL,
    gas_used NUMERIC NOT NULL,
) PARTITION BY (chain_id);

CREATE TABLE IF NOT EXISTS blockchain.evm_smart_contracts(
    id BIGSERIAL PRIMARY KEY,
    chain_id INTEGER NOT NULL,
    name VARCHAR(255) NOT NULL,
    compiler_version VARCHAR(255) NOT NULL,
    address_hash BYTEA NOT NULL,
    inserted_at TIMESTAMPTZ NOT NULL,
    constructor_arguments TEXT NOT NULL,
) PARTITION BY (chain_id);

CREATE TABLE IF NOT EXISTS blockchain.evm_tokens(
    contract_address_hash bytea PRIMARY KEY,
    chain_id INTEGER NOT NULL,
    name TEXT NOT NULL,
    symbol TEXT NOT NULL,
    decimals INTEGER NOT NULL,
    inserted_at TIMESTAMPTZ NOT NULL,
    decimals NUMERIC NOT NULL,
    total_supply NUMERIC NOT NULL,
    type VARCHAR(255) NOT NULL
) PARTITION BY (chain_id);

COMMIT;
