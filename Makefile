SHELL := /bin/bash

build:
	docker build . \
		-t go-notes-service

up:
	make build && \
	docker run \
  	--name go-notes-service \
		-p 3000:3000 \
		go-notes-service

down:
	docker stop go-notes-service

clean:
	docker rm -f go-notes-service
