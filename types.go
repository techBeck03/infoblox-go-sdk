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
	Value             interface{}        `json:"value,omitempty"`
	InheritanceSource *InheritanceSource `json:"inheritance_source,omitempty"`
}

// InheritanceSource defines inheritance of an EA
type InheritanceSource struct {
	Ref string `json:"_ref,omitempty"`
}

// ExtensibleAttributeJSONMap ea object in terraform friendly JSON
type ExtensibleAttributeJSONMap map[string]ExtensibleAttributeJSONMapValue

// ExtensibleAttributeJSONMapValue value of ea in terraform friendly JSON
type ExtensibleAttributeJSONMapValue struct {
	Value interface{} `json:"value,omitempty"`
	Type  string      `json:"type,omitempty"`
}

// Network object
type Network struct {
	Ref                   string               `json:"_ref,omitempty"`
	NetworkView           string               `json:"network_view,omitempty"`
	CIDR                  string               `json:"network,omitempty"`
	ExtensibleAttributes  *ExtensibleAttribute `json:"extattrs,omitempty"`
	Comment               string               `json:"comment,omitempty"`
	RestartIfNeeded       *bool                `json:"restart_if_needed,omitempty"`
	AutoCreateReverseZone *bool                `json:"auto_create_reversezone,omitempty"`
	DisableDHCP           *bool                `json:"disable,omitempty"`
	Members               []Member             `json:"members,omitempty"`
	Options               []Option             `json:"options,omitempty"`
}

// Member defines grid members
type Member struct {
	StructType  string `json:"_struct,omitempty"`
	Name        string `json:"name,omitempty"`
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
	Ref                  string               `json:"_ref,omitempty"`
	Hostname             string               `json:"name,omitempty"`
	Comment              string               `json:"comment,omitempty"`
	EnableDNS            *bool                `json:"configure_for_dns,omitempty"`
	ExtensibleAttributes *ExtensibleAttribute `json:"extattrs,omitempty"`
	IPv4Addrs            []IPv4Addr           `json:"ipv4addrs,omitempty"`
	NetworkView          string               `json:"network_view,omitempty"`
	RestartIfNeeded      *bool                `json:"restart_if_needed,omitempty"`
	View                 string               `json:"view,omitempty"`
	Zone                 string               `json:"zone,omitempty"`
}

// IPv4Addr object
type IPv4Addr struct {
	Ref              string                 `json:"_ref,omitempty"`
	Host             string                 `json:"host,omitempty"`
	IPAddress        string                 `json:"ipv4addr,omitempty"`
	Mac              string                 `json:"mac,omitempty"`
	CIDR             string                 `json:"network,omitempty"`
	ConfigureForDHCP *bool                  `json:"configure_for_dhcp,omitempty"`
	NextServer       string                 `json:"nextserver,omitempty"`
	ObjectFunction   string                 `json:"_object_function,omitempty"`
	Parameters       map[string]interface{} `json:"_parameters,omitempty"`
	ResultField      string                 `json:"_result_field,omitempty"`
	Object           string                 `json:"_object,omitempty"`
	ObjectParameters map[string]interface{} `json:"_object_parameters,omitempty"`
}

// FixedAddress object
type FixedAddress struct {
	Ref                  string               `json:"_ref,omitempty"`
	NetworkView          string               `json:"network_view,omitempty"`
	CIDR                 string               `json:"network,omitempty"`
	Comment              string               `json:"comment,omitempty"`
	IPAddress            string               `json:"ipv4addr,omitempty"`
	Mac                  string               `json:"mac,omitempty"`
	Hostname             string               `json:"name,omitempty"`
	MatchClient          string               `json:"match_client,omitempty"`
	ExtensibleAttributes *ExtensibleAttribute `json:"extattrs,omitempty"`
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
	NetworkView  string
	CIDR         string
	Count        int
	StartAddress string
	EndAddress   string
}

func (aq *AddressQuery) fillDefaults() {
	if aq.NetworkView == "" {
		aq.NetworkView = "default"
	}
}

// Range object
type Range struct {
	Ref                  string               `json:"_ref,omitempty"`
	Comment              string               `json:"comment,omitempty"`
	DisableDHCP          *bool                `json:"disable,omitempty"`
	StartAddress         string               `json:"start_addr,omitempty"`
	EndAddress           string               `json:"end_addr,omitempty"`
	NetworkView          string               `json:"network_view,omitempty"`
	CIDR                 string               `json:"network,omitempty"`
	ExtensibleAttributes *ExtensibleAttribute `json:"extattrs,omitempty"`
}
