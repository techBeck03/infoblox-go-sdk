package infoblox

import (
	"fmt"
	"net/http"
)

const (
	fixedAddressBasePath = "fixedaddress"
)

// GetFixedAddressByRef gets host record by reference
func (c *Client) GetFixedAddressByRef(ref string) (FixedAddress, error) {
	var ret FixedAddress
	queryParams := map[string]string{
		"_return_fields": "extattrs,ipv4addr,network_view,comment,name",
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

// CreateFixedAddress creates host record
func (c *Client) CreateFixedAddress(fixedAddress *FixedAddress) error {
	queryParams := map[string]string{
		"_return_fields": "extattrs,ipv4addr,network_view,comment,name",
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

// UpdateFixedAddress creates host record
func (c *Client) UpdateFixedAddress(ref string, fixedAddress FixedAddress) (FixedAddress, error) {
	var ret FixedAddress
	queryParams := map[string]string{
		"_return_fields": "extattrs,ipv4addr,network_view,comment,name",
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

// DeleteFixedAddress creates host record
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
