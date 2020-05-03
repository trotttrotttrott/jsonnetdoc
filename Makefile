test:
	go test

docker-build:
	docker build -t trotttrotttrott/jsonnetdoc:$$(git rev-parse --short HEAD) .

docker-push:
	docker push trotttrotttrott/jsonnetdoc:$$(git rev-parse --short HEAD)
