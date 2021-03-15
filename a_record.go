package infoblox

import (
	"fmt"
	"net/http"
)

const (
	aRecordBasePath     = "record:a"
	aRecordReturnFields = "ipv4addr,name,view,dns_name,comment,zone,extattrs"
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

	err = c.Call(request, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

// GetARecordByQuery gets A records by query parameters
func (c *Client) GetARecordByQuery(queryParams map[string]string) ([]ARecord, error) {
	var ret []ARecord
	queryParams["_return_fields"] = aRecordReturnFields

	queryParamString := c.BuildQuery(queryParams)
	request, err := c.CreateJSONRequest(http.MethodGet, fmt.Sprintf("%s?%s", aRecordBasePath, queryParamString), nil)
	if err != nil {
		return ret, err
	}

	err = c.Call(request, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

// CreateARecord creates A record
func (c *Client) CreateARecord(network *ARecord) error {
	queryParams := map[string]string{
		"_return_fields": aRecordReturnFields,
	}
	queryParamString := c.BuildQuery(queryParams)
	request, err := c.CreateJSONRequest(http.MethodPost, fmt.Sprintf("%s?%s", aRecordBasePath, queryParamString), network)
	if err != nil {
		return err
	}

	err = c.Call(request, &network)
	if err != nil {
		return err
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

	err = c.Call(request, &ret)
	if err != nil {
		return ret, err
	}
	return ret, nil
}

// DeleteARecord creates A record
func (c *Client) DeleteARecord(ref string) error {
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
