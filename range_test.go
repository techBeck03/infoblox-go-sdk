// +build all unittests

package infoblox

import (
	"os"
	"testing"
)

var (
	rangeConfig = Config{
		Host:                   os.Getenv("INFOBLOX_HOST"),
		Port:                   os.Getenv("INFOBLOX_PORT"),
		Username:               os.Getenv("INFOBLOX_USERNAME"),
		Password:               os.Getenv("INFOBLOX_PASSWORD"),
		Version:                os.Getenv("INFOBLOX_VERSION"),
		DisableTLSVerification: true,
	}
	rangeClient = New(rangeConfig)
	testRange   = Range{
		CIDR:         "172.19.4.0/24",
		NetworkView:  "default",
		StartAddress: "172.19.4.10",
		EndAddress:   "172.19.4.20",
		Comment:      "Test Block",
		DisableDHCP:  newBool(true),
		ExtensibleAttributes: newExtensibleAttribute(ExtensibleAttribute{
			"Owner": ExtensibleAttributeValue{
				Value: "testUser",
			},
		}),
	}
	rangeTestNetwork = Network{
		CIDR:        "172.19.4.0/24",
		NetworkView: "default",
		Comment:     "testing",
		ExtensibleAttributes: newExtensibleAttribute(ExtensibleAttribute{
			"Owner": ExtensibleAttributeValue{
				Value: "testUser",
			},
			"Gateway": ExtensibleAttributeValue{
				Value: "172.19.4.1",
			},
		}),
	}
)

func TestCreateRange(t *testing.T) {
	err := rangeClient.CreateNetwork(&rangeTestNetwork)
	if err != nil {
		t.Errorf("Error creating network: %s", err)
	}
	err = rangeClient.CreateRange(&testRange)
	if err != nil {
		t.Errorf("Error creating range: %s", err)
	}
}

func TestGetRange(t *testing.T) {
	rangeObject, err := rangeClient.GetRangeByRef(testRange.Ref)
	if err != nil {
		t.Errorf("Error retrieving range: %s", err)
	}
	prettyPrint(rangeObject)
}

func TestUpdateRange(t *testing.T) {
	updates := Range{
		EndAddress: "172.19.4.22",
		ExtensibleAttributes: newExtensibleAttribute(ExtensibleAttribute{
			"Location": ExtensibleAttributeValue{
				Value: "austin",
			},
		}),
	}
	rangeObject, err := rangeClient.UpdateRange(testRange.Ref, updates)
	if err != nil {
		t.Errorf("Error retrieving range: %s", err)
	}
	eas := *rangeObject.ExtensibleAttributes
	if eas["Location"].Value.(string) != "austin" {
		t.Errorf("Error updating range. EA value does not match expected value")
	}
	prettyPrint(rangeObject)
	testRange = rangeObject
}

func TestDeleteRange(t *testing.T) {
	err := rangeClient.DeleteRange(testRange.Ref)
	if err != nil {
		t.Errorf("Error deleting range: %s", err)
	}
	err = rangeClient.DeleteNetwork(rangeTestNetwork.Ref)
	if err != nil {
		t.Errorf("Error deleting network: %s", err)
	}
}

func TestLogoutRange(t *testing.T) {
	err := rangeClient.Logout()
	if err != nil {
		t.Errorf("Error logging out: %s", err)
	}
}
