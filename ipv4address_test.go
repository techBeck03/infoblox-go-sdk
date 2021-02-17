// +build all unittests

package infoblox

import (
	"os"
	"testing"
)

var (
	ipv4AddressConfig = Config{
		Host:                   os.Getenv("INFOBLOX_HOST"),
		Port:                   os.Getenv("INFOBLOX_PORT"),
		Username:               os.Getenv("INFOBLOX_USERNAME"),
		Password:               os.Getenv("INFOBLOX_PASSWORD"),
		Version:                os.Getenv("INFOBLOX_VERSION"),
		DisableTLSVerification: true,
	}
	ipv4AddressClient          = New(ipv4AddressConfig)
	ipv4AddressSequentialQuery = AddressQuery{
		CIDR:  "172.16.106.0/24",
		Count: 7,
	}
	ipv4AddressVerifyQuery = AddressQuery{
		CIDR:         "172.19.4.0/24",
		StartAddress: "172.19.4.10",
		EndAddress:   "172.19.4.20",
	}
	ipv4AddressTestRange = Range{
		CIDR:         "172.19.4.0/24",
		NetworkView:  "default",
		StartAddress: "172.19.4.10",
		EndAddress:   "172.19.4.20",
		Comment:      "Test Block",
		DisableDHCP:  true,
		ExtensibleAttributes: newExtensibleAttribute(ExtensibleAttribute{
			"Owner": ExtensibleAttributeValue{
				Value: "testUser",
			},
		}),
	}
	ipv4AddressTestNetwork = Network{
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
	ipv4AddressHostRecord = HostRecord{
		Hostname:    "test-api-12.auslab.cisco.com",
		NetworkView: "default",
		EnableDNS:   true,
		IPv4Addrs: []IPv4Addr{
			{
				//IPAddress: "func:nextavailableip:172.19.4.0/24",
				IPAddress: "172.19.4.12",
			},
		},
		ExtensibleAttributes: newExtensibleAttribute(ExtensibleAttribute{
			"Owner": ExtensibleAttributeValue{
				Value: "testUser",
			},
		}),
	}
)

func TestGetSequentialAddressRange(t *testing.T) {
	addresses, err := ipv4AddressClient.GetSequentialAddressRange(ipv4AddressSequentialQuery)
	if err != nil {
		t.Errorf("Error retrieving host record: %s", err)
	}
	prettyPrint(addresses)
}

func TestGetUsedAddressesWithinRange(t *testing.T) {
	err := ipv4AddressClient.CreateNetwork(&ipv4AddressTestNetwork)
	if err != nil {
		t.Errorf("Error creating network: %s", err)
	}
	err = ipv4AddressClient.CreateRange(&ipv4AddressTestRange)
	if err != nil {
		t.Errorf("Error creating range: %s", err)
	}
	err = ipv4AddressClient.CreateHostRecord(&ipv4AddressHostRecord)
	if err != nil {
		t.Errorf("Error creating host record: %s", err)
	}
	addresses, err := ipv4AddressClient.GetUsedAddressesWithinRange(ipv4AddressVerifyQuery)
	if err != nil {
		t.Errorf("Error retrieving host record: %s", err)
	}
	prettyPrint(addresses)
	err = ipv4AddressClient.DeleteHostRecord(ipv4AddressHostRecord.Ref)
	if err != nil {
		t.Errorf("Error deleting host record: %s", err)
	}
	err = ipv4AddressClient.DeleteRange(ipv4AddressTestRange.Ref)
	if err != nil {
		t.Errorf("Error deleting range: %s", err)
	}
	err = ipv4AddressClient.DeleteNetwork(ipv4AddressTestNetwork.Ref)
	if err != nil {
		t.Errorf("Error deleting network: %s", err)
	}
}

func TestLogoutIPv4Address(t *testing.T) {
	err := ipv4AddressClient.Logout()
	if err != nil {
		t.Errorf("Error logging out: %s", err)
	}
}
