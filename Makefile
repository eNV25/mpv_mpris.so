
all: build

build:
	go build -buildmode=pie -buildmode=c-shared

fmt:
	goimports -local github.com/env25/mpv_mpris.so -w -l .
	gofumpt -w -l .

gen:
	go generate ./...

