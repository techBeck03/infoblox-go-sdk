package infoblox

import (
	"fmt"
	"log"
	"net/http"
	"time"
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

// GetPaginatedCidrRanges gets ranges within CIDR by page
func (c *Client) GetPaginatedCidrRanges(cidr string, pageID string) (rangePage RangeQueryResult, err error) {
	var ret RangeQueryResult

	queryParams := map[string]string{
		"network":           cidr,
		"_return_as_object": "1",
		"_paging":           "1",
		"_max_results":      "100",
		"_return_fields":    rangeReturnFields,
	}
	if pageID != "" {
		queryParams["_page_id"] = pageID
	}

	queryParamString := c.BuildQuery(queryParams)
	request, err := c.CreateJSONRequest(http.MethodGet, fmt.Sprintf("%s?%s", rangeBasePath, queryParamString), nil)
	if err != nil {
		return rangePage, err
	}

	err = c.Call(request, &ret)
	if err != nil {
		return rangePage, err
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

// UpdateRange updates range
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

// DeleteRange deletes range
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

// CreateSequentialRange creates sequential address range
func (c *Client) CreateSequentialRange(rangeObject *Range, query AddressQuery) error {
	retryCount := 0
	verified := false

	for verified == false && retryCount <= query.Retries {
		sequentialAddresses, err := c.GetSequentialAddressRange(query)
		if err != nil {
			return err
		}
		rangeObject.StartAddress = (*sequentialAddresses)[0].IPAddress
		rangeObject.EndAddress = (*sequentialAddresses)[len(*sequentialAddresses)-1].IPAddress

		err = c.CreateRange(rangeObject)
		if err != nil {
			return err
		}

		time.Sleep(1 * time.Second)

		// Check for used addresses within range
		usedAddresses, err := c.GetUsedAddressesWithinRange(AddressQuery{
			CIDR:         query.CIDR,
			StartAddress: rangeObject.StartAddress,
			EndAddress:   rangeObject.EndAddress,
		})
		if err != nil {
			return err
		}
		if len((*usedAddresses)) > 0 {
			log.Println("Found allocated addresses within newly created range.  Deleting and Recreating.....")
			retryCount++
			err := c.DeleteRange(rangeObject.Ref)
			if err != nil {
				return err
			}
		} else {
			verified = true
		}
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
