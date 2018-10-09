package device

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/phanletrunghieu/notifier/domain"
	deviceEndpoints "github.com/phanletrunghieu/notifier/endpoints/device"
)

// DecodeCreateDevice .
func DecodeCreateDevice(ctx context.Context, r *http.Request) (interface{}, error) {
	var req deviceEndpoints.RequestCreateDevice
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// DecodeFindDevice .
func DecodeFindDevice(ctx context.Context, r *http.Request) (interface{}, error) {
	var req deviceEndpoints.RequestFindDevice
	strID := chi.URLParam(r, "device_id")
	deviceID, err := domain.UUIDFromString(strID)
	if err != nil {
		return req, err
	}

	req.DeviceID = deviceID
	return req, err
}

// DecodeUpdateDevice .
func DecodeUpdateDevice(ctx context.Context, r *http.Request) (interface{}, error) {
	var req deviceEndpoints.RequestUpdateDevice
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	strID := chi.URLParam(r, "device_id")
	deviceID, err := domain.UUIDFromString(strID)
	if err != nil {
		return nil, err
	}

	req.Device.ID = deviceID
	return req, nil
}
