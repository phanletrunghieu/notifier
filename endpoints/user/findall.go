package user

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"github.com/phanletrunghieu/notifier/domain"
	"github.com/phanletrunghieu/notifier/service"
)

// ResponseFindAllUser .
type ResponseFindAllUser struct {
	Users []domain.User `json:"users"`
}

// MakeFindAllEndpoint .
func MakeFindAllEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		users, err := s.UserService.FindAll()
		if err != nil {
			return nil, err
		}

		return ResponseFindAllUser{Users: users}, nil
	}
}
