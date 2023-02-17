package types

import (
	"context"
	"time"
)

// Model 所有模型应实现的接口
type Model interface {
	// TableName returns the table name of the model
	TableName() string
}

// TimeModel is the base model for all models with time fields.
type TimeModel struct {
	CreateTime time.Time `gorm:"column:create_time;autoCreateTime"`
	UpdateTime time.Time `gorm:"column:update_time;autoUpdateTime"`
}

// Repository 存储库接口 用于实体的获取、保存（新增+更新）、删除
type Repository[T Model] interface {
	// Save 保存模型到存储库
	Save(ctx context.Context, model T) error
	// FindOne 从存储库查找模型
	FindOne(ctx context.Context, id uint32) (T, error)
	// Remove 删除模型
	Remove(ctx context.Context, model T) error
}

// PageQuery 分页查询
type PageQuery[T any] struct {
	// PageNo 分页编号从1开始
	PageNo int32 `json:"page_no"`
	// PageSize 分页页码 最大100 默认10
	PageSize int32 `json:"page_size"`
	// Condition 分页查询条件
	Condition T
}
