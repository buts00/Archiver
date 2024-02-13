.PHONY: build
build:
	rm -rf build && mkdir build &&  go build -o build/archiver -v ./cmd/app/

.PHONY: pack_sf
pack_sf:
	build/archiver pack -m shannon-fano example.txt

.PHONY: unpack_sf
unpack_sf:
	build/archiver unpack -m shannon-fano example.sf

.PHONY: run
run:
	go run cmd/app/main.go

.DEFAULT_GOAL:=build
