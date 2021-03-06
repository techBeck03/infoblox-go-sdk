package infoblox

import (
	"fmt"
	"net/http"
)

const (
	hostRecordBasePath     = "record:host"
	hostRecordReturnFields = "name,view,network_view,configure_for_dns,comment,zone,ipv4addrs,ipv4addrs.host,ipv4addrs.network,ipv4addrs.ipv4addr,ipv4addrs.mac,ipv4addrs.configure_for_dhcp,ipv4addrs.nextserver,ipv4addrs.use_for_ea_inheritance,extattrs"
)

// GetHostRecordByRef gets host record by reference
func (c *Client) GetHostRecordByRef(ref string, queryParams map[string]string) (HostRecord, error) {
	var ret HostRecord

	if queryParams == nil {
		queryParams = map[string]string{
			"_return_fields": hostRecordReturnFields,
		}
	} else {
		queryParams["_return_fields"] = hostRecordReturnFields
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

// GetHostRecordByQuery gets host record by query parameters
func (c *Client) GetHostRecordByQuery(queryParams map[string]string) ([]HostRecord, error) {
	var ret HostRecordQueryResult
	queryParams["_return_fields"] = hostRecordReturnFields
	queryParams["_return_as_object"] = "1"
	queryParams["_paging"] = "0"
	queryParams["_max_results"] = "2"

	queryParamString := c.BuildQuery(queryParams)
	request, err := c.CreateJSONRequest(http.MethodGet, fmt.Sprintf("%s?%s", hostRecordBasePath, queryParamString), nil)
	if err != nil {
		return nil, err
	}

	response := c.Call(request, &ret)
	if response != nil {
		return nil, fmt.Errorf(response.ErrorMessage)
	}

	return ret.Results, nil
}

// CreateHostRecord creates host record
func (c *Client) CreateHostRecord(hostRecord *HostRecord) error {
	queryParams := map[string]string{
		"_return_fields": hostRecordReturnFields,
	}
	queryParamString := c.BuildQuery(queryParams)
	request, err := c.CreateJSONRequest(http.MethodPost, fmt.Sprintf("%s?%s", hostRecordBasePath, queryParamString), hostRecord)
	if err != nil {
		return err
	}

	response := c.Call(request, &hostRecord)
	if response != nil {
		return fmt.Errorf(response.ErrorMessage)
	}
	return nil
}

// UpdateHostRecord creates host record
func (c *Client) UpdateHostRecord(ref string, hostRecord HostRecord) (HostRecord, error) {
	var ret HostRecord
	queryParams := map[string]string{
		"_return_fields": hostRecordReturnFields,
	}
	queryParamString := c.BuildQuery(queryParams)
	request, err := c.CreateJSONRequest(http.MethodPut, fmt.Sprintf("%s?%s", ref, queryParamString), hostRecord)
	if err != nil {
		return ret, err
	}

	response := c.Call(request, &ret)
	if response != nil {
		return ret, fmt.Errorf(response.ErrorMessage)
	}
	return ret, nil
}

// DeleteHostRecord creates host record
func (c *Client) DeleteHostRecord(ref string) error {
	request, err := c.CreateJSONRequest(http.MethodDelete, fmt.Sprintf("%s", ref), nil)
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
