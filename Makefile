build:
	go build -ldflags="-X main.VERSION=DEV -X main.COMMIT=$(shell git rev-parse HEAD)" -o "metaprint-DEV-$(shell uname | tr '[:upper:]' '[:lower:]')-$(shell uname -m)" .

install: build
	sudo mv metaprint /usr/bin
