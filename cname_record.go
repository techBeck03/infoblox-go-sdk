package infoblox

import (
	"fmt"
	"net/http"
)

const (
	cNameRecordBasePath     = "record:cname"
	cNameRecordReturnFields = "name,canonical,view,dns_name,dns_canonical,disable,comment,zone,extattrs"
)

// GetCNameRecordByRef gets cname record by reference
func (c *Client) GetCNameRecordByRef(ref string, queryParams map[string]string) (CNameRecord, error) {
	var ret CNameRecord
	if queryParams == nil {
		queryParams = map[string]string{
			"_return_fields": cNameRecordReturnFields,
		}
	} else {
		queryParams["_return_fields"] = cNameRecordReturnFields
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

// GetCNameRecordByQuery gets cname records by query parameters
func (c *Client) GetCNameRecordByQuery(queryParams map[string]string) ([]CNameRecord, error) {
	var ret CNameRecordQueryResult
	queryParams["_return_fields"] = cNameRecordReturnFields
	queryParams["_return_as_object"] = "1"
	queryParams["_paging"] = "0"
	queryParams["_max_results"] = "2"

	queryParamString := c.BuildQuery(queryParams)
	request, err := c.CreateJSONRequest(http.MethodGet, fmt.Sprintf("%s?%s", cNameRecordBasePath, queryParamString), nil)
	if err != nil {
		return nil, err
	}

	err = c.Call(request, &ret)
	if err != nil {
		return nil, err
	}

	return ret.Results, nil
}

// CreateCNameRecord creates cname record
func (c *Client) CreateCNameRecord(network *CNameRecord) error {
	queryParams := map[string]string{
		"_return_fields": cNameRecordReturnFields,
	}
	queryParamString := c.BuildQuery(queryParams)
	request, err := c.CreateJSONRequest(http.MethodPost, fmt.Sprintf("%s?%s", cNameRecordBasePath, queryParamString), network)
	if err != nil {
		return err
	}

	err = c.Call(request, &network)
	if err != nil {
		return err
	}
	return nil
}

// UpdateCNameRecord creates cname record
func (c *Client) UpdateCNameRecord(ref string, network CNameRecord) (CNameRecord, error) {
	var ret CNameRecord
	queryParams := map[string]string{
		"_return_fields": cNameRecordReturnFields,
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

// DeleteCNameRecord creates cname record
func (c *Client) DeleteCNameRecord(ref string) error {
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
