FROM alpine

WORKDIR /go/src/github.com/tauki/bluebeak-test-pe
COPY main /go/src/github.com/tauki/bluebeak-test-pe
COPY keys /go/src/github.com/tauki/bluebeak-test-pe/keys

# RUN apk add --update --no-cache ca-certificates

ENV GOPATH /go

ENTRYPOINT /go/src/github.com/tauki/bluebeak-test-pe/main run

# Service listens on port 9010
EXPOSE 9010