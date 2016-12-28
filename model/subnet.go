package model

import "strings"

type Subneter interface {
	LogicalName() string
}

type Subnet struct {
	//ID                string `yaml:"id,omitempty"`
	AvailabilityZone  string `yaml:"availabilityZone,omitempty"`
	InstanceCIDR      string `yaml:"instanceCIDR,omitempty"`
	RouteTableID      string `yaml:"routeTableId,omitempty"`
	NatGateway        NatGateway `yaml:"natGateway,omitempty"`
}

func (c Subnet) LogicalName() string {
	return "Subnet" + strings.Replace(strings.Title(c.AvailabilityZone), "-", "", -1)
}

type PrivateSubnet struct {
	Subnet `yaml:",inline"`
}

func (c PrivateSubnet) LogicalName() string {
	return "PrivateSubnet" + strings.Replace(strings.Title(c.AvailabilityZone), "-", "", -1)
}

type NatGateway struct {
	ID              string `yaml:"id,omitempty"`
	EIPAllocationID string `yaml:"eipAllocationId,omitempty"`
}
