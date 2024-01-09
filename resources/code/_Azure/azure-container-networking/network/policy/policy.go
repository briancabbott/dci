package policy

import (
	"encoding/json"
)

const (
	NetworkPolicy     CNIPolicyType = "NetworkPolicy"
	EndpointPolicy    CNIPolicyType = "EndpointPolicy"
	OutBoundNatPolicy CNIPolicyType = "OutBoundNAT"
	RoutePolicy       CNIPolicyType = "ROUTE"
	PortMappingPolicy CNIPolicyType = "NAT"
	ACLPolicy         CNIPolicyType = "ACL"
	L4WFPProxyPolicy  CNIPolicyType = "L4WFPPROXY"
	LoopbackDSRPolicy CNIPolicyType = "LoopbackDSR"
)

type CNIPolicyType string

type Policy struct {
	Type CNIPolicyType
	Data json.RawMessage
}

// NATInfo contains information about NAT rules
type NATInfo struct {
	Destinations []string
	VirtualIP    string
}
