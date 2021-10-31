package http

import (
	"context"
	"log"
	"net/http"
	proto "stockbit-golang/gen/proto"
	"stockbit-golang/http/handler"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

type App struct {
	Router *runtime.ServeMux
}

func (a *App) InitApp() {
	a.Router = runtime.NewServeMux()
	proto.RegisterMoviesApiHandlerServer(context.Background(), a.Router, &handler.MoviesApiServer{})
	log.Fatalln(http.ListenAndServe("localhost:8083", a.Router))
}

