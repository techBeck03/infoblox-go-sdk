package infoblox

import (
	"fmt"
	"net/http"
)

const (
	networkBasePath = "network"
)

// GetNetworkByRef gets host record by reference
func (c *Client) GetNetworkByRef(ref string) (Network, error) {
	var ret Network
	queryParams := map[string]string{
		"_return_fields": "network,network_view,comment,extattrs",
	}
	queryParamString := c.BuildQuery(queryParams)
	request, err := c.CreateJSONRequest(http.MethodGet, fmt.Sprintf("%s?%s", ref, queryParamString), nil)
	if err != nil {
		return ret, err
	}

	err = c.Call(request, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

// CreateNetwork creates host record
func (c *Client) CreateNetwork(network *Network) error {
	queryParams := map[string]string{
		"_return_fields": "network,network_view,comment,extattrs",
	}
	queryParamString := c.BuildQuery(queryParams)
	request, err := c.CreateJSONRequest(http.MethodPost, fmt.Sprintf("%s?%s", networkBasePath, queryParamString), network)
	if err != nil {
		return err
	}

	err = c.Call(request, &network)
	if err != nil {
		return err
	}
	return nil
}

// UpdateNetwork creates host record
func (c *Client) UpdateNetwork(ref string, network Network) (Network, error) {
	var ret Network
	queryParams := map[string]string{
		"_return_fields": "network,network_view,comment,extattrs",
	}
	queryParamString := c.BuildQuery(queryParams)
	request, err := c.CreateJSONRequest(http.MethodPut, fmt.Sprintf("%s?%s", ref, queryParamString), network)
	if err != nil {
		return ret, err
	}

	err = c.Call(request, &ret)
	if err != nil {
		return ret, err
	}
	return ret, nil
}

// DeleteNetwork creates host record
func (c *Client) DeleteNetwork(ref string) error {
	request, err := c.CreateJSONRequest(http.MethodDelete, fmt.Sprintf("%s", ref), nil)
	if err != nil {
		return err
	}

	err = c.Call(request, nil)
	if err != nil {
		return err
	}
	return nil
}
