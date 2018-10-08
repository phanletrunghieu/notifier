package http

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"

	"github.com/phanletrunghieu/notifier/endpoints"
	decodeUser "github.com/phanletrunghieu/notifier/http/decode/user"
)

// NewHTTPHandler .
func NewHTTPHandler(logger log.Logger, endpoints endpoints.Endpoints) http.Handler {
	r := chi.NewRouter()

	options := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
	}

	r.Get("/_warm", httptransport.NewServer(
		endpoint.Nop,
		httptransport.NopRequestDecoder,
		httptransport.EncodeJSONResponse,
		options...,
	).ServeHTTP)

	r.Route("/users", func(r chi.Router) {
		r.Post("/", httptransport.NewServer(
			endpoints.UserCreate,
			decodeUser.DecodeCreateUser,
			httptransport.EncodeJSONResponse,
			options...,
		).ServeHTTP)
	})

	return r
}
