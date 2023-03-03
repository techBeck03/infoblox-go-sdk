//go:build all || unittests || specific
// +build all unittests specific

package infoblox

import (
	"os"
	"testing"
)

var (
	containerConfig = Config{
		Host:                   os.Getenv("INFOBLOX_HOST"),
		Port:                   os.Getenv("INFOBLOX_PORT"),
		Username:               os.Getenv("INFOBLOX_USERNAME"),
		Password:               os.Getenv("INFOBLOX_PASSWORD"),
		Version:                os.Getenv("INFOBLOX_VERSION"),
		DisableTLSVerification: true,
	}
	containerClient = New(containerConfig)
	testContainer   = NetworkContainer{
		CIDR:        "172.19.10.0/23",
		NetworkView: "default",
		Comment:     "Container Testing",
		ExtensibleAttributes: newExtensibleAttribute(ExtensibleAttribute{
			"Owner": ExtensibleAttributeValue{
				Value: "testUser",
			},
		}),
	}
)

func TestCreateContainer(t *testing.T) {
	err := containerClient.CreateContainer(&testContainer)
	if err != nil {
		t.Errorf("Error creating container: %s", err)
	}
}

func TestGetContainer(t *testing.T) {
	container, err := containerClient.GetContainerByRef(testContainer.Ref, nil)
	if err != nil {
		t.Errorf("Error retrieving container: %s", err)
	}
	prettyPrint(container)
}

func TestUpdateContainer(t *testing.T) {
	updates := NetworkContainer{
		Comment: "Container Testing Updated",
		ExtensibleAttributes: newExtensibleAttribute(ExtensibleAttribute{
			"Location": ExtensibleAttributeValue{
				Value: "austin",
			},
		}),
	}
	container, err := containerClient.UpdateContainer(testContainer.Ref, updates)
	if err != nil {
		t.Errorf("Error retrieving container: %s", err)
	}
	eas := *container.ExtensibleAttributes
	if eas["Location"].Value.(string) != "austin" {
		t.Errorf("Error updating container. EA value does not match expected value")
	}
	testContainer = container
}

func TestDeleteContainer(t *testing.T) {
	err := containerClient.DeleteContainer(testContainer.Ref)
	if err != nil {
		t.Errorf("Error deleting container: %s", err)
	}
}

func TestLogoutContainer(t *testing.T) {
	err := containerClient.Logout()
	if err != nil {
		t.Errorf("Error logging out: %s", err)
	}
}
