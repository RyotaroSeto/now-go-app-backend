-- +goose NO TRANSACTION

-- +goose Up

CREATE SCHEMA IF NOT EXISTS nowgo;

-- +goose Down

DROP SCHEMA IF EXISTS nowgo;
