package http

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"

	"github.com/phanletrunghieu/notifier/endpoints"
	decodeDevice "github.com/phanletrunghieu/notifier/http/decode/device"
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

		r.Get("/", httptransport.NewServer(
			endpoints.UserFindAll,
			httptransport.NopRequestDecoder,
			httptransport.EncodeJSONResponse,
			options...,
		).ServeHTTP)

		r.Route("/{user_id}", func(r chi.Router) {
			r.Get("/", httptransport.NewServer(
				endpoints.UserFind,
				decodeUser.DecodeFindUser,
				httptransport.EncodeJSONResponse,
				options...,
			).ServeHTTP)
		})
	})

	r.Route("/devices", func(r chi.Router) {
		r.Post("/", httptransport.NewServer(
			endpoints.DeviceCreate,
			decodeDevice.DecodeCreateDevice,
			httptransport.EncodeJSONResponse,
			options...,
		).ServeHTTP)

		r.Get("/", httptransport.NewServer(
			endpoints.DeviceFindAll,
			httptransport.NopRequestDecoder,
			httptransport.EncodeJSONResponse,
			options...,
		).ServeHTTP)

		r.Route("/{device_id}", func(r chi.Router) {
			r.Get("/", httptransport.NewServer(
				endpoints.DeviceFind,
				decodeDevice.DecodeFindDevice,
				httptransport.EncodeJSONResponse,
				options...,
			).ServeHTTP)

			r.Put("/", httptransport.NewServer(
				endpoints.DeviceUpdate,
				decodeDevice.DecodeUpdateDevice,
				httptransport.EncodeJSONResponse,
				options...,
			).ServeHTTP)
		})
	})

	return r
}
