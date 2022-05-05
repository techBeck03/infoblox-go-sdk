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
		queryParams["ip_address>"] = query.StartAddress
	}
	if query.EndAddress != "" {
		queryParams["ip_address<"] = query.EndAddress
	}
	queryParamString := c.BuildQuery(queryParams)

	request, err := c.CreateJSONRequest(http.MethodGet, fmt.Sprintf("%s?%s", ipv4AddressBasePath, queryParamString), nil)
	if err != nil {
		return &addresses, err
	}

	response := c.Call(request, &ret)
	if response != nil {
		return &addresses, fmt.Errorf(response.ErrorMessage)
	}

	_, network, _ := net.ParseCIDR(query.CIDR)
	rangePage, err := c.GetPaginatedCidrRanges(query.CIDR, "")
	if err != nil {
		return &addresses, err
	}
	for !matchFlag {
		resultsCount := len(ret.Results)
		if ret.NextPageID == "" && ((len(prevPage) == 0 && resultsCount < query.Count) || (len(prevPage) > 0 && ((len(prevPage)-startIndex)+resultsCount) < query.Count)) {
			return &addresses, fmt.Errorf("no sequential block found for supplied count")
		}
		if startIndex == -1 {
			startIndex = 0
			endIndex = query.Count - 1
		} else {
			endIndex = 0
		}
		for endIndex <= resultsCount && !matchFlag {
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
					for !rangeMatchFlag {
						for _, addressRange := range rangePage.Results {
							if ipWithinRange(addressRange.StartAddress, addressRange.EndAddress, currentMatch.Address.String()) || ipWithinRange(addressRange.StartAddress, addressRange.EndAddress, lastMatch.Address.String()) {
								rangeMatchFlag = true
								break
							}
						}
						if !rangeMatchFlag && rangePage.NextPageID != "" {
							rangePage, err = c.GetPaginatedCidrRanges(query.CIDR, rangePage.NextPageID)
							if err != nil {
								return &addresses, err
							}
						} else if !rangeMatchFlag && rangePage.NextPageID == "" {
							matchFlag = true
							break
						}
					}
				} else {
					matchFlag = true
				}
				if matchFlag {
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
		if !matchFlag && ret.NextPageID != "" {
			prevPage = ret.Results
			queryParams["_page_id"] = ret.NextPageID
			queryParamString := c.BuildQuery(queryParams)

			request, err := c.CreateJSONRequest(http.MethodGet, fmt.Sprintf("ipv4address%s", queryParamString), nil)
			if err != nil {
				return &addresses, err
			}

			response := c.Call(request, &ret)
			if response != nil {
				return &addresses, fmt.Errorf(response.ErrorMessage)
			}
		} else if !matchFlag && ret.NextPageID == "" {
			return &addresses, fmt.Errorf("no sequential block found for supplied count")
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
		"_return_as_object": "1",
		"ip_address>":       query.StartAddress,
		"ip_address<":       query.EndAddress,
		"_return_fields":    "ip_address,network,network_view,status,names,objects",
	}
	queryParamString := c.BuildQuery(queryParams)
	request, err := c.CreateJSONRequest(http.MethodGet, fmt.Sprintf("%s?%s", ipv4AddressBasePath, queryParamString), nil)
	if err != nil {
		return &addresses, err
	}

	response := c.Call(request, &ret)
	if response != nil {
		return &addresses, fmt.Errorf(response.ErrorMessage)
	}
	var filteredResults []IPv4Address
	if *query.FilterEmptyHostnames {
		for _, result := range ret.Results {
			if (len(result.Hostnames) > 0 || len(result.Objects) > 0) && result.Status == "USED" {
				filteredResults = append(filteredResults, result)
			}
		}
	} else {
		for _, result := range ret.Results {
			if result.Status == "USED" {
				filteredResults = append(filteredResults, result)
			}
		}
	}
	return &filteredResults, nil
}
