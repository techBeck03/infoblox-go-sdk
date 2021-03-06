package infoblox

import (
	"fmt"
	"net/http"
)

const (
	aliasRecordBasePath     = "record:alias"
	aliasRecordReturnFields = "name,target_name,target_type,dns_name,dns_target_name,disable,view,dns_name,comment,zone,extattrs"
)

// GetAliasRecordByRef gets alias record by reference
func (c *Client) GetAliasRecordByRef(ref string, queryParams map[string]string) (AliasRecord, error) {
	var ret AliasRecord
	if queryParams == nil {
		queryParams = map[string]string{
			"_return_fields": aliasRecordReturnFields,
		}
	} else {
		queryParams["_return_fields"] = aliasRecordReturnFields
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

// GetAliasRecordByQuery gets alias records by query parameters
func (c *Client) GetAliasRecordByQuery(queryParams map[string]string) ([]AliasRecord, error) {
	var ret AliasRecordQueryResult
	queryParams["_return_fields"] = aliasRecordReturnFields
	queryParams["_return_as_object"] = "1"
	queryParams["_paging"] = "0"
	queryParams["_max_results"] = "2"

	queryParamString := c.BuildQuery(queryParams)
	request, err := c.CreateJSONRequest(http.MethodGet, fmt.Sprintf("%s?%s", aliasRecordBasePath, queryParamString), nil)
	if err != nil {
		return nil, err
	}

	response := c.Call(request, &ret)
	if response != nil {
		return nil, fmt.Errorf(response.ErrorMessage)
	}

	return ret.Results, nil
}

// CreateAliasRecord creates alias record
func (c *Client) CreateAliasRecord(record *AliasRecord) error {
	queryParams := map[string]string{
		"_return_fields": aliasRecordReturnFields,
	}
	queryParamString := c.BuildQuery(queryParams)
	request, err := c.CreateJSONRequest(http.MethodPost, fmt.Sprintf("%s?%s", aliasRecordBasePath, queryParamString), record)
	if err != nil {
		return err
	}

	response := c.Call(request, &record)
	if response != nil {
		return fmt.Errorf(response.ErrorMessage)
	}
	return nil
}

// UpdateAliasRecord creates alias record
func (c *Client) UpdateAliasRecord(ref string, network AliasRecord) (AliasRecord, error) {
	var ret AliasRecord
	queryParams := map[string]string{
		"_return_fields": aliasRecordReturnFields,
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

// DeleteAliasRecord creates alias record
func (c *Client) DeleteAliasRecord(ref string) error {
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
