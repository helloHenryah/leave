build:
	go build -o /bin/leave main.go

run:
	go build -o /bin/leave main.go && ./bin/leave -port 8080