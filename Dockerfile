# golang image where workspace (GOPATH) configured at /go.
FROM golang:1.8.0
MAINTAINER Sara Taha

# Install dependencies
RUN go get -u github.com/kardianos/govendor
RUN go get github.com/onsi/ginkgo/ginkgo
RUN go get github.com/onsi/gomega

# add the current ctgaccounts dir
ADD . /go/src/ctgAccounts
WORKDIR /go/src/ctgAccounts

# installing the dependency
RUN govendor sync

# run with: docker build -t ctgaccounts .
# connect: docker run -i -t ctgaccounts /bin/bash
