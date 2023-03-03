package infoblox

import (
	"fmt"
	"net/http"
)

const (
	networkBasePath     = "network"
	networkReturnFields = "network,network_view,comment,extattrs,members,options"
)

// GetNetworkByRef gets network by reference
func (c *Client) GetNetworkByRef(ref string, queryParams map[string]string) (Network, error) {
	var ret Network
	if queryParams == nil {
		queryParams = map[string]string{
			"_return_fields": networkReturnFields,
		}
	} else {
		queryParams["_return_fields"] = networkReturnFields
	}

	queryParamString := c.BuildQuery(queryParams)

	request, err := c.CreateJSONRequest(http.MethodGet, fmt.Sprintf("%s?%s", ref, queryParamString), nil)
	if err != nil {
		return ret, err
	}

	response := c.Call(request, &ret)
	if response != nil {
		return ret, fmt.Errorf(response.ErrorMessage)
	}

	return ret, nil
}

// GetNetworkByQuery gets network by query parameters
func (c *Client) GetNetworkByQuery(queryParams map[string]string) ([]Network, error) {
	var ret NetworkQueryResult
	queryParams["_return_fields"] = networkReturnFields
	queryParams["_return_as_object"] = "1"
	queryParams["_paging"] = "0"
	queryParams["_max_results"] = "2"

	queryParamString := c.BuildQuery(queryParams)
	request, err := c.CreateJSONRequest(http.MethodGet, fmt.Sprintf("%s?%s", networkBasePath, queryParamString), nil)
	if err != nil {
		return nil, err
	}

	response := c.Call(request, &ret)
	if response != nil {
		return nil, fmt.Errorf(response.ErrorMessage)
	}

	return ret.Results, nil
}

// CreateNetwork creates network
func (c *Client) CreateNetwork(network *Network) error {
	queryParams := map[string]string{
		"_return_fields": networkReturnFields,
	}
	queryParamString := c.BuildQuery(queryParams)
	request, err := c.CreateJSONRequest(http.MethodPost, fmt.Sprintf("%s?%s", networkBasePath, queryParamString), network)
	if err != nil {
		return err
	}

	response := c.Call(request, &network)
	if response != nil {
		return fmt.Errorf(response.ErrorMessage)
	}
	return nil
}

// CreateNetworkFromContainer creates network
func (c *Client) CreateNetworkFromContainer(container *NetworkFromContainer) (Network, error) {
	var ret Network
	queryParams := map[string]string{
		"_return_fields":    networkReturnFields,
		"_return_as_object": "1",
	}
	queryParamString := c.BuildQuery(queryParams)
	request, err := c.CreateJSONRequest(http.MethodPost, fmt.Sprintf("%s?%s", networkBasePath, queryParamString), container)
	if err != nil {
		return ret, err
	}

	var result NetworkFromContainerResult
	response := c.Call(request, &result)
	if response != nil {
		return ret, fmt.Errorf(response.ErrorMessage)
	}
	ret, err = c.GetNetworkByRef(result.Result.Ref, nil)
	if err != nil {
		return ret, fmt.Errorf(response.ErrorMessage)
	}

	return ret, nil
}

// UpdateNetwork updates network
func (c *Client) UpdateNetwork(ref string, network Network) (Network, error) {
	var ret Network
	queryParams := map[string]string{
		"_return_fields": networkReturnFields,
	}
	queryParamString := c.BuildQuery(queryParams)
	request, err := c.CreateJSONRequest(http.MethodPut, fmt.Sprintf("%s?%s", ref, queryParamString), network)
	if err != nil {
		return ret, err
	}

	response := c.Call(request, &ret)
	if response != nil {
		return ret, fmt.Errorf(response.ErrorMessage)
	}
	return ret, nil
}

// DeleteNetwork deletes network
func (c *Client) DeleteNetwork(ref string) error {
	request, err := c.CreateJSONRequest(http.MethodDelete, ref, nil)
	if err != nil {
		return err
	}

	response := c.Call(request, nil)
	if response != nil {
		if response.StatusCode == 404 {
			return nil
		}
		return err
	}
	return nil
}
