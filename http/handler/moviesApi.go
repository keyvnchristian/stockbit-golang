package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"stockbit-golang/config"
	proto "stockbit-golang/gen/proto"

	"github.com/spf13/viper"
)

type MoviesApiServer struct {
	proto.UnimplementedMoviesApiServer
}

func (c *MoviesApiServer) GetMovies(ctx context.Context, req *proto.GetMoviesRequest) (*proto.MoviesResponse, error) {
	url := fmt.Sprintf("%s?apikey=%s&s=%s&page=%d", viper.GetString(config.OmdbURL), viper.GetString(config.OmdbKey), req.Searchword, req.Page)

	response := proto.MoviesResponse{}
	err := getJson(url, &response)
	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}

	return &response, nil
}

func (c *MoviesApiServer) GetMovieByID(ctx context.Context, req *proto.GetMovieByIDRequest) (*proto.MovieByIDResponse, error) {
	url := fmt.Sprintf("%s?apikey=%s&i=%s", viper.GetString(config.OmdbURL), viper.GetString(config.OmdbKey), req.Id)

	response := proto.MovieByIDResponse{}
	err := getJson(url, &response)
	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}

	return &response, nil
}

// func getMovies(url string) {
// 	req, _ := http.NewRequest("GET", url, nil)
// 	res, _ := http.DefaultClient.Do(req)

// 	defer res.Body.Close()

// 	var v interface{}
// 	decoder := json.NewDecoder(res.Body)

// 	err := decoder.Decode(&v)
// 	if err != nil {
// 		log.Printf("%v", err)
// 	}

// 	fmt.Println(v)
// }

func getJson(url string, target interface{}) error {
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)

	err := decoder.Decode(target)
	if err != nil {
		return err
	}

    return err
}