// +build all unittests specific

package infoblox

import (
	"os"
	"testing"
)

var (
	hostRecordConfig = Config{
		Host:                   os.Getenv("INFOBLOX_HOST"),
		Port:                   os.Getenv("INFOBLOX_PORT"),
		Username:               os.Getenv("INFOBLOX_USERNAME"),
		Password:               os.Getenv("INFOBLOX_PASSWORD"),
		Version:                os.Getenv("INFOBLOX_VERSION"),
		DisableTLSVerification: true,
	}
	hostRecordClient = New(hostRecordConfig)
	testHostRecord   = HostRecord{
		Hostname:    "test-api-1.auslab.cisco.com",
		NetworkView: "default",
		EnableDNS:   true,
		IPv4Addrs: []IPv4Addr{
			{
				//IPAddress: "func:nextavailableip:172.19.4.0/24",
				IPAddress: "172.19.4.2",
			},
		},
		ExtensibleAttributes: newExtensibleAttribute(ExtensibleAttribute{
			"Owner": ExtensibleAttributeValue{
				Value: "testUser",
			},
		}),
	}
	hostRecordTestNetwork = Network{
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

func TestCreateHostRecord(t *testing.T) {
	err := hostRecordClient.CreateNetwork(&hostRecordTestNetwork)
	if err != nil {
		t.Errorf("Error creating network: %s", err)
	}
	err = hostRecordClient.CreateHostRecord(&testHostRecord)
	if err != nil {
		t.Errorf("Error creating host record: %s", err)
	}
}

func TestGetHostRecord(t *testing.T) {
	record, err := hostRecordClient.GetHostRecordByRef(testHostRecord.Ref, nil)
	if err != nil {
		t.Errorf("Error retrieving host record: %s", err)
	}
	prettyPrint(record)
	flattenedEAs, err := hostRecordClient.ConvertEAsToJSONString(*record.ExtensibleAttributes)
	if err != nil {
		t.Errorf("Error flattening eas for host record: %s", err)
	}
	prettyPrint(flattenedEAs)
}

func TestUpdateHostRecord(t *testing.T) {
	updates := HostRecord{
		Hostname: "test-api-2.auslab.cisco.com",
		ExtensibleAttributes: newExtensibleAttribute(ExtensibleAttribute{
			"Location": ExtensibleAttributeValue{
				Value: "austin",
			},
		}),
	}
	record, err := hostRecordClient.UpdateHostRecord(testHostRecord.Ref, updates)
	if err != nil {
		t.Errorf("Error retrieving host record: %s", err)
	}
	eas := *record.ExtensibleAttributes
	if eas["Location"].Value.(string) != "austin" {
		t.Errorf("Error updating host record. EA value does not match expected value")
	}
	prettyPrint(record)
	testHostRecord = record
}

func TestDeleteHostRecord(t *testing.T) {
	err := hostRecordClient.DeleteNetwork(hostRecordTestNetwork.Ref)
	if err != nil {
		t.Errorf("Error deleting network: %s", err)
	}
	err = hostRecordClient.DeleteHostRecord(testHostRecord.Ref)
	if err != nil {
		t.Errorf("Error deleting host record: %s", err)
	}
}

func TestLogoutHostRecord(t *testing.T) {
	err := hostRecordClient.Logout()
	if err != nil {
		t.Errorf("Error logging out: %s", err)
	}
}
