# Pull my version of go from docker hub.  Alpine version
FROM golang:1.12.1-alpine3.9

# Set my working directory
WORKDIR $GOPATH/src/github.com/martinrocket/rktServer

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

RUN go build .

EXPOSE 8080
CMD ["./rktServer"]