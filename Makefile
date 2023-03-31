.PHONY: build

build:
	@ go build -o ./build/server ./cmd/server

run: .PHONY build
	@ ./build/server
