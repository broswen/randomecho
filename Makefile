.PHONY: build test
DOCKER_REPO=broswen/randomecho

up:
	docker-compose up --build

down:
	docker-compose down

docker-build:
	docker build . -t $(DOCKER_REPO):$(tag)

docker-push:
	docker push $(DOCKER_REPO):$(tag)

test:
	go test -race ./...
