// +build all unittests

package infoblox

import (
	"os"
	"testing"
)

var (
	fixedAddressConfig = Config{
		Host:                   os.Getenv("INFOBLOX_HOST"),
		Port:                   os.Getenv("INFOBLOX_PORT"),
		Username:               os.Getenv("INFOBLOX_USERNAME"),
		Password:               os.Getenv("INFOBLOX_PASSWORD"),
		Version:                os.Getenv("INFOBLOX_VERSION"),
		DisableTLSVerification: true,
	}
	fixedAddressClient = New(fixedAddressConfig)
	testFixedAddress   = FixedAddress{
		IPAddress:   "172.19.4.1",
		Comment:     "Gateway",
		MatchClient: "RESERVED",
		ExtensibleAttributes: newExtensibleAttribute(ExtensibleAttribute{
			"Owner": ExtensibleAttributeValue{
				Value: "testUser",
			},
		}),
	}
	fixedAddressTestNetwork = Network{
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

func TestCreateFixedAddress(t *testing.T) {
	err := networkClient.CreateNetwork(&fixedAddressTestNetwork)
	if err != nil {
		t.Errorf("Error creating network: %s", err)
	}
	err = fixedAddressClient.CreateFixedAddress(&testFixedAddress)
	if err != nil {
		t.Errorf("Error creating fixed address: %s", err)
	}
}

func TestGetFixedAddress(t *testing.T) {
	record, err := fixedAddressClient.GetFixedAddressByRef(testFixedAddress.Ref)
	if err != nil {
		t.Errorf("Error retrieving fixed address: %s", err)
	}
	prettyPrint(record)
}

func TestUpdateFixedAddress(t *testing.T) {
	updates := FixedAddress{
		Comment: "testing2",
		ExtensibleAttributes: newExtensibleAttribute(ExtensibleAttribute{
			"Location": ExtensibleAttributeValue{
				Value: "austin",
			},
		}),
	}
	record, err := fixedAddressClient.UpdateFixedAddress(testFixedAddress.Ref, updates)
	if err != nil {
		t.Errorf("Error retrieving fixed address: %s", err)
	}
	eas := *record.ExtensibleAttributes
	if eas["Location"].Value.(string) != "austin" {
		t.Errorf("Error updating fixed address. EA value does not match expected value")
	}
	if record.Comment != "testing2" {
		t.Errorf("Error updating fixed address.  Comment string does not match expected value")
	}
	prettyPrint(record)
	testFixedAddress = record
}

func TestDeleteFixedAddress(t *testing.T) {
	err := fixedAddressClient.DeleteFixedAddress(testFixedAddress.Ref)
	if err != nil {
		t.Errorf("Error deleting fixed address: %s", err)
	}
	err = networkClient.DeleteNetwork(fixedAddressTestNetwork.Ref)
	if err != nil {
		t.Errorf("Error deleting network: %s", err)
	}
}

func TestLogoutFixedAddress(t *testing.T) {
	err := fixedAddressClient.Logout()
	if err != nil {
		t.Errorf("Error logging out: %s", err)
	}
}
