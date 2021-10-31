package main

import (
	"log"
	"net"
	"stockbit-golang/config"
	proto "stockbit-golang/gen/proto"
	"stockbit-golang/http"
	"stockbit-golang/http/handler"

	"google.golang.org/grpc"
)


func init() {
	config.InitConfig()
}

func main() {
	go func() {
		// Allow apps to run on different port, enable to communicate with JSON
		app := http.App{}
		app.InitApp()
	}()

	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	grpcServer := grpc.NewServer()

	proto.RegisterMoviesApiServer(grpcServer, &handler.MoviesApiServer{})

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Println(err)
	}
}