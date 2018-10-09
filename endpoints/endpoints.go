package endpoints

import (
	"github.com/go-kit/kit/endpoint"

	deviceEndpoints "github.com/phanletrunghieu/notifier/endpoints/device"
	userEndpoints "github.com/phanletrunghieu/notifier/endpoints/user"
	"github.com/phanletrunghieu/notifier/service"
)

// Endpoints .
type Endpoints struct {
	UserCreate  endpoint.Endpoint
	UserFind    endpoint.Endpoint
	UserFindAll endpoint.Endpoint

	DeviceCreate  endpoint.Endpoint
	DeviceFind    endpoint.Endpoint
	DeviceFindAll endpoint.Endpoint
	DeviceUpdate  endpoint.Endpoint
}

// NewEndpoints .
func NewEndpoints(s service.Service) Endpoints {
	return Endpoints{
		UserCreate:  userEndpoints.MakeCreateEndpoint(s),
		UserFind:    userEndpoints.MakeFindEndpoint(s),
		UserFindAll: userEndpoints.MakeFindAllEndpoint(s),

		DeviceCreate:  deviceEndpoints.MakeCreateEndpoint(s),
		DeviceFind:    deviceEndpoints.MakeFindEndpoint(s),
		DeviceFindAll: deviceEndpoints.MakeFindAllEndpoint(s),
		DeviceUpdate:  deviceEndpoints.MakeUpdateEndpoint(s),
	}
}
