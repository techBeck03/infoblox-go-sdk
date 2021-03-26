package infoblox

import (
	"fmt"
	"net/http"
)

const (
	gridBasePath       = "grid"
	memberBasePath     = "member"
	gridReturnFields   = "name,service_status,dns_resolver_setting"
	memberReturnFields = "config_addr_type,host_name,platform,service_type_configuration"
)

// GetGridByRef gets grid by ref
func (c *Client) GetGridByRef(ref string) (Grid, error) {
	var ret Grid

	queryParams := map[string]string{
		"_return_fields": gridReturnFields,
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

// GetGridsByQuery gets grid list
func (c *Client) GetGridsByQuery(queryParams map[string]string) ([]Grid, error) {
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

// GetGridMembersByRef gets grid member list
func (c *Client) GetGridMembersByRef(ref string) (GridMember, error) {
	var ret GridMember

	queryParams := map[string]string{
		"_return_fields": memberReturnFields,
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

// GetGridMembersByQuery gets grid member list
func (c *Client) GetGridMembersByQuery(queryParams map[string]string) ([]GridMember, error) {
	var ret []GridMember

	if queryParams == nil {
		queryParams = map[string]string{
			"_return_fields": memberReturnFields,
		}
	} else {
		queryParams["_return_fields"] = memberReturnFields
	}

	queryParamString := c.BuildQuery(queryParams)
	request, err := c.CreateJSONRequest(http.MethodGet, fmt.Sprintf("%s?%s", memberBasePath, queryParamString), nil)
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
