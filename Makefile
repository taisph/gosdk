.PHONY: all build install clean

all: build

build:
	go build -o out/gosdk ./cmd/gosdk

install:
	go install ./cmd/gosdk

clean:
	rm -r out
	go clean -i ./cmd/gosdk
