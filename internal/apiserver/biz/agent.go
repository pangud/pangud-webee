package biz

import "context"

type AgentStatus uint8
type AgentType uint8

const (
	AgentStatusUnknown AgentStatus = iota
	AgentStatusOnline
	AgentStatusOffline
)
const (
	AgentTypeUnknown AgentType = iota
	AgentTypeServer
	AgentTypeEdge
	AgentTypeKube
)

type Agent struct {
	ID     uint32      `gorm:"column:id;primaryKey;autoIncrement;not null"`
	Name   string      `gorm:"column:name;type:varchar(255);not null"`
	Addr   string      `gorm:"column:addr;not null"`
	Token  string      `gorm:"column:token;not null"`
	Type   AgentType   `gorm:"column:type;not null"`
	Status AgentStatus `gorm:"column:status;not null"`
}

func (a *Agent) TableName() string {
	return "m_endpoints"
}

// AgentReadRepository agent读存储库
type AgentReadRepository interface {
	FindOne(ctx context.Context, id uint32) (*Agent, error)
}
