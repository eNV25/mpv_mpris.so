
all: build

gen:
	go generate ./...

build:
	go build -buildmode=pie -buildmode=c-shared
