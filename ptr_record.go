package infoblox

import (
	"fmt"
	"net/http"
)

const (
	ptrRecordBasePath     = "record:ptr"
	ptrRecordReturnFields = "name,ptrdname,ipv4addr,ipv6addr,dns_name,dns_ptrdname,disable,view,dns_name,comment,zone,extattrs"
)

// GetPtrRecordByRef gets ptr record by reference
func (c *Client) GetPtrRecordByRef(ref string, queryParams map[string]string) (PtrRecord, error) {
	var ret PtrRecord
	if queryParams == nil {
		queryParams = map[string]string{
			"_return_fields": ptrRecordReturnFields,
		}
	} else {
		queryParams["_return_fields"] = ptrRecordReturnFields
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

// GetPtrRecordByQuery gets ptr records by query parameters
func (c *Client) GetPtrRecordByQuery(queryParams map[string]string) ([]PtrRecord, error) {
	var ret PtrRecordQueryResult
	queryParams["_return_fields"] = ptrRecordReturnFields
	queryParams["_return_as_object"] = "1"
	queryParams["_paging"] = "0"
	queryParams["_max_results"] = "2"

	queryParamString := c.BuildQuery(queryParams)
	request, err := c.CreateJSONRequest(http.MethodGet, fmt.Sprintf("%s?%s", ptrRecordBasePath, queryParamString), nil)
	if err != nil {
		return nil, err
	}

	response := c.Call(request, &ret)
	if response != nil {
		return nil, fmt.Errorf(response.ErrorMessage)
	}

	return ret.Results, nil
}

// CreatePtrRecord creates ptr record
func (c *Client) CreatePtrRecord(record *PtrRecord) error {
	queryParams := map[string]string{
		"_return_fields": ptrRecordReturnFields,
	}
	queryParamString := c.BuildQuery(queryParams)
	request, err := c.CreateJSONRequest(http.MethodPost, fmt.Sprintf("%s?%s", ptrRecordBasePath, queryParamString), record)
	if err != nil {
		return err
	}

	response := c.Call(request, &record)
	if response != nil {
		return fmt.Errorf(response.ErrorMessage)
	}
	return nil
}

// UpdatePtrRecord creates ptr record
func (c *Client) UpdatePtrRecord(ref string, network PtrRecord) (PtrRecord, error) {
	var ret PtrRecord
	queryParams := map[string]string{
		"_return_fields": ptrRecordReturnFields,
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

// DeletePtrRecord creates ptr record
func (c *Client) DeletePtrRecord(ref string) error {
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
