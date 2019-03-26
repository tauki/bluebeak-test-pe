# bluebeak-test-pe

There are a number of ways to run the program.

To run the program by building from scrach type
```
go build -o "main"
```
This will create an executable file named main, change the quoted name to change it to anything else.

The executable file is a cli app, to see more into it, type
```
./main --help
```

To run the server type
```
./main run
```
It will spin-up the server and start listening and serving the APIs. Defaultly the server is set to a dev environment so the logs are expected to be printed after the server is spun-up.

To run the individual scripts without spinning-up the server, just type `./main --help` or `./main -h` or just `./main` and it will printout the flags and arguments required to run individual script.

Just running `./main` will printout the basic information of the app which incudes all necessary flags and commands to run the app and individual components.

There are two files included for dockerinzing the app. The `Dockerfile` dockerizes the file but it requires the program to be already built. Defaultly the `Dockerfile` expects the executable to be named `main`, to change that open the `Dockerfile` and change the following line
```
ENTRYPOINT /go/src/github.com/tauki/bluebeak-test-pe/main run
```
change the word main with the name of your executable.

The other Dockerfile or `Dockerfile.build` is meant for building the app and also meant to be run before building the `Dockerfile`.

To build with the `Dockerfile` and `Dockerfile.build` the following sequence can be considered.

```
docker build -t $(BUILD_NAME):$(IMAGE_TAG) -f Dockerfile.build .
docker run --name $(BUILD_NAME) -t $(BUILD_NAME):$(IMAGE_TAG) /bin/true
docker cp `sudo docker ps -q -n=1`:go/src/github.com/tauki/bluebeak-test-pe/main ./main
chmod 755 ./main
```
When this is done you might consider removing the image and container by typing the following lines if required
```
docker rm $(BUILD_NAME)
docker rmi $(BUILD_NAME):$(IMAGE_TAG)
```
after they are done to build the container and run the app, the following lines can be considered
```
# Contain the binary with alpineOS
docker build -t $(DOCKER_REPO)/$(IMAGE_NAME):$(IMAGE_TAG) .
# run the docker
sudo docker run -d --name $(IMAGE_NAME) -p 9010:80 $(DOCKER_REPO)/$(IMAGE_NAME):$(IMAGE_TAG)
```
The commands might require sudo or root access

There's also a Makefile included to automate the process.

To build the app with `Dockerfile.build` using the Makefile, simply write
```
make build
```
To build the image containing the app and deploy the container write
```
make deploy
```
To do both at the same time with single command, just type the following line and the `Makefile` will defaultly build and run the containers
```
make
```

The make defaultly removes the first image and container. To change this behavior, change/remove the following lines from the `Makefile`
```
docker rm $(BUILD_NAME)
docker rmi $(BUILD_NAME):$(IMAGE_TAG)
```
