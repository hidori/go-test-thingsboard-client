package devicemanagementapi

import (
	"context"
	"net/http"

	"github.com/hidori/go-test-thingsboard-client/thingsboard"
	"github.com/pkg/errors"
)

type DeviceProfileId = thingsboard.DeviceProfileId

type DeviceCredentialsCredentialsType = thingsboard.DeviceCredentialsCredentialsType

type DeviceCredentialsValue struct {
	CredentialsId   string                                        `json:"credentialsId"`
	CredentialsType *thingsboard.DeviceCredentialsCredentialsType `json:"credentialsType,omitempty"`
}

type DeviceValues struct {
	Name      string                      `json:"name"`
	Label     *string                     `json:"label,omitempty"`
	ProfileId thingsboard.DeviceProfileId `json:"deviceProfileId"`
}

type PostDeviceRequest struct {
	Device      DeviceValues           `json:"device"`
	Credentials DeviceCredentialsValue `json:"credentials"`
}

type PostDeviceResponse thingsboard.Device

func (c *DeviceManagementAPI) GetPostDeviceURL() string {
	return c.server + "/api/device"
}

func (c *DeviceManagementAPI) PostDevice(ctx context.Context, request *PostDeviceRequest) (*PostDeviceResponse, error) {
	url := c.GetPostDeviceURL()

	response, err := Post[PostDeviceRequest, PostDeviceResponse](ctx, url, request, http.StatusOK)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return response, nil
}
