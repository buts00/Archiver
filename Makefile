.PHONY: build
build:
	rm -rf build && mkdir build &&  go build -o build/archiver -v ./cmd/app/

.PHONY: run pack vlc
run pack vlc:
	build/archiver pack -m vlc example.txt

.PHONY: run unpack vlc
run unpack vlc:
	build/archiver unpack -m vlc example.vlc

.PHONY: run
run:
	go run cmd/app/main.go

.DEFAULT_GOAL:=build
