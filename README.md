# Stockbit-Golang

Stockbit-golang is created to fulfill Stockbit Go Test

## Local

### Install tools

1. Install go: [read here](https://golang.org/dl/)
2. Install [protoc](http://google.github.io/proto-lens/installing-protoc.html) to compile protocol buffers definitions files

```bash
brew install protobuf
```

### Usage

1. First of all you need to run `make create` to compile protobuf def files
2. To run server, execute `make run`
3. To run client, execute `make client`
4. To run test, execute `make test`
5. To send HTTP request and view the response in Visual Studio Code directly, I've prepared `req.http` files or

```bash
curl http://localhost:8083/v1/movies?searchword=Batman&page=1

curl http://localhost:8083/v1/movies/tt0372784
```

## Answer for number 1,3,4

Answer for those numbers are located inside `others` folder