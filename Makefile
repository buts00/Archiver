.PHONY: build
build:
	rm -rf build && mkdir build &&  go build -o build/archiver -v ./cmd/app/

.PHONY: run archiver
run archiver:
	build/archiver pack vlc example.txt

.PHONY: run
run:
	go run cmd/app/main.go

.DEFAULT_GOAL:=build
