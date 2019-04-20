FROM golang:latest
ADD . /go/src/github.com/scarbo87/fasthttp_test/
WORKDIR /go/src/github.com/scarbo87/fasthttp_test/
EXPOSE 8080