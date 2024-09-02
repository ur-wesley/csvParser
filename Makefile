run:
	go run main.go

build:
	go build -o bin\csvParser.exe main.go
	copy config_exmaple.yml bin\config_exmaple.yml