FROM golang:1.7.3
MAINTAINER mxpetit <maximepetit@hotmail.fr>

# Setup work directory
ADD . /go/src/github.com/mxpetit/bookx
WORKDIR /go/src/github.com/mxpetit/bookx

# Get wait-for-it
ADD https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh ./wait-for-it.sh
RUN chmod +x ./wait-for-it.sh

# Install dependencies
RUN go get -u github.com/kardianos/govendor
RUN govendor sync -v

# Install bookx
RUN go install github.com/mxpetit/bookx/bookx

# Setup environment
ENV BOOKX_KEYSPACE=bookx
ENV BOOKX_IP=172.17.0.2
ENV BOOKX_PORT=8080

EXPOSE 8080
