.PHONY: build
# build
build:
	GOOS=darwin GOARCH=amd64  CGO_ENABLED=0
	mkdir -p bin/
	go build -ldflags "-X main.Version=v0.1" -o ./bin/gold main.go

.PHONY: build_linux
# build_linux
build_linux:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "-X main.Version=v0.1" -o ./bin_linux/gold main.go