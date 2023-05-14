package biz

import "context"

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

type Endpoint struct {
	ID     uint32         `gorm:"column:id;primaryKey;autoIncrement;not null"`
	Name   string         `gorm:"column:name;type:varchar(255);not null"`
	Addr   string         `gorm:"column:addr;not null"`
	Token  string         `gorm:"column:token;not null"`
	Type   EndpointType   `gorm:"column:type;not null"`
	Status EndpointStatus `gorm:"column:status;not null"`
}

func (a *Endpoint) TableName() string {
	return "endpoints"
}

// EndpointReadRepository agent读存储库
type EndpointReadRepository interface {
	FindOne(ctx context.Context, id uint32) (*Endpoint, error)
}
