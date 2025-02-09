.DEFAULT_GOAL := all

.PHONY: run all docker_build docker_run docker_stop

all: docker_run run docker_stop

run:
	go run .

docker_build:
	docker build -t ssh-test docker/.

docker_run: docker_build
	docker run --rm --name clip-test -dp 8022:22 ssh-test

docker_stop:
	docker stop clip-test 
