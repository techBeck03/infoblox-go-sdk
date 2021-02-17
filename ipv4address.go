package infoblox

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

const (
	ipv4AddressBasePath = "ipv4address"
)

// GetSequentialAddressRange retrieves count number of sequential IPs from supplied network
func (c *Client) GetSequentialAddressRange(query AddressQuery) (*[]IPv4Address, error) {
	var addresses []IPv4Address
	var ret AddressQueryResult
	var prevPage []IPv4Address
	startIndex := -1
	var endIndex int
	matchFlag := false

	query.fillDefaults()
	queryParams := map[string]string{
		"network":           query.CIDR,
		"network_view":      query.NetworkView,
		"status":            "UNUSED",
		"_return_as_object": "1",
		"_paging":           "1",
		"_max_results":      "100",
		"_return_fields":    "ip_address,network,network_view,status",
	}
	queryParamString := c.BuildQuery(queryParams)

	request, err := c.CreateJSONRequest(http.MethodGet, fmt.Sprintf("%s?%s", ipv4AddressBasePath, queryParamString), nil)
	if err != nil {
		return &addresses, err
	}

	err = c.Call(request, &ret)
	if err != nil {
		return &addresses, err
	}

	for matchFlag == false {
		resultsCount := len(ret.Results)
		if resultsCount < query.Count {
			return &addresses, fmt.Errorf("No sequential block found for supplied count")
		}
		if startIndex == -1 {
			startIndex = 0
			endIndex = query.Count - 1
		} else {
			endIndex = 0
		}
		for endIndex <= resultsCount {
			currentMatch := strings.Split(ret.Results[startIndex].IPAddress, ".")
			lastMatch := strings.Split(ret.Results[endIndex].IPAddress, ".")
			a, _ := strconv.Atoi(currentMatch[3])
			b, _ := strconv.Atoi(lastMatch[3])
			if a+query.Count > 255 {
				b = b + 256
			}
			if b-a == (query.Count - 1) {
				matchFlag = true
				for i := 0; i <= query.Count-1; i++ {
					addresses = append(addresses, ret.Results[startIndex])
					if len(prevPage) > 0 && startIndex == len(prevPage)-1 {
						startIndex = 0
					} else {
						startIndex++
					}
				}
				break
			}
			if len(prevPage) > 0 && startIndex == len(prevPage)-1 {
				startIndex = 0
			} else {
				startIndex++
			}
			endIndex++
		}
		if matchFlag == false && ret.NextPageID != "" {
			prevPage = ret.Results
			queryParams["_page_id"] = ret.NextPageID
			queryParamString := c.BuildQuery(queryParams)

			request, err := c.CreateJSONRequest(http.MethodGet, fmt.Sprintf("ipv4address%s", queryParamString), nil)
			if err != nil {
				return &addresses, err
			}

			err = c.Call(request, &ret)
			if err != nil {
				return &addresses, err
			}
		} else if matchFlag == false && ret.NextPageID == "" {
			return &addresses, fmt.Errorf("No sequential block found for supplied count")
		}
	}
	return &addresses, nil
}

// GetUsedAddressesWithinRange gets used addresses within selected network range
func (c *Client) GetUsedAddressesWithinRange(query AddressQuery) (*[]IPv4Address, error) {
	var addresses []IPv4Address
	var ret AddressQueryResult

	query.fillDefaults()
	queryParams := map[string]string{
		"network":           query.CIDR,
		"network_view":      query.NetworkView,
		"status":            "USED",
		"_return_as_object": "1",
		"ip_address>":       query.StartAddress,
		"ip_address<":       query.EndAddress,
		"_return_fields":    "ip_address,network,network_view,status",
	}
	queryParamString := c.BuildQuery(queryParams)
	request, err := c.CreateJSONRequest(http.MethodGet, fmt.Sprintf("%s?%s", ipv4AddressBasePath, queryParamString), nil)
	if err != nil {
		return &addresses, err
	}

	err = c.Call(request, &ret)
	if err != nil {
		return &addresses, err
	}
	return &ret.Results, nil
}
