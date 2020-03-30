SHELL := /bin/bash

build:
	docker build . \
		-t go-notes-service

up:
	make build && \
	docker run \
  	--name go-notes-service \
		-p 3000:3000 \
		-d \
		go-notes-service && \
	make log

log:
	docker logs -f go-notes-service

down:
	docker stop go-notes-service

clean:
	make down && \
	docker rm -f go-notes-service
