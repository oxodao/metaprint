build:
	go build

install: build
	sudo mv metaprint /usr/bin