package main

import (
	"context"
	"fmt"
	proto "stockbit-golang/gen/proto"

	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
	}

	client := proto.NewMoviesApiClient(conn)

	res, err := client.GetMovies(context.Background(), &proto.GetMoviesRequest{
		Searchword: "Batman",
		Page: int32(1),
	})

	if err != nil {
		log.Println(err)
	}

	fmt.Println(res)
}