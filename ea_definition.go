package infoblox

import (
	"fmt"
	"net/http"
)

const (
	eaDefintionBasePath = "extensibleattributedef"
)

// GetEADefinitions retrieves extensible attribute definitions
func (c *Client) GetEADefinitions(ref string) ([]EADefinition, error) {
	var ret []EADefinition
	queryParams := map[string]string{
		"_return_fields": "name,default_value,type,min,max,list_values",
	}
	queryParamString := c.BuildQuery(queryParams)
	request, err := c.CreateJSONRequest(http.MethodGet, fmt.Sprintf("%s%s", eaDefintionBasePath, queryParamString), nil)
	if err != nil {
		return ret, err
	}

	err = c.Call(request, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}
