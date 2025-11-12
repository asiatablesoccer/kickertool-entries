build:
	mkdir -p functions
	go get ./...
	GOOS=linux GOARCH=amd64 go build -o functions/tournamentEntries main.go
