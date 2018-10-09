package device

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"github.com/phanletrunghieu/notifier/domain"
	"github.com/phanletrunghieu/notifier/service"
)

// RequestUpdateDevice .
type RequestUpdateDevice struct {
	Device domain.Device `json:"device"`
}

// ResponseUpdateDevice .
type ResponseUpdateDevice struct {
	Device domain.Device `json:"device"`
}

// MakeUpdateEndpoint .
func MakeUpdateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			req    = request.(RequestUpdateDevice)
			device = &req.Device
		)

		device, err := s.DeviceService.Update(device)
		if err != nil {
			return nil, err
		}

		return ResponseUpdateDevice{Device: *device}, nil
	}
}
