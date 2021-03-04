package infoblox

import (
	"fmt"
	"net/http"
)

const (
	rangeBasePath     = "range"
	rangeReturnFields = "network,network_view,start_addr,end_addr,disable,comment,extattrs"
)

// GetRangeByRef gets range by reference
func (c *Client) GetRangeByRef(ref string, queryParams map[string]string) (Range, error) {
	var ret Range

	if queryParams == nil {
		queryParams = map[string]string{
			"_return_fields": rangeReturnFields,
		}
	} else {
		queryParams["_return_fields"] = rangeReturnFields
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

// GetRangeByQuery gets range by reference
func (c *Client) GetRangeByQuery(queryParams map[string]string) (Range, error) {
	var ret Range

	queryParams["_return_fields"] = rangeReturnFields

	queryParamString := c.BuildQuery(queryParams)
	request, err := c.CreateJSONRequest(http.MethodGet, fmt.Sprintf("%s?%s", rangeBasePath, queryParamString), nil)
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
		"_return_fields": rangeReturnFields,
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
		"_return_fields": rangeReturnFields,
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

// CheckIfRangeContainsRange checks if a range exists containing ip range
func (c *Client) CheckIfRangeContainsRange(query IPsWithinRangeQuery) (bool, error) {
	var ret RangeQueryResult

	queryParams := map[string]string{
		"network":           query.CIDR,
		"_return_as_object": "1",
		"_paging":           "1",
		"_max_results":      "100",
		"_return_fields":    rangeReturnFields,
	}
	queryParamString := c.BuildQuery(queryParams)

	request, err := c.CreateJSONRequest(http.MethodGet, fmt.Sprintf("%s?%s", rangeBasePath, queryParamString), nil)

	if err != nil {
		return true, err
	}

	err = c.Call(request, &ret)
	if err != nil {
		return true, err
	}

	if len(ret.Results) == 0 {
		return false, nil
	}

	matchFlag := false

	for matchFlag == false {
		for _, addressRange := range ret.Results {
			if addressRange.Ref != query.Ref && (ipWithinRange(addressRange.StartAddress, addressRange.EndAddress, query.StartAddress) || ipWithinRange(addressRange.StartAddress, addressRange.EndAddress, query.EndAddress)) {
				matchFlag = true
				break
			}
		}
		if matchFlag == false && ret.NextPageID != "" {
			queryParams["_page_id"] = ret.NextPageID
			queryParamString := c.BuildQuery(queryParams)

			request, err := c.CreateJSONRequest(http.MethodGet, fmt.Sprintf("%s?%s", rangeBasePath, queryParamString), nil)
			if err != nil {
				return true, err
			}

			err = c.Call(request, &ret)
			if err != nil {
				return true, err
			}
		} else if matchFlag == false && ret.NextPageID == "" {
			return false, nil
		}
	}

	return true, nil
}
