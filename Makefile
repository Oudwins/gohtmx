BIN_NAME=main
.DEFAULT_GOAL := run

build:
	go build -o ./tmp/${BIN_NAME} ./cmd/server/*.go

run: build
	./tmp/${BIN_NAME}	

db clean: 
	go run ./cmd/cleanDB/cleanMain.go

workbench:
	go run ./cmd/workbench/main.go