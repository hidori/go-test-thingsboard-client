package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/hidori/go-pointer"
	devicemanagementapi "github.com/hidori/go-test-thingsboard-client/devicemanagemntapi"
	"github.com/pkg/errors"
)

var (
	server = os.Getenv("SERVER")
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	name := "Name"
	credentialsId := "CredentialsId"

	request := &devicemanagementapi.PostDeviceRequest{
		Device: devicemanagementapi.DeviceValues{
			Name:  name,
			Label: nil,
			ProfileId: devicemanagementapi.DeviceProfileId{
				Id:         uuid.Must(uuid.NewRandom()),
				EntityType: "DEVICE_PROFILE",
			},
		},
		Credentials: devicemanagementapi.DeviceCredentialsValue{
			CredentialsId:   credentialsId,
			CredentialsType: pointer.Of(devicemanagementapi.DeviceCredentialsCredentialsType("ACCESS_TOKEN")),
		},
	}

	b, err := json.MarshalIndent(request, "", "  ")
	if err != nil {
		return errors.WithStack(err)
	}

	log.Println(string(b))

	return nil
}
