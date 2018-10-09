package user

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/phanletrunghieu/notifier/domain"
	userEndpoints "github.com/phanletrunghieu/notifier/endpoints/user"
)

// DecodeCreateUser .
func DecodeCreateUser(ctx context.Context, r *http.Request) (interface{}, error) {
	var req userEndpoints.RequestCreateUser
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// DecodeFindUser .
func DecodeFindUser(ctx context.Context, r *http.Request) (interface{}, error) {
	var req userEndpoints.RequestFindUser
	strID := chi.URLParam(r, "user_id")
	userID, err := domain.UUIDFromString(strID)
	if err != nil {
		return req, err
	}

	req.UserID = userID
	return req, err
}
