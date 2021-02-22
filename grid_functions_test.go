// +build all unittests

package infoblox

import (
	"os"
	"testing"
	"time"
)

var (
	gridConfig = Config{
		Host:                   os.Getenv("INFOBLOX_HOST"),
		Port:                   os.Getenv("INFOBLOX_PORT"),
		Username:               os.Getenv("INFOBLOX_USERNAME"),
		Password:               os.Getenv("INFOBLOX_PASSWORD"),
		Version:                os.Getenv("INFOBLOX_VERSION"),
		DisableTLSVerification: true,
	}
	gridClient      = New(gridConfig)
	gridTestNetwork = Network{
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
		Options: []Option{
			Option{
				Name:        "routers",
				UseOption:   newBool(true),
				Value:       "172.19.4.1",
				VendorClass: "DHCP",
			},
		},
	}
	grid               Grid
	gridMember         GridMember
	gridRestartRequest = GridServiceRestartRequest{
		RestartOption: "RESTART_IF_NEEDED",
		Services:      []string{"DHCP"},
		Members:       []string{gridMember.Hostname},
	}
)

func TestGridCreateNetwork(t *testing.T) {
	members, err := gridClient.GetGridMembers(nil)
	if err != nil {
		t.Errorf("Error retrieving grid members: %s", err)
	}
	gridMember = members[0]
	gridTestNetwork.Members = []Member{
		Member{
			StructType: "dhcpmember",
			Name:       gridMember.Hostname,
		},
	}
	err = gridClient.CreateNetwork(&gridTestNetwork)
	if err != nil {
		t.Errorf("Error creating network: %s", err)
	}
	grids, err := gridClient.GetGrids(nil)
	if err != nil {
		t.Errorf("Error retrieving grids: %s", err)
	}
	grid = grids[0]
	err = gridClient.RestartServices(grid.Ref, gridRestartRequest)
	if err != nil {
		t.Errorf("Error restarting grid services: %s", err)
	}
	time.Sleep(5 * time.Second)
}

func TestGridDeleteNetwork(t *testing.T) {
	err := gridClient.DeleteNetwork(gridTestNetwork.Ref)
	if err != nil {
		t.Errorf("Error deleting network: %s", err)
	}

	err = gridClient.RestartServices(grid.Ref, gridRestartRequest)
	if err != nil {
		t.Errorf("Error restarting grid services: %s", err)
	}
}

func TestLogoutGrid(t *testing.T) {
	err := gridClient.Logout()
	if err != nil {
		t.Errorf("Error logging out: %s", err)
	}
}
