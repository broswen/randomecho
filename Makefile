.PHONY: build test
DOCKER_REPO=broswen/randomecho

up:
	docker-compose up --build

down:
	docker-compose down

build:
	docker build . -t $(DOCKER_REPO):$(tag)

test:
	go test -race ./...
