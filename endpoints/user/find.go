package user

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"github.com/phanletrunghieu/notifier/domain"
	"github.com/phanletrunghieu/notifier/service"
)

// RequestFindUser .
type RequestFindUser struct {
	UserID domain.UUID `json:"user_id"`
}

// ResponseFindUser .
type ResponseFindUser struct {
	User domain.User `json:"user"`
}

// MakeFindEndpoint .
func MakeFindEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RequestFindUser)
		user := &domain.User{Model: domain.Model{ID: req.UserID}}

		user, err := s.UserService.Find(user)
		if err != nil {
			return nil, err
		}

		return ResponseFindUser{User: *user}, nil
	}
}
