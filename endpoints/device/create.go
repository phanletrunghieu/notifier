package device

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"github.com/phanletrunghieu/notifier/domain"
	"github.com/phanletrunghieu/notifier/service"
)

// RequestCreateDevice .
type RequestCreateDevice struct {
	Device domain.Device `json:"device"`
}

// ResponseCreateDevice .
type ResponseCreateDevice struct {
	Device domain.Device `json:"device"`
}

// MakeCreateEndpoint .
func MakeCreateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			req    = request.(RequestCreateDevice)
			device = &req.Device
		)

		err := s.DeviceService.Create(device)
		if err != nil {
			return nil, err
		}

		return ResponseCreateDevice{Device: *device}, nil
	}
}
