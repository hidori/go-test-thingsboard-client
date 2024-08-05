package devicemanagementapi

import (
	"context"
)

type DeviceManagementAPIClient interface {
	PostDevice(ctx context.Context, request *PostDeviceRequest) (*PostDeviceResponse, error)
}

type DeviceManagementAPI struct {
	server string
}

func NewDeviceAPIClient(server string) *DeviceManagementAPI {
	return &DeviceManagementAPI{
		server: server,
	}
}
