build:
	go build -o cmd/main main.go

run: 
	nodemon --exec "go run" main.go --signal SIGTERM