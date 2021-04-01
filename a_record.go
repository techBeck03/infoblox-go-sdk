package infoblox

import (
	"fmt"
	"net/http"
)

const (
	aRecordBasePath     = "record:a"
	aRecordReturnFields = "ipv4addr,name,view,dns_name,disable,comment,zone,extattrs"
)

// GetARecordByRef gets A record by reference
func (c *Client) GetARecordByRef(ref string, queryParams map[string]string) (ARecord, error) {
	var ret ARecord
	if queryParams == nil {
		queryParams = map[string]string{
			"_return_fields": aRecordReturnFields,
		}
	} else {
		queryParams["_return_fields"] = aRecordReturnFields
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

// GetARecordByQuery gets A records by query parameters
func (c *Client) GetARecordByQuery(queryParams map[string]string) ([]ARecord, error) {
	var ret ARecordQueryResult
	queryParams["_return_fields"] = aRecordReturnFields
	queryParams["_return_as_object"] = "1"
	queryParams["_paging"] = "0"
	queryParams["_max_results"] = "2"

	queryParamString := c.BuildQuery(queryParams)
	request, err := c.CreateJSONRequest(http.MethodGet, fmt.Sprintf("%s?%s", aRecordBasePath, queryParamString), nil)
	if err != nil {
		return nil, err
	}

	response := c.Call(request, &ret)
	if response != nil {
		return nil, fmt.Errorf(response.ErrorMessage)
	}

	return ret.Results, nil
}

// CreateARecord creates A record
func (c *Client) CreateARecord(record *ARecord) error {
	queryParams := map[string]string{
		"_return_fields": aRecordReturnFields,
	}
	queryParamString := c.BuildQuery(queryParams)
	request, err := c.CreateJSONRequest(http.MethodPost, fmt.Sprintf("%s?%s", aRecordBasePath, queryParamString), record)
	if err != nil {
		return err
	}

	response := c.Call(request, &record)
	if response != nil {
		return fmt.Errorf(response.ErrorMessage)
	}
	return nil
}

// UpdateARecord creates A record
func (c *Client) UpdateARecord(ref string, network ARecord) (ARecord, error) {
	var ret ARecord
	queryParams := map[string]string{
		"_return_fields": aRecordReturnFields,
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

// DeleteARecord creates A record
func (c *Client) DeleteARecord(ref string) error {
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
