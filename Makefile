DOCKER_REPO?=bluebeak-test
IMAGE_NAME?=btp
IMAGE_TAG = $(shell date +"%y.%m.%d"-alpha)
BUILD_NAME = $(IMAGE_NAME)_build

default: build_deploy

build_deploy:
		make build && make deploy

# todo: better container and image naming, tagging and management
build:
		# check if dep is installed in the system
		if [ ! -f /usr/bin/dep  ] && [ ! -f ${GOPATH}/bin/dep  ]; then make dep-install; fi;
		# config and create vendor
		make dep
		# Build docker image with golang docker
		sudo docker build -t $(BUILD_NAME):$(IMAGE_TAG) -f Dockerfile.build .
		sudo docker run --name $(BUILD_NAME) -t $(BUILD_NAME):$(IMAGE_TAG) /bin/true
		sudo docker cp `sudo docker ps -q -n=1`:go/src/github.com/tauki/bluebeak-test-pe/main ./main
		ls -la
		sudo chmod 755 ./main

		# remove the build container; it's not required anymore; remove this line when it's required
		sudo docker rm $(BUILD_NAME)
		sudo docker rmi $(BUILD_NAME):$(IMAGE_TAG)


deploy:
		# Contain the binary with alpineOS
		sudo docker build -t $(DOCKER_REPO)/$(IMAGE_NAME):$(IMAGE_TAG) .
		# run the docker
		# It's set on host mode to give access to the host network, remove the tag if not needed
		sudo docker run --network="host" -d --name $(IMAGE_NAME) -p 9010:9010 $(DOCKER_REPO)/$(IMAGE_NAME):$(IMAGE_TAG)


dep:
		if [ ! -f Gopkg.toml  ]; then dep init -v; fi;
		dep ensure -v


test:
		# todo



# DEV #####


rm-rmi:
		sudo docker kill $(IMAGE_NAME)
		sudo docker rm $(IMAGE_NAME)
		sudo docker rmi $(DOCKER_REPO)/$(IMAGE_NAME):$(IMAGE_TAG)


rm-rmi-build:
		sudo docker rm $(BUILD_NAME)
		sudo docker rmi $(BUILD_NAME):$(IMAGE_TAG)


dep-dev:
		sudo rm -f Gopkg.lock
		sudo rm -f Gopkg.toml
		sudo rm -rf vendor/
		if [ ! -f /usr/bin/dep  ] && [ ! -f ${GOPATH}/bin/dep  ]; then make dep-install; fi;
		dep init -v


dep-install:
		curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh