
all: build

build:
	go build -buildmode=pie -buildmode=c-shared

fmt:
	goimports -w -l .
	gofumpt -w -l .

gen:
	go generate ./...

