.PHONY: build install clean

build:
	go build -o bin/cmg ./cmd/cmg

install:
	go install ./cmd/cmg

clean:
	rm -rf bin/