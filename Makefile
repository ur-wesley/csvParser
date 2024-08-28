run:
	go run main.go

build:
	go build -o bin\upsParser.exe main.go
	copy config.yml bin\config.yml