package infoblox

import (
	"fmt"
	"net/http"
)

const (
	rangeBasePath = "range"
)

// GetRangeByRef gets range by reference
func (c *Client) GetRangeByRef(ref string) (Range, error) {
	var ret Range
	queryParams := map[string]string{
		"_return_fields": "network,network_view,start_addr,end_addr,disable,comment,extattrs",
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

// CreateRange creates range
func (c *Client) CreateRange(rangeObject *Range) error {
	queryParams := map[string]string{
		"_return_fields": "network,network_view,start_addr,end_addr,disable,comment,extattrs",
	}
	queryParamString := c.BuildQuery(queryParams)
	request, err := c.CreateJSONRequest(http.MethodPost, fmt.Sprintf("%s?%s", rangeBasePath, queryParamString), rangeObject)
	if err != nil {
		return err
	}

	err = c.Call(request, &rangeObject)
	if err != nil {
		return err
	}
	return nil
}

// UpdateRange creates range
func (c *Client) UpdateRange(ref string, rangeObject Range) (Range, error) {
	var ret Range
	queryParams := map[string]string{
		"_return_fields": "network,network_view,start_addr,end_addr,disable,comment,extattrs",
	}
	queryParamString := c.BuildQuery(queryParams)
	request, err := c.CreateJSONRequest(http.MethodPut, fmt.Sprintf("%s?%s", ref, queryParamString), rangeObject)
	if err != nil {
		return ret, err
	}

	err = c.Call(request, &ret)
	if err != nil {
		return ret, err
	}
	return ret, nil
}

// DeleteRange creates range
func (c *Client) DeleteRange(ref string) error {
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
