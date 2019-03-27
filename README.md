# bluebeak-test-pe

There are a number of ways to run the program.

To run the program by building from scratch type
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
It will spin-up the server and start listening and serving the APIs. In default the server is set to a dev environment so the logs are expected to be printed after the server is spun-up.

To run the individual scripts without spinning-up the server, just type `./main --help` or `./main -h` or just `./main` and it will printout the flags and arguments required to run individual script.

Just running `./main` will printout the basic information of the app which includes all necessary flags and commands to run the app and individual components.

There are two files included for dockerinzing the app. The `Dockerfile` dockerizes the file but it requires the program to be already built. In default the `Dockerfile` expects the executable to be named `main`, to change that open the `Dockerfile` and change the following line
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
sudo docker run --network="host" -d --name $(IMAGE_NAME) -p 9010:80 $(DOCKER_REPO)/$(IMAGE_NAME):$(IMAGE_TAG)
```

If you notice, the docker is running on the host mode, it is done so it can access to the locally installed mysql since if you look into the configs in `main.go` you'll notice the mysql host is set to `127.0.0.1`, if it is changed, the tag won't be needed for this purpose.

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
To do both at the same time with single command, just type the following line and the `Makefile` will build and run the containers in default.
```
make
```

The make removes the first image and container in default. To change this behavior, change/remove the following lines from the `Makefile`
```
docker rm $(BUILD_NAME)
docker rmi $(BUILD_NAME):$(IMAGE_TAG)
```

As it is mentioned earlier, the docker container is also run on host mode to bridge connection to the locally deployed mysql. If it is changed remove the `--network="host"` tag from the following line of the `Makefile`
```
sudo docker run --network="host" -d --name $(IMAGE_NAME) -p 9010:9010 $(DOCKER_REPO)/$(IMAGE_NAME):$(IMAGE_TAG)
```

# APIs

There is not API documentation included in the project. But the following http routes are currently defined

#
#### Reviews

```
GET /review
```
 This returns a list of reviews from the database as a JSON object.
 
 The structure of the JSON object contains 2 keys, `data` which is an array of review object and the key `next` is the url to fetch the next portion of the reviews. It returns 50 reviews at a time in default.
    
```
POST /review
```
Posting to this address allows to add review to the review database.

This expects a JSON object in the request body with the structure matching the structure defined in the `models/review.Reviews` object, the tags as the json key and datatype as the type of the values.

```
GET /review/list/unique
```
This returns an array in a JSON object with the list of unique reviewers from the DB.

```
GET /review/list/regular
```
This returns an array as a JSON object containing the names of the reviewers with 5 or more reviews.

```
GET review/user/:name
```
This returns the reviews made by the an user that's mentioned in the url-param

#
#### user

```
GET /user
```
This returns a list of users from the `userinfo` table of the database as a JSON object.
 
 The structure of the JSON object contains 2 keys, `data` which is an array of review object and the key `next` is the url to fetch the next portion of the reviews. It returns 50 reviews at a time in default.

 ```
 POST /user
 ```
 Posting to this address allows to add user to the `userinfo` table of the database.
 
 This expects a JSON object in the request body with the structure matching the structure defined in the `models/user.UserInfo` object, the tags as the json key and datatype as the type of the values.
 
 #
 ### HTTP return codes and error message
 
 Error messages are sent in as a JSON object that has the structure as the structure defined in the `models.msg.Message`, in the JSON the `code` key contains the code of the error and the `Message` object contains details about the error. Only the required data are serialized into the object.
 
```
400 - Bad request
Sent when the request sent to the server is to a route fails it's requirements.
```
```
404 - Not Found
Sent when a searched query resulted in 0 results.
```
```
302 - Found
Sent when a searched query resulted in 1 or more results.
```
```
500 - Internal server error
Sent when the server failed to deliver a respond.

```
