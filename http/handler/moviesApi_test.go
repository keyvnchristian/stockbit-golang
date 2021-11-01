package handler_test

import (
	"context"
	"log"
	"net"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"

	"stockbit-golang/gen/proto"
	"stockbit-golang/http/handler"
)

var _ = Describe("GetMoviesHandler", func() {
	var ctx context.Context
	var listener *bufconn.Listener

	BeforeEach(func() {
		ctx = context.Background()
		listener = bufconn.Listen(1024 * 1024)
		s := grpc.NewServer()
		proto.RegisterMoviesApiServer(s, &handler.MoviesApiServer{})
		go func() {
			if err := s.Serve(listener); err != nil {
				log.Fatalf("Serve exited with error: %v", err)
			}
		}()
	})

	Describe("Fetching movies", func() {
		var client proto.MoviesApiClient


		BeforeEach(func ()  {
			conn, _ := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
				return listener.Dial()
			}), grpc.WithInsecure())
			client = proto.NewMoviesApiClient(conn)
			
			viper.SetConfigType("yaml")
			viper.AddConfigPath("./../../config")
			viper.SetConfigName("config")
			if err := viper.ReadInConfig(); err != nil {
				log.Printf("Fatal error config file: %s\n", err)
			}
		})

		It("return list of movies with searchword and pagination", func() {
			response, err := client.GetMovies(ctx, &proto.GetMoviesRequest{
				Searchword: "Batman",
				Page: 1,
			})
			if err != nil {
				log.Printf("Error %v", err)
			}
			Expect(err).NotTo(HaveOccurred())
			Expect(len(response.Search)).To(Equal(10))
		})

		It("return movie detail with imdbID", func() {
			response, err := client.GetMovieByID(ctx, &proto.GetMovieByIDRequest{
				Id: "tt5671882",
			})
			if err != nil {
				log.Printf("Error %v", err)
			}
			Expect(err).NotTo(HaveOccurred())
			Expect(response.Actors).To(Equal("Virgil Aioanei, Sabrina Iaschevici, Emilian Marnea"))
			Expect(response.Title).To(Equal("Spid"))
		})
	})
})