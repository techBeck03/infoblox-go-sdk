package infoblox

import (
	"fmt"
	"net/http"
)

const (
	fixedAddressBasePath     = "fixedaddress"
	fixedAddressReturnFields = "extattrs,ipv4addr,network_view,disable,comment,name,match_client,mac,network"
)

// GetFixedAddressByRef gets fixed address by reference
func (c *Client) GetFixedAddressByRef(ref string, queryParams map[string]string) (FixedAddress, error) {
	var ret FixedAddress

	if queryParams == nil {
		queryParams = map[string]string{
			"_return_fields": fixedAddressReturnFields,
		}
	} else {
		queryParams["_return_fields"] = fixedAddressReturnFields
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

// GetFixedAddressByQuery gets fixed address by query parameters
func (c *Client) GetFixedAddressByQuery(queryParams map[string]string) ([]FixedAddress, error) {
	var ret []FixedAddress

	queryParams["_return_fields"] = fixedAddressReturnFields

	queryParamString := c.BuildQuery(queryParams)
	request, err := c.CreateJSONRequest(http.MethodGet, fmt.Sprintf("%s?%s", fixedAddressBasePath, queryParamString), nil)
	if err != nil {
		return ret, err
	}

	err = c.Call(request, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

// CreateFixedAddress creates fixed address
func (c *Client) CreateFixedAddress(fixedAddress *FixedAddress) error {
	queryParams := map[string]string{
		"_return_fields": fixedAddressReturnFields,
	}
	queryParamString := c.BuildQuery(queryParams)
	request, err := c.CreateJSONRequest(http.MethodPost, fmt.Sprintf("%s?%s", fixedAddressBasePath, queryParamString), fixedAddress)
	if err != nil {
		return err
	}

	err = c.Call(request, &fixedAddress)
	if err != nil {
		return err
	}
	return nil
}

// UpdateFixedAddress creates fixed address
func (c *Client) UpdateFixedAddress(ref string, fixedAddress FixedAddress) (FixedAddress, error) {
	var ret FixedAddress
	queryParams := map[string]string{
		"_return_fields": fixedAddressReturnFields,
	}
	queryParamString := c.BuildQuery(queryParams)
	request, err := c.CreateJSONRequest(http.MethodPut, fmt.Sprintf("%s?%s", ref, queryParamString), fixedAddress)
	if err != nil {
		return ret, err
	}

	err = c.Call(request, &ret)
	if err != nil {
		return ret, err
	}
	return ret, nil
}

// DeleteFixedAddress creates fixed address
func (c *Client) DeleteFixedAddress(ref string) error {
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
