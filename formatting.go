package infoblox

import (
	"bytes"
	"encoding/json"
	"log"
	"net"
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

func ipWithinRange(startAddress string, endAddress string, ip string) bool {
	trial := net.ParseIP(ip)
	if trial.To4() == nil {
		return false
	}
	if bytes.Compare(trial, net.ParseIP(startAddress)) >= 0 && bytes.Compare(trial, net.ParseIP(endAddress)) <= 0 {
		return true
	}
	return false
}
