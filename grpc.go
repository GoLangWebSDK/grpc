package grpc

import (
	"net/http"

	grpchealth "github.com/bufbuild/connect-grpchealth-go"
)

type GRPCHandlers interface {
	LoadHandlers()
	LoadCheckers()
}

type GRPC struct {
	Mux *http.ServeMux
}

func NewRouter() *GRPC {
	return &GRPC{
		Mux: http.NewServeMux(),
	}
}

func (router *GRPC) Load(handlers GRPCHandlers) {
	if handlers != nil {
		handlers.LoadCheckers()
		handlers.LoadHandlers()
	}
}

func (router *GRPC) AddGrpcHandler(pattern string, handler http.Handler) {
	router.Mux.Handle(pattern, handler)
}

func (router *GRPC) AddGrpcChecker(checker *grpchealth.StaticChecker) {
	router.Mux.Handle(grpchealth.NewHandler(checker))
}
