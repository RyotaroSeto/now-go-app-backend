api:
	docker-compose exec nowapi ash

db:
	docker-compose exec nowdb ash

wire:
	wire ./pkg/injection/wire.go
	go generate ./...

createdb:
	docker exec -it nowdb createdb --username=postgres --owner=postgres nowdb

swag:
	swag init

lint:
	golangci-lint run

.PHONY: api db wire createdb swag lint
