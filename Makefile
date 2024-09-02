setup:
	go mod tidy

run:
	go run .

build:
	go build -o bin\csvParser.exe main.go
	copy config_exmaple.yml bin\config_exmaple.yml

clear:
	git clean -Xdf
