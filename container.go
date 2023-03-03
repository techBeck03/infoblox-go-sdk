package infoblox

import (
	"fmt"
	"net/http"
)

const (
	containerBasePath     = "networkcontainer"
	containerReturnFields = "comment,network,network_view,extattrs"
)

// GetContainerByRef gets A record by reference
func (c *Client) GetContainerByRef(ref string, queryParams map[string]string) (NetworkContainer, error) {
	var ret NetworkContainer
	if queryParams == nil {
		queryParams = map[string]string{
			"_return_fields": containerReturnFields,
		}
	} else {
		queryParams["_return_fields"] = containerReturnFields
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

// GetContainerByQuery gets A records by query parameters
func (c *Client) GetContainerByQuery(queryParams map[string]string) ([]NetworkContainer, error) {
	var ret []NetworkContainer
	queryParams["_return_fields"] = containerReturnFields
	queryParams["_paging"] = "0"
	queryParams["_max_results"] = "2"

	queryParamString := c.BuildQuery(queryParams)
	request, err := c.CreateJSONRequest(http.MethodGet, fmt.Sprintf("%s?%s", containerBasePath, queryParamString), nil)
	if err != nil {
		return nil, err
	}

	response := c.Call(request, &ret)
	if response != nil {
		return nil, fmt.Errorf(response.ErrorMessage)
	}

	return ret, nil
}

// CreateContainer creates A record
func (c *Client) CreateContainer(record *NetworkContainer) error {
	queryParams := map[string]string{
		"_return_fields": containerReturnFields,
	}
	queryParamString := c.BuildQuery(queryParams)
	request, err := c.CreateJSONRequest(http.MethodPost, fmt.Sprintf("%s?%s", containerBasePath, queryParamString), record)
	if err != nil {
		return err
	}

	response := c.Call(request, &record)
	if response != nil {
		return fmt.Errorf(response.ErrorMessage)
	}
	return nil
}

// UpdateContainer creates A record
func (c *Client) UpdateContainer(ref string, network NetworkContainer) (NetworkContainer, error) {
	var ret NetworkContainer
	queryParams := map[string]string{
		"_return_fields": containerReturnFields,
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

// DeleteContainer creates A record
func (c *Client) DeleteContainer(ref string) error {
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
