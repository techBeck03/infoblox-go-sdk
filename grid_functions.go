package infoblox

import (
	"fmt"
	"net/http"
)

const (
	gridBasePath     = "grid"
	memberBasePath   = "member"
	gridReturnFields = "name,service_status,dns_resolver_setting"
)

// GetGrids gets grid member list
func (c *Client) GetGrids(queryParams map[string]string) ([]Grid, error) {
	var ret []Grid
	if queryParams == nil {
		queryParams = map[string]string{
			"_return_fields": gridReturnFields,
		}
	} else {
		queryParams["_return_fields"] = gridReturnFields
	}

	queryParamString := c.BuildQuery(queryParams)
	request, err := c.CreateJSONRequest(http.MethodGet, fmt.Sprintf("%s?%s", gridBasePath, queryParamString), nil)
	if err != nil {
		return ret, err
	}

	err = c.Call(request, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

// GetGridMembers gets grid member list
func (c *Client) GetGridMembers() ([]GridMember, error) {
	var ret []GridMember

	request, err := c.CreateJSONRequest(http.MethodGet, fmt.Sprintf("%s", memberBasePath), nil)
	if err != nil {
		return ret, err
	}

	err = c.Call(request, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

// RestartServices restarts selected grid services
func (c *Client) RestartServices(ref string, restartRequest GridServiceRestartRequest) error {
	queryParams := map[string]string{
		"_function": "restartservices",
	}

	queryParamString := c.BuildQuery(queryParams)
	request, err := c.CreateJSONRequest(http.MethodPost, fmt.Sprintf("%s?%s", ref, queryParamString), nil)
	if err != nil {
		return err
	}

	err = c.Call(request, nil)
	if err != nil {
		return err
	}

	return nil
}
