package user

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/phanletrunghieu/notifier/domain"
	"github.com/phanletrunghieu/notifier/service"
)

// RequestCreateUser .
type RequestCreateUser struct {
	User domain.User `json:"user"`
}

// ResponseCreateUser .
type ResponseCreateUser struct {
	User domain.User `json:"user"`
}

// MakeCreateEndpoint .
func MakeCreateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			req  = request.(RequestCreateUser)
			user = &domain.User{
				Name: req.User.Name,
			}
		)

		err := s.UserService.Create(user)
		if err != nil {
			return nil, err
		}

		return ResponseCreateUser{User: *user}, nil
	}
}
