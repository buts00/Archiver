.PHONY: build
build:
	rm -rf build && mkdir build &&  go build -o build/archiver -v ./cmd/app/

.PHONY: run archiver
run archiver:
	build/archiver pack vlc cmd/app/main.go

.PHONY: run
run:
	go run cmd/app/main.go

.DEFAULT_GOAL:=build
