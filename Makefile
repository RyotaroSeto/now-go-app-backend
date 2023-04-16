DB_URL=postgresql://nowdb:postgres@postgres:5432/nowdb?sslmode=disable

api:
	docker-compose exec nowapi ash

db:
	docker-compose exec nowdb ash

wire:
	wire ./pkg/injection/wire.go
	go generate ./...

createdb:
	docker exec -it nowdb createdb --username=postgres --owner=postgres nowdb

gooseup:
	goose postgres "user=postgres dbname=nowdb password=postgres sslmode=disable" up
	goose postgres "user=postgres dbname=nowdb sslmode=disable" status
	goose postgres -database "$(DB_URL)" up

