package infoblox

import (
	"fmt"
	"net/http"
)

const (
	aliasRecordBasePath     = "record:a"
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

	err = c.Call(request, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

// GetAliasRecordByQuery gets alias records by query parameters
func (c *Client) GetAliasRecordByQuery(queryParams map[string]string) ([]AliasRecord, error) {
	var ret []AliasRecord
	queryParams["_return_fields"] = aliasRecordReturnFields

	queryParamString := c.BuildQuery(queryParams)
	request, err := c.CreateJSONRequest(http.MethodGet, fmt.Sprintf("%s?%s", aliasRecordBasePath, queryParamString), nil)
	if err != nil {
		return ret, err
	}

	err = c.Call(request, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

// CreateAliasRecord creates alias record
func (c *Client) CreateAliasRecord(network *AliasRecord) error {
	queryParams := map[string]string{
		"_return_fields": aliasRecordReturnFields,
	}
	queryParamString := c.BuildQuery(queryParams)
	request, err := c.CreateJSONRequest(http.MethodPost, fmt.Sprintf("%s?%s", aliasRecordBasePath, queryParamString), network)
	if err != nil {
		return err
	}

	err = c.Call(request, &network)
	if err != nil {
		return err
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

	err = c.Call(request, &ret)
	if err != nil {
		return ret, err
	}
	return ret, nil
}

// DeleteAliasRecord creates alias record
func (c *Client) DeleteAliasRecord(ref string) error {
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
