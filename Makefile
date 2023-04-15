wire:
	wire ./pkg/injection/wire.go
	go generate ./...

createdb:
	docker exec -it nowdb createdb --username=postgres --owner=postgres nowdb
