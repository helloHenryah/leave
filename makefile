build:
	go build -o ./bin/leave .

run:
	go run .

linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/leave-linux-amd64 --ldflags="-s -w" .

windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./bin/leave-windows-amd64.exe --ldflags="-s -w" .

pull:
	git pull

push:
	git add .
	git commit -m "update submit button"
	git push origin main
	git push github main