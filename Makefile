autoserver:
		nodemon --watch './**/*.go' --signal SIGTERM --exec 'go' run cmd/main.go

swaggo:
	swag init -g ./cmd/main.go