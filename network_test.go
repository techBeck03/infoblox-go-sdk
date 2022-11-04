//go:build all || unittests || specific
// +build all unittests specific

package infoblox

import (
	"os"
	"testing"
)

var (
	networkConfig = Config{
		Host:                   os.Getenv("INFOBLOX_HOST"),
		Port:                   os.Getenv("INFOBLOX_PORT"),
		Username:               os.Getenv("INFOBLOX_USERNAME"),
		Password:               os.Getenv("INFOBLOX_PASSWORD"),
		Version:                os.Getenv("INFOBLOX_VERSION"),
		DisableTLSVerification: true,
	}
	networkClient = New(networkConfig)
	testNetwork   = Network{
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
	testNetworkFromContainer = NetworkFromContainer{
		Network: NetworkContainerFunction{
			Function:    "next_available_network",
			ResultField: "networks",
			Object:      "networkcontainer",
			ObjectParameters: newNetworkContainerObjectParameter(NetworkContainerObjectParameter{
				Label: "Autonets",
			}),
			Parameters: newNetworkContainerParameter(NetworkContainerParameter{
				Prefix: 24,
			}),
		},
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

var testNetworkFromContainerResolved Network

func TestCreateNetwork(t *testing.T) {
	err := networkClient.CreateNetwork(&testNetwork)
	if err != nil {
		t.Errorf("Error creating network: %s", err)
	}
}

func TestGetNetwork(t *testing.T) {
	network, err := networkClient.GetNetworkByRef(testNetwork.Ref, nil)
	if err != nil {
		t.Errorf("Error retrieving network: %s", err)
	}
	prettyPrint(network)
}

func TestUpdateNetwork(t *testing.T) {
	updates := Network{
		Comment: "testing2",
		ExtensibleAttributes: newExtensibleAttribute(ExtensibleAttribute{
			"Location": ExtensibleAttributeValue{
				Value: "austin",
			},
		}),
	}
	network, err := networkClient.UpdateNetwork(testNetwork.Ref, updates)
	if err != nil {
		t.Errorf("Error retrieving network: %s", err)
	}
	eas := *network.ExtensibleAttributes
	if eas["Location"].Value.(string) != "austin" {
		t.Errorf("Error updating network. EA value does not match expected value")
	}
	prettyPrint(network)
	testNetwork = network
}

func TestDeleteNetwork(t *testing.T) {
	err := networkClient.DeleteNetwork(testNetwork.Ref)
	if err != nil {
		t.Errorf("Error deleting network: %s", err)
	}
}

func TestCreateNetworkFromContainer(t *testing.T) {
	network, err := networkClient.CreateNetworkFromContainer(&testNetworkFromContainer)
	if err != nil {
		t.Errorf("Error creating network: %s", err)
	}
	testNetworkFromContainerResolved = network
}

func TestGetNetworkCreatedFromContainer(t *testing.T) {
	network, err := networkClient.GetNetworkByRef(testNetworkFromContainerResolved.Ref, nil)
	if err != nil {
		t.Errorf("Error retrieving network: %s", err)
	}
	prettyPrint(network)
}

func TestUpdateNetworkCreatedFromContainer(t *testing.T) {
	updates := Network{
		Comment: "testing2",
		ExtensibleAttributesAdd: newExtensibleAttribute(ExtensibleAttribute{
			"Location": ExtensibleAttributeValue{
				Value: "austin",
			},
		}),
	}
	network, err := networkClient.UpdateNetwork(testNetworkFromContainerResolved.Ref, updates)
	if err != nil {
		t.Errorf("Error retrieving network: %s", err)
	}
	eas := *network.ExtensibleAttributes
	if eas["Location"].Value.(string) != "austin" {
		t.Errorf("Error updating network. EA value does not match expected value")
	}
	prettyPrint(network)
}

func TestDeleteNetworkCreatedFromContainer(t *testing.T) {
	err := networkClient.DeleteNetwork(testNetworkFromContainerResolved.Ref)
	if err != nil {
		t.Errorf("Error deleting network: %s", err)
	}
}

func TestLogoutNetwork(t *testing.T) {
	err := networkClient.Logout()
	if err != nil {
		t.Errorf("Error logging out: %s", err)
	}
}
