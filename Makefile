wire-authentication:
	wire ./src/authentication/injection/wire.go
	go generate ./...

wire-user:
	wire ./src/user/injection/wire.go
	go generate ./...
