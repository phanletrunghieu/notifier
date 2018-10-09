package device

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"github.com/phanletrunghieu/notifier/domain"
	"github.com/phanletrunghieu/notifier/service"
)

// ResponseFindAllDevice .
type ResponseFindAllDevice struct {
	Devices []domain.Device `json:"devices"`
}

// MakeFindAllEndpoint .
func MakeFindAllEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		devices, err := s.DeviceService.FindAll()
		if err != nil {
			return nil, err
		}

		return ResponseFindAllDevice{Devices: devices}, nil
	}
}
