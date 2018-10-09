package device

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"github.com/phanletrunghieu/notifier/domain"
	"github.com/phanletrunghieu/notifier/service"
)

// RequestFindDevice .
type RequestFindDevice struct {
	DeviceID domain.UUID `json:"device_id"`
}

// ResponseFindDevice .
type ResponseFindDevice struct {
	Device domain.Device `json:"device"`
}

// MakeFindEndpoint .
func MakeFindEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RequestFindDevice)
		device := &domain.Device{Model: domain.Model{ID: req.DeviceID}}

		device, err := s.DeviceService.Find(device)
		if err != nil {
			return nil, err
		}

		return ResponseFindDevice{Device: *device}, nil
	}
}
