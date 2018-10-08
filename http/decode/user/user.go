package user

import (
	"context"
	"encoding/json"
	"net/http"

	userEndpoints "github.com/phanletrunghieu/notifier/endpoints/user"
)

// DecodeCreateUser .
func DecodeCreateUser(ctx context.Context, r *http.Request) (interface{}, error) {
	var req userEndpoints.RequestCreateUser
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}
