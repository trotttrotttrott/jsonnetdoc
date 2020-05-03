TAG=0.0.1

test:
	go test

docker-build:
	docker build -t jsonnetdoc:$(TAG) .
