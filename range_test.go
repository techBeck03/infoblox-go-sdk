//go:build all || unittests || specific
// +build all unittests specific

package infoblox

import (
	"log"
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
		CIDR:         "172.19.10.0/24",
		NetworkView:  "default",
		StartAddress: "172.19.10.10",
		EndAddress:   "172.19.10.19",
		Comment:      "Test Block",
		DisableDHCP:  newBool(true),
		ExtensibleAttributes: newExtensibleAttribute(ExtensibleAttribute{
			"Owner": ExtensibleAttributeValue{
				Value: "testUser",
			},
		}),
	}
	testRangeSequential = Range{
		CIDR:         "172.19.10.0/24",
		NetworkView:  "default",
		StartAddress: "172.19.10.10",
		EndAddress:   "172.19.10.19",
		Comment:      "Test Block",
		DisableDHCP:  newBool(true),
		ExtensibleAttributes: newExtensibleAttribute(ExtensibleAttribute{
			"Owner": ExtensibleAttributeValue{
				Value: "testUser",
			},
		}),
	}
	rangeTestSequentialQuery1 = AddressQuery{
		CIDR:  "172.19.10.0/24",
		Count: 20,
	}
	rangeTestSequentialQuery2 = AddressQuery{
		CIDR:  "172.19.10.0/24",
		Count: 20,
	}
	rangeTestSequentialQuery3 = AddressQuery{
		CIDR:  "172.19.10.0/24",
		Count: 20,
	}
	rangeTestNetwork = Network{
		CIDR:        "172.19.10.0/24",
		NetworkView: "default",
		Comment:     "testing",
		ExtensibleAttributes: newExtensibleAttribute(ExtensibleAttribute{
			"Owner": ExtensibleAttributeValue{
				Value: "testUser",
			},
			"Gateway": ExtensibleAttributeValue{
				Value: "172.19.10.1",
			},
		}),
	}
)

func TestRange(t *testing.T) {
	t.Cleanup(cleanup)
	t.Run("", createRange)
	t.Run("", getRange)
	t.Run("", updateRange)
	t.Run("", deleteRange)
	t.Run("", createSequentialRange1)
	t.Run("", createSequentialRange2)
	t.Run("", createSequentialRange3)
}

func createRange(t *testing.T) {
	err := rangeClient.CreateNetwork(&rangeTestNetwork)
	if err != nil {
		t.Errorf("Error creating network: %s", err)
	}
	err = rangeClient.CreateRange(&testRange)
	if err != nil {
		t.Errorf("Error creating range: %s", err)
	}
}

func createSequentialRange1(t *testing.T) {
	t.Parallel()
	err := rangeClient.CreateSequentialRange(&testRangeSequential, rangeTestSequentialQuery1)
	if err != nil {
		t.Errorf("Error creating range: %s", err)
	}
}
func createSequentialRange2(t *testing.T) {
	t.Parallel()
	err := rangeClient.CreateSequentialRange(&testRangeSequential, rangeTestSequentialQuery2)
	if err != nil {
		t.Errorf("Error creating range: %s", err)
	}
}
func createSequentialRange3(t *testing.T) {
	t.Parallel()
	err := rangeClient.CreateSequentialRange(&testRangeSequential, rangeTestSequentialQuery3)
	if err != nil {
		t.Errorf("Error creating range: %s", err)
	}
}

func getRange(t *testing.T) {
	rangeObject, err := rangeClient.GetRangeByRef(testRange.Ref, nil)
	if err != nil {
		t.Errorf("Error retrieving range: %s", err)
	}
	prettyPrint(rangeObject)
}

func updateRange(t *testing.T) {
	updates := Range{
		EndAddress: "172.19.10.22",
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

func deleteRange(t *testing.T) {
	err := rangeClient.DeleteRange(testRange.Ref)
	if err != nil {
		t.Errorf("Error deleting range: %s", err)
	}
}

func cleanup() {
	err := rangeClient.DeleteNetwork(rangeTestNetwork.Ref)
	if err != nil {
		log.Printf("Error deleting network: %s", err)
	}
	err = rangeClient.Logout()
	if err != nil {
		log.Printf("Error logging out: %s", err)
	}
}
