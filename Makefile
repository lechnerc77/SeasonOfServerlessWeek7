RELEASE=0.0.2
APP=recipe-connector
DOCKER_ACCOUNT=lechnerc77
CONTAINER_IMAGE=${DOCKER_ACCOUNT}/${APP}:${RELEASE}

.PHONY: build-image push-image

build-image:
	docker build -t $(CONTAINER_IMAGE) --no-cache --rm .

push-image: build-image
	docker push $(CONTAINER_IMAGE)

docker-run:
	docker run -it -p 8080:80 $(CONTAINER_IMAGE)