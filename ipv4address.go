package infoblox

import (
	"fmt"
	"net"
	"net/http"

	"github.com/techBeck03/go-ipmath"
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
	rangeMatchFlag := false

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
	if query.StartAddress != "" {
		queryParams["start_address>"] = query.StartAddress
	}
	if query.EndAddress != "" {
		queryParams["end_address<"] = query.StartAddress
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

	_, network, _ := net.ParseCIDR(query.CIDR)
	rangePage, err := c.GetPaginatedCidrRanges(query.CIDR, "")
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
		for endIndex <= resultsCount && matchFlag == false {
			var currentMatch ipmath.IP
			var lastMatch ipmath.IP

			if startIndex > endIndex {
				currentMatch = ipmath.IP{
					Address: net.ParseIP(prevPage[startIndex].IPAddress),
					Network: network,
				}
			} else {
				currentMatch = ipmath.IP{
					Address: net.ParseIP(ret.Results[startIndex].IPAddress),
					Network: network,
				}
			}
			lastMatch = ipmath.IP{
				Address: net.ParseIP(ret.Results[endIndex].IPAddress),
				Network: network,
			}

			if currentMatch.Difference(lastMatch.Address) == (query.Count - 1) {
				if len(rangePage.Results) > 0 {
					for rangeMatchFlag == false {
						for _, addressRange := range rangePage.Results {
							if ipWithinRange(addressRange.StartAddress, addressRange.EndAddress, currentMatch.Address.String()) || ipWithinRange(addressRange.StartAddress, addressRange.EndAddress, lastMatch.Address.String()) {
								rangeMatchFlag = true
								break
							}
						}
						if rangeMatchFlag == false && rangePage.NextPageID != "" {
							rangePage, err = c.GetPaginatedCidrRanges(query.CIDR, rangePage.NextPageID)
							if err != nil {
								return &addresses, err
							}
						} else if rangeMatchFlag == false && rangePage.NextPageID == "" {
							matchFlag = true
							break
						}
					}
				} else {
					matchFlag = true
				}
				if matchFlag == true {
					for i := 0; i <= query.Count-1; i++ {
						if startIndex > endIndex {
							addresses = append(addresses, prevPage[startIndex])
						} else {
							addresses = append(addresses, ret.Results[startIndex])
						}
						if len(prevPage) > 0 && startIndex == len(prevPage)-1 {
							startIndex = 0
						} else {
							startIndex++
						}
					}
				}
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
	if *query.FilterEmptyHostnames == true {
		var filteredResults []IPv4Address
		for _, result := range ret.Results {
			if len(result.Hostnames) > 0 {
				filteredResults = append(filteredResults, result)
			}
		}
		return &filteredResults, nil
	}
	return &ret.Results, nil
}
