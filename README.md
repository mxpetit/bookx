[![Build Status](https://travis-ci.org/mxpetit/bookx.svg?branch=master)](https://travis-ci.org/mxpetit/bookx)

## Installation

Simply clone this project :
```sh
git clone https://github.com/mxpetit/bookx.git
```

And then run :
```sh
cd $GOPATH/src/github.com/mxpetit/bookx
go run bookx/main.go
```

Bookx need a [cassandra](http://cassandra.apache.org/) instance to be up. You should set these environnement variable :
```sh
BOOKX_IP=127.0.0.1
BOOKX_KEYSPACE=bookx
```

## Tests

You can use [govendor](https://github.com/kardianos/govendor) to run tests :
```sh
govendor test -short -race -bench=. +local
```

Or the vanilla way :
```sh
go test -short -race $(go list ./... | grep -v /vendor/)
```
