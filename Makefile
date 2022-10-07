build:
	go build -o cmd/main main.go

run: 
	nodemon --exec "go run" ./cmd/web --watch cmd/web/main.go --watch cmd/web/handlers.go --signal SIGTERM