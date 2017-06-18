FROM golang:1.8.0
MAINTAINER Sara Taha

# installing cachable go dependency
RUN go get -u github.com/kardianos/govendor
RUN go get github.com/onsi/ginkgo/ginkgo
RUN go get github.com/onsi/gomega

# add the current ctgaccounts dir
ADD . /go/src/ctgaccounts
WORKDIR /go/src/ctgaccounts

# installing the dependency
RUN govendor sync

# run with: docker build -t ctgaccounts .
# connect: docker run -i -t ctgaccounts /bin/bash
