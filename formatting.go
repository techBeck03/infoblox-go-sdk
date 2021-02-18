package infoblox

import (
	"encoding/json"
	"log"
)

func prettyPrint(object interface{}) {
	output, _ := json.MarshalIndent(object, "", "    ")
	log.Printf("%s", string(output))
}

func newExtensibleAttribute(ea ExtensibleAttribute) *ExtensibleAttribute {
	return &ea
}

func newBool(b bool) *bool {
	return &b
}
