package infoblox

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/techBeck03/go-ipmath"
)

const (
	rangeBasePath     = "range"
	rangeReturnFields = "network,network_view,start_addr,end_addr,disable,comment,extattrs,member"
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

	response := c.Call(request, &ret)
	if response != nil {
		return ret, fmt.Errorf(response.ErrorMessage)
	}

	startingIP := ipmath.IP{
		Address: net.ParseIP(ret.StartAddress),
	}
	count := startingIP.Difference(net.ParseIP(ret.EndAddress)) + 1
	ret.IPAddressList = getRangeAddressList(ret.StartAddress, count)

	return ret, nil
}

// GetRangeByQuery gets range by query
func (c *Client) GetRangeByQuery(queryParams map[string]string) ([]Range, error) {
	var ret RangeQueryResult

	queryParams["_return_fields"] = rangeReturnFields
	queryParams["_return_as_object"] = "1"
	queryParams["_paging"] = "0"
	queryParams["_max_results"] = "2"

	queryParamString := c.BuildQuery(queryParams)
	request, err := c.CreateJSONRequest(http.MethodGet, fmt.Sprintf("%s?%s", rangeBasePath, queryParamString), nil)
	if err != nil {
		return nil, err
	}

	response := c.Call(request, &ret)
	if response != nil {
		return nil, fmt.Errorf(response.ErrorMessage)
	}

	for i, r := range ret.Results {
		startingIP := ipmath.IP{
			Address: net.ParseIP(r.StartAddress),
		}
		count := startingIP.Difference(net.ParseIP(r.EndAddress)) + 1
		ret.Results[i].IPAddressList = getRangeAddressList(r.StartAddress, count)
	}

	return ret.Results, nil
}

func getRangeAddressList(startAddress string, count int) []string {
	ipAddressList := []string{}
	startingIP := ipmath.IP{
		Address: net.ParseIP(startAddress),
	}
	for i := 0; i < count; i++ {
		ipAddressList = append(ipAddressList, startingIP.ToIPString())
		startingIP.Inc()
	}
	return ipAddressList
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

	response := c.Call(request, &ret)
	if response != nil {
		return rangePage, fmt.Errorf(response.ErrorMessage)
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

	response := c.Call(request, &rangeObject)
	if response != nil {
		return fmt.Errorf(response.ErrorMessage)
	}
	startingIP := ipmath.IP{
		Address: net.ParseIP(rangeObject.StartAddress),
	}
	count := startingIP.Difference(net.ParseIP(rangeObject.EndAddress)) + 1
	rangeObject.IPAddressList = getRangeAddressList(rangeObject.StartAddress, count)
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

	response := c.Call(request, &ret)
	if response != nil {
		return ret, fmt.Errorf(response.ErrorMessage)
	}
	return ret, nil
}

// DeleteRange deletes range
func (c *Client) DeleteRange(ref string) error {
	request, err := c.CreateJSONRequest(http.MethodDelete, ref, nil)
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

// CreateSequentialRange creates sequential address range
func (c *Client) CreateSequentialRange(rangeObject *Range, query AddressQuery) error {
	c.SequentialLock.Lock()
	defer c.SequentialLock.Unlock()
	query.fillDefaults()
	retryCount := 0
	verified := false
	for !verified && retryCount <= query.Retries {
		prettyPrint(query)
		log.Println("Getting sequential range")
		sequentialAddresses, err := c.GetSequentialAddressRange(query)
		if err != nil {
			log.Printf("The following error occurred while getting sequential address range: %s", err)
			return err
		}
		rangeObject.StartAddress = (*sequentialAddresses)[0].IPAddress
		rangeObject.EndAddress = (*sequentialAddresses)[len(*sequentialAddresses)-1].IPAddress

		log.Println("Creating range")
		err = c.CreateRange(rangeObject)
		if err != nil {
			verified = false
			time.Sleep(2 * time.Second)
			retryCount++
			log.Printf("An error occurred creating range: %s", err)
			prettyPrint(rangeObject)
		} else {
			log.Println("Pausing for race condition checks")
			time.Sleep(1 * time.Second)

			// Check for used addresses within range
			usedAddresses, err := c.GetUsedAddressesWithinRange(AddressQuery{
				CIDR:                 query.CIDR,
				StartAddress:         rangeObject.StartAddress,
				EndAddress:           rangeObject.EndAddress,
				FilterEmptyHostnames: newBool(true),
			})
			if err != nil {
				return err
			}
			if len((*usedAddresses)) > 0 {
				prettyPrint(rangeObject)
				log.Println("Found allocated addresses within newly created range.  Deleting and Recreating.....")
				retryCount++
				err := c.DeleteRange(rangeObject.Ref)
				if err != nil {
					log.Printf("The following error occurred deleting range: %s", err)
					return err
				}
				time.Sleep(1 * time.Second)
			} else {
				verified = true
			}
		}
	}

	if !verified {
		return fmt.Errorf("unable to create sequential range within %s", query.CIDR)
	}
	rangeObject.IPAddressList = getRangeAddressList(rangeObject.StartAddress, query.Count)
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

	response := c.Call(request, &ret)
	if response != nil {
		return true, fmt.Errorf(response.ErrorMessage)
	}

	if len(ret.Results) == 0 {
		return false, nil
	}

	matchFlag := false

	for !matchFlag {
		for _, addressRange := range ret.Results {
			if addressRange.Ref != query.Ref && (ipWithinRange(addressRange.StartAddress, addressRange.EndAddress, query.StartAddress) || ipWithinRange(addressRange.StartAddress, addressRange.EndAddress, query.EndAddress)) {
				matchFlag = true
				break
			}
		}
		if !matchFlag && ret.NextPageID != "" {
			queryParams["_page_id"] = ret.NextPageID
			queryParamString := c.BuildQuery(queryParams)

			request, err := c.CreateJSONRequest(http.MethodGet, fmt.Sprintf("%s?%s", rangeBasePath, queryParamString), nil)
			if err != nil {
				return true, err
			}

			response := c.Call(request, &ret)
			if response != nil {
				return true, fmt.Errorf(response.ErrorMessage)
			}
		} else if !matchFlag && ret.NextPageID == "" {
			return false, nil
		}
	}

	return true, nil
}
