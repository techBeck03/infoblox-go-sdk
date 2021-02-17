package infoblox

import (
	"fmt"
	"net/http"
)

const (
	hostRecordBasePath = "record:host"
)

// GetHostRecordByRef gets host record by reference
func (c *Client) GetHostRecordByRef(ref string) (HostRecord, error) {
	var ret HostRecord
	queryParams := map[string]string{
		"_return_fields": "ipv4addrs,extattrs,name,view",
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

// CreateHostRecord creates host record
func (c *Client) CreateHostRecord(hostRecord *HostRecord) error {
	queryParams := map[string]string{
		"_return_fields": "ipv4addrs,extattrs,name,view",
	}
	queryParamString := c.BuildQuery(queryParams)
	request, err := c.CreateJSONRequest(http.MethodPost, fmt.Sprintf("%s?%s", hostRecordBasePath, queryParamString), hostRecord)
	if err != nil {
		return err
	}

	err = c.Call(request, &hostRecord)
	if err != nil {
		return err
	}
	return nil
}

// UpdateHostRecord creates host record
func (c *Client) UpdateHostRecord(ref string, hostRecord HostRecord) (HostRecord, error) {
	var ret HostRecord
	queryParams := map[string]string{
		"_return_fields": "ipv4addrs,extattrs,name,view",
	}
	queryParamString := c.BuildQuery(queryParams)
	request, err := c.CreateJSONRequest(http.MethodPut, fmt.Sprintf("%s?%s", ref, queryParamString), hostRecord)
	if err != nil {
		return ret, err
	}

	err = c.Call(request, &ret)
	if err != nil {
		return ret, err
	}
	return ret, nil
}

// DeleteHostRecord creates host record
func (c *Client) DeleteHostRecord(ref string) error {
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
