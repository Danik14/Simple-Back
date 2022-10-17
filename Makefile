build:
	go build -o cmd/main main.go

run: 
	nodemon --exec "go run" ./cmd/web --watch cmd/web/main.go --watch cmd/web/handlers.go --watch internal/models/snippets.go --watch cmd/web/helpers.go --watch cmd/web/routes.go --signal SIGTERM