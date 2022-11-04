package infoblox

// Grid defines grid properties
type Grid struct {
	Ref                string             `json:"_ref,omitempty"`
	Name               string             `json:"name,omitempty"`
	ServiceStatus      string             `json:"service_status,omitempty"`
	DNSResolverSetting DNSResolverSetting `json:"dns_resolver_setting,omitempty"`
}

// DNSResolverSetting defines grid dns resolver configuration
type DNSResolverSetting struct {
	Resolvers     []string `json:"resolvers,omitempty"`
	SearchDomains []string `json:"search_domains,omitempty"`
}

// GridMember defines grid member properties
type GridMember struct {
	Ref                      string `json:"_ref,omitempty"`
	Hostname                 string `json:"host_name,omitempty"`
	ConfigAddressType        string `json:"config_addr_type,omitempty"`
	Platform                 string `json:"platform,omitempty"`
	ServiceTypeConfiguration string `json:"service_type_configuration,omitempty"`
}

// GridServiceRestartRequest defines properties for grid restart request
type GridServiceRestartRequest struct {
	RestartOption string   `json:"restart_option,omitempty"`
	Services      []string `json:"services,omitempty"`
	Members       []string `json:"members,omitempty"`
}

// ExtensibleAttribute extensible attribute object
type ExtensibleAttribute map[string]ExtensibleAttributeValue

// ExtensibleAttributeValue return value of ea
type ExtensibleAttributeValue struct {
	Value                interface{}        `json:"value,omitempty"`
	InheritanceSource    *InheritanceSource `json:"inheritance_source,omitempty"`
	InheritanceOperation string             `json:"inheritance_operation,omitempty"`
	DescendantsAction    *DescendantsAction `json:"descendants_action,omitempty"`
}

// InheritanceSource defines inheritance of an EA
type InheritanceSource struct {
	Ref string `json:"_ref,omitempty"`
}

// DescendantsAction defines inheritance of an EA
type DescendantsAction struct {
	OptionDeleteEA  string `json:"option_delete_ea,omitempty"`
	OptionWithEA    string `json:"option_with_ea,omitempty"`
	OptionWithoutEA string `json:"option_without_ea,omitempty"`
}

// ExtensibleAttributeJSONMap ea object in terraform friendly JSON
type ExtensibleAttributeJSONMap map[string]ExtensibleAttributeJSONMapValue

// ExtensibleAttributeJSONMapValue value of ea in terraform friendly JSON
type ExtensibleAttributeJSONMapValue struct {
	Value                interface{}        `json:"value,omitempty"`
	Type                 string             `json:"type,omitempty"`
	InheritanceSource    *InheritanceSource `json:"inheritance_source,omitempty"`
	InheritanceOperation string             `json:"inheritance_operation,omitempty"`
	DescendantsAction    *DescendantsAction `json:"descendants_action,omitempty"`
}

// Network object
type Network struct {
	Ref                        string               `json:"_ref,omitempty"`
	NetworkView                string               `json:"network_view,omitempty"`
	CIDR                       string               `json:"network,omitempty"`
	Comment                    string               `json:"comment,omitempty"`
	DisableDHCP                *bool                `json:"disable,omitempty"`
	Members                    []Member             `json:"members,omitempty"`
	Options                    []Option             `json:"options,omitempty"`
	ExtensibleAttributes       *ExtensibleAttribute `json:"extattrs,omitempty"`
	ExtensibleAttributesAdd    *ExtensibleAttribute `json:"extattrs+,omitempty"`
	ExtensibleAttributesRemove *ExtensibleAttribute `json:"extattrs-,omitempty"`
}

// NetworkFromContainer object
type NetworkFromContainer struct {
	Ref                        string                   `json:"_ref,omitempty"`
	NetworkView                string                   `json:"network_view,omitempty"`
	Network                    NetworkContainerFunction `json:"network,omitempty"`
	Comment                    string                   `json:"comment,omitempty"`
	DisableDHCP                *bool                    `json:"disable,omitempty"`
	Members                    []Member                 `json:"members,omitempty"`
	Options                    []Option                 `json:"options,omitempty"`
	ExtensibleAttributes       *ExtensibleAttribute     `json:"extattrs,omitempty"`
	ExtensibleAttributesAdd    *ExtensibleAttribute     `json:"extattrs+,omitempty"`
	ExtensibleAttributesRemove *ExtensibleAttribute     `json:"extattrs-,omitempty"`
}

// NetworkContainerFunction object
type NetworkContainerFunction struct {
	Function         string                           `json:"_object_function,omitempty"`
	ResultField      string                           `json:"_result_field,omitempty"`
	Object           string                           `json:"_object,omitempty"`
	ObjectParameters *NetworkContainerObjectParameter `json:"_object_parameters,omitempty"`
	Parameters       *NetworkContainerParameter       `json:"_parameters,omitempty"`
}

// NetworkContainerObjectParameter object
type NetworkContainerObjectParameter struct {
	Label string `json:"*Label,omitempty"`
}

// NetworkContainerParameter object
type NetworkContainerParameter struct {
	Prefix int `json:"cidr,omitempty"`
}

// NetworkQueryResult object
type NetworkQueryResult struct {
	NextPageID string    `json:"next_page_id,omitempty"`
	Results    []Network `json:"result,omitempty"`
}

// Member defines grid members
type Member struct {
	StructType  string `json:"_struct,omitempty"`
	Hostname    string `json:"name,omitempty"`
	IPV4Address string `json:"ipv4addr,omitempty"`
	IPV6Address string `json:"ipv6addr,omitempty"`
}

// Option defines dhcp options
type Option struct {
	Name        string `json:"name,omitempty"`
	Code        int    `json:"num,omitempty"`
	UseOption   *bool  `json:"use_option,omitempty"`
	Value       string `json:"value,omitempty"`
	VendorClass string `json:"vendor_class,omitempty"`
}

// EADefinition extensible attribute definition
type EADefinition struct {
	Ref                string      `json:"_ref,omitempty"`
	AllowedObjectTypes string      `json:"allowed_object_types,omitempty"`
	Comment            string      `json:"comment,omitempty"`
	DefaultValue       string      `json:"default_value,omitempty"`
	DescendantsAction  string      `json:"descendants_action,omitempty"`
	Flags              string      `json:"flags,omitempty"`
	ListValues         []ListValue `json:"list_values,omitempty"`
	Max                string      `json:"max,omitempty"`
	Min                string      `json:"min,omitempty"`
	Name               string      `json:"name,omitempty"`
	Namespace          string      `json:"namespace,omitempty"`
	Type               string      `json:"type,omitempty"`
}

// ListValue defines possible list values
type ListValue struct {
	Value string `json:"value,omitempty"`
}

// HostRecord object
type HostRecord struct {
	Ref                        string               `json:"_ref,omitempty"`
	Hostname                   string               `json:"name,omitempty"`
	Comment                    string               `json:"comment,omitempty"`
	EnableDNS                  *bool                `json:"configure_for_dns,omitempty"`
	IPv4Addrs                  []IPv4Addr           `json:"ipv4addrs,omitempty"`
	NetworkView                string               `json:"network_view,omitempty"`
	RestartIfNeeded            *bool                `json:"restart_if_needed,omitempty"`
	View                       string               `json:"view,omitempty"`
	Zone                       string               `json:"zone,omitempty"`
	ExtensibleAttributes       *ExtensibleAttribute `json:"extattrs,omitempty"`
	ExtensibleAttributesAdd    *ExtensibleAttribute `json:"extattrs+,omitempty"`
	ExtensibleAttributesRemove *ExtensibleAttribute `json:"extattrs-,omitempty"`
}

// HostRecordQueryResult object
type HostRecordQueryResult struct {
	NextPageID string       `json:"next_page_id,omitempty"`
	Results    []HostRecord `json:"result,omitempty"`
}

// IPv4Addr object
type IPv4Addr struct {
	Ref                 string                 `json:"_ref,omitempty"`
	Host                string                 `json:"host,omitempty"`
	IPAddress           string                 `json:"ipv4addr,omitempty"`
	Mac                 string                 `json:"mac,omitempty"`
	CIDR                string                 `json:"network,omitempty"`
	ConfigureForDHCP    *bool                  `json:"configure_for_dhcp,omitempty"`
	NextServer          string                 `json:"nextserver,omitempty"`
	ObjectFunction      string                 `json:"_object_function,omitempty"`
	UseForEAInheritance *bool                  `json:"use_for_ea_inheritance,omitempty"`
	Parameters          map[string]interface{} `json:"_parameters,omitempty"`
	ResultField         string                 `json:"_result_field,omitempty"`
	Object              string                 `json:"_object,omitempty"`
	ObjectParameters    map[string]interface{} `json:"_object_parameters,omitempty"`
}

// FixedAddress object
type FixedAddress struct {
	Ref                        string               `json:"_ref,omitempty"`
	NetworkView                string               `json:"network_view,omitempty"`
	CIDR                       string               `json:"network,omitempty"`
	Comment                    string               `json:"comment,omitempty"`
	Disable                    *bool                `json:"disable,omitempty"`
	IPAddress                  string               `json:"ipv4addr,omitempty"`
	Mac                        string               `json:"mac,omitempty"`
	Hostname                   string               `json:"name,omitempty"`
	MatchClient                string               `json:"match_client,omitempty"`
	Options                    []Option             `json:"options,omitempty"`
	ExtensibleAttributes       *ExtensibleAttribute `json:"extattrs,omitempty"`
	ExtensibleAttributesAdd    *ExtensibleAttribute `json:"extattrs+,omitempty"`
	ExtensibleAttributesRemove *ExtensibleAttribute `json:"extattrs-,omitempty"`
}

// FixedAddressQueryResult object
type FixedAddressQueryResult struct {
	NextPageID string         `json:"next_page_id,omitempty"`
	Results    []FixedAddress `json:"result,omitempty"`
}

// IPv4Address object
type IPv4Address struct {
	Ref         string   `json:"_ref,omitempty"`
	Hostnames   []string `json:"names,omitempty"`
	IPAddress   string   `json:"ip_address,omitempty"`
	Mac         string   `json:"mac,omitempty"`
	NetworkView string   `json:"network_view,omitempty"`
	CIDR        string   `json:"network,omitempty"`
	Usage       []string `json:"usage,omitempty"`
	Types       []string `json:"types,omitempty"`
	Objects     []string `json:"objects,omitempty"`
	Status      string   `json:"status,omitempty"`
}

// AddressQueryResult object
type AddressQueryResult struct {
	NextPageID string        `json:"next_page_id,omitempty"`
	Results    []IPv4Address `json:"result,omitempty"`
}

// AddressQuery object
type AddressQuery struct {
	NetworkView          string
	FilterEmptyHostnames *bool
	Retries              int
	CIDR                 string
	Count                int
	StartAddress         string
	EndAddress           string
}

func (aq *AddressQuery) fillDefaults() {
	if aq.NetworkView == "" {
		aq.NetworkView = "default"
	}
	if aq.Retries == 0 {
		aq.Retries = 3
	}
	if aq.FilterEmptyHostnames == nil {
		aq.FilterEmptyHostnames = newBool(false)
	}
}

// Range object
type Range struct {
	Ref                        string               `json:"_ref,omitempty"`
	Comment                    string               `json:"comment,omitempty"`
	DisableDHCP                *bool                `json:"disable,omitempty"`
	StartAddress               string               `json:"start_addr,omitempty"`
	EndAddress                 string               `json:"end_addr,omitempty"`
	NetworkView                string               `json:"network_view,omitempty"`
	CIDR                       string               `json:"network,omitempty"`
	Member                     *Member              `json:"member,omitempty"`
	Options                    []Option             `json:"options,omitempty"`
	ExtensibleAttributes       *ExtensibleAttribute `json:"extattrs,omitempty"`
	ExtensibleAttributesAdd    *ExtensibleAttribute `json:"extattrs+,omitempty"`
	ExtensibleAttributesRemove *ExtensibleAttribute `json:"extattrs-,omitempty"`
	IPAddressList              []string             `json:"ip_address_list,omitempty"`
}

// RangeQueryResult object
type RangeQueryResult struct {
	NextPageID string  `json:"next_page_id,omitempty"`
	Results    []Range `json:"result,omitempty"`
}

// IPsWithinRangeQuery object
type IPsWithinRangeQuery struct {
	Ref          string
	CIDR         string
	StartAddress string
	EndAddress   string
}

// ARecord object
type ARecord struct {
	Ref                        string               `json:"_ref,omitempty"`
	Hostname                   string               `json:"name,omitempty"`
	DNSName                    string               `json:"dns_name,omitempty"`
	IPAddress                  string               `json:"ipv4addr,omitempty"`
	Zone                       string               `json:"zone,omitempty"`
	Comment                    string               `json:"comment,omitempty"`
	Disable                    *bool                `json:"disable,omitempty"`
	View                       string               `json:"view,omitempty"`
	ExtensibleAttributes       *ExtensibleAttribute `json:"extattrs,omitempty"`
	ExtensibleAttributesAdd    *ExtensibleAttribute `json:"extattrs+,omitempty"`
	ExtensibleAttributesRemove *ExtensibleAttribute `json:"extattrs-,omitempty"`
}

// ARecordQueryResult object
type ARecordQueryResult struct {
	NextPageID string    `json:"next_page_id,omitempty"`
	Results    []ARecord `json:"result,omitempty"`
}

// CNameRecord object
type CNameRecord struct {
	Ref                        string               `json:"_ref,omitempty"`
	Alias                      string               `json:"name,omitempty"`
	Canonical                  string               `json:"canonical,omitempty"`
	DNSName                    string               `json:"dns_name,omitempty"`
	Zone                       string               `json:"zone,omitempty"`
	Comment                    string               `json:"comment,omitempty"`
	Disable                    *bool                `json:"disable,omitempty"`
	View                       string               `json:"view,omitempty"`
	ExtensibleAttributes       *ExtensibleAttribute `json:"extattrs,omitempty"`
	ExtensibleAttributesAdd    *ExtensibleAttribute `json:"extattrs+,omitempty"`
	ExtensibleAttributesRemove *ExtensibleAttribute `json:"extattrs-,omitempty"`
}

// CNameRecordQueryResult object
type CNameRecordQueryResult struct {
	NextPageID string        `json:"next_page_id,omitempty"`
	Results    []CNameRecord `json:"result,omitempty"`
}

// AliasRecord object
type AliasRecord struct {
	Ref                        string               `json:"_ref,omitempty"`
	Name                       string               `json:"name,omitempty"`
	Target                     string               `json:"target_name,omitempty"`
	TargetType                 string               `json:"target_type,omitempty"`
	DNSName                    string               `json:"dns_name,omitempty"`
	DNSTargetName              string               `json:"dns_target_name,omitempty"`
	Zone                       string               `json:"zone,omitempty"`
	Comment                    string               `json:"comment,omitempty"`
	Disable                    *bool                `json:"disable,omitempty"`
	View                       string               `json:"view,omitempty"`
	ExtensibleAttributes       *ExtensibleAttribute `json:"extattrs,omitempty"`
	ExtensibleAttributesAdd    *ExtensibleAttribute `json:"extattrs+,omitempty"`
	ExtensibleAttributesRemove *ExtensibleAttribute `json:"extattrs-,omitempty"`
}

// AliasRecordQueryResult object
type AliasRecordQueryResult struct {
	NextPageID string        `json:"next_page_id,omitempty"`
	Results    []AliasRecord `json:"result,omitempty"`
}

// PtrRecord object
type PtrRecord struct {
	Ref                        string               `json:"_ref,omitempty"`
	Name                       string               `json:"name,omitempty"`
	PointerDomainName          string               `json:"ptrdname,omitempty"`
	IPv4Address                string               `json:"ipv4addr,omitempty"`
	IPv6Address                string               `json:"ipv6addr,omitempty"`
	DNSName                    string               `json:"dns_name,omitempty"`
	DNSPointerDomainName       string               `json:"dns_ptrdname,omitempty"`
	Zone                       string               `json:"zone,omitempty"`
	Comment                    string               `json:"comment,omitempty"`
	Disable                    *bool                `json:"disable,omitempty"`
	View                       string               `json:"view,omitempty"`
	ExtensibleAttributes       *ExtensibleAttribute `json:"extattrs,omitempty"`
	ExtensibleAttributesAdd    *ExtensibleAttribute `json:"extattrs+,omitempty"`
	ExtensibleAttributesRemove *ExtensibleAttribute `json:"extattrs-,omitempty"`
}

// PtrRecordQueryResult object
type PtrRecordQueryResult struct {
	NextPageID string      `json:"next_page_id,omitempty"`
	Results    []PtrRecord `json:"result,omitempty"`
}

// ResponseError object
type ResponseError struct {
	StatusCode   int
	Request      string
	ResponseBody string
	ErrorMessage string
}
