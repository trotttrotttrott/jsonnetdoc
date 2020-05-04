IMAGE=trotttrotttrott/jsonnetdoc
TAG=$$(git rev-parse --short HEAD)

test:
	go test

docker-build:
	docker build -t ${IMAGE}:${TAG} .
	docker tag ${IMAGE}:${TAG} ${IMAGE}:latest

docker-push:
	docker push ${IMAGE}:${TAG}
	docker push ${IMAGE}:latest
