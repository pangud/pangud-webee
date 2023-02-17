package biz

import "pangud.io/pangud/pkg/types"

type Group struct {
	GID  uint8 `gorm:"column:gid"`
	Name string
}

func (g *Group) TableName() string {
	return "t_group"
}

// GroupRepository user repository
type GroupRepository interface {
	types.Repository[*User]
}
