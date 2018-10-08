package endpoints

import (
	"github.com/go-kit/kit/endpoint"
	userEndpoints "github.com/phanletrunghieu/notifier/endpoints/user"
	"github.com/phanletrunghieu/notifier/service"
)

// Endpoints .
type Endpoints struct {
	UserCreate endpoint.Endpoint
}

// NewEndpoints .
func NewEndpoints(s service.Service) Endpoints {
	return Endpoints{
		UserCreate: userEndpoints.MakeCreateEndpoint(s),
	}
}
