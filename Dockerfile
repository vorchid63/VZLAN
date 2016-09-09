FROM golang:1.5
RUN go get github.com/tools/godep
RUN go get github.com/Sirupsen/logrus
RUN go get github.com/codegangsta/cli
RUN go get github.com/docker/go-plugins-helpers/network
RUN go get github.com/docker/libnetwork/types
RUN go get github.com/samalba/dockerclient
RUN go get github.com/vishvananda/netlink
COPY . /go/src/github.com/vorchid63/vzlan
WORKDIR /go/src/github.com/vorchid63/vzlan
RUN cd vzdriver
RUN godep go build 
RUN godep go install
RUN cd ..
RUN godep go install -v
ENTRYPOINT ["vzlan"]