SHELL := /bin/bash

build:
	docker build . \
		-t go-notes-service

run:
	make build && \
	docker run \
  	--name go-notes-service \
		-p 3000:3000 \
		-d \
		go-notes-service && \
	make logs

logs:
	docker logs -f go-notes-service

stop:
	docker stop go-notes-service

clean:
	make stop && \
	docker rm -f go-notes-service
