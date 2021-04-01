package infoblox

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	eaDefintionBasePath = "extensibleattributedef"
)

// GetEADefinitions retrieves extensible attribute definitions
func (c *Client) GetEADefinitions(force bool) error {
	var ret []EADefinition

	if len(c.eaDefinitions) > 0 && force != false {
		return nil
	}
	queryParams := map[string]string{
		"_return_fields": "name,default_value,type,min,max,list_values",
	}
	queryParamString := c.BuildQuery(queryParams)
	request, err := c.CreateJSONRequest(http.MethodGet, fmt.Sprintf("%s?%s", eaDefintionBasePath, queryParamString), nil)
	if err != nil {
		return err
	}

	response := c.Call(request, &ret)
	if response != nil {
		return fmt.Errorf(response.ErrorMessage)
	}

	c.eaDefinitions = ret

	return nil
}

// ConvertEAsToJSONString converts extensible attributes to json format
func (c *Client) ConvertEAsToJSONString(eas ExtensibleAttribute) (map[string]string, error) {
	ret := make(map[string]string)
	if len(c.eaDefinitions) == 0 {
		c.GetEADefinitions(false)
	}
	for name, ea := range eas {
		var target EADefinition
		for _, def := range c.eaDefinitions {
			if def.Name == name {
				target = def
			}
		}
		if target.Ref == "" {
			return ret, fmt.Errorf("No ea definition found for ea: %s", name)
		}
		stringVal, _ := json.Marshal(ExtensibleAttributeJSONMapValue{
			Type:                 target.Type,
			Value:                ea.Value,
			InheritanceSource:    ea.InheritanceSource,
			InheritanceOperation: ea.InheritanceOperation,
			DescendantsAction:    ea.DescendantsAction,
		})
		ret[name] = string(stringVal)
	}
	return ret, nil
}

// ConvertEAsToJSONString converts extensible attributes to json format
// func (c *Client) ConvertJSONStringToEAs(eaString string) (*ExtensibleAttribute, error) {
// 	var ret ExtensibleAttribute
// 	var eaMap ExtensibleAttributeJSONMap
// 	json.Unmarshal([]byte(eaString), eaMap)
// 	for name, ea := range eaMap {
// 		switch ea.Type {
// 		case "STRING":
// 			ret[name] = ExtensibleAttributeValue{
// 				Value: ea.Value.(string),
// 			}
// 		}
// 	}
// 	preparedEAs, err := json.Marshal(ret)
// 	if err != nil {
// 		return "", err
// 	}
// 	return string(preparedEAs), nil
// }
