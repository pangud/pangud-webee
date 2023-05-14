package biz

import (
	"github.com/pangud/pangud/pkg/types"
)

type EndpointStatus uint8
type EndpointType uint8

const (
	EndpointStatusUnknown EndpointStatus = iota
	EndpointStatusOnline
	EndpointStatusOffline
)
const (
	EndpointTypeUnknown EndpointType = iota
	EndpointTypeServer
	EndpointTypeEdge
	EndpointTypeKube
)

// Endpoint endpoint model
type Endpoint struct {
	types.IDModel
	Name   string         `gorm:"column:name;type:string;size:100;not null"`
	Addr   string         `gorm:"column:addr;not null;type:string;size:255"`
	Token  string         `gorm:"column:token;not null;type:string;size:255"`
	Type   EndpointType   `gorm:"column:type;not null"`
	Status EndpointStatus `gorm:"column:status;not null"`
}

// TableName 表名
func (a *Endpoint) TableName() string {
	return "endpoints"
}

// EndpointReadRepository endpoint repository
type EndpointReadRepository interface {
	types.Repository[*Endpoint]
}
