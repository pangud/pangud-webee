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

type IDModel struct {
	ID uint32 `json:"id" gorm:"autoIncrement;primaryKey;column:id;not null"`
}

// TimeModel is the base model for all models with time fields.
type TimeModel struct {
	CreateTime time.Time `json:"create_time" gorm:"column:create_time;autoCreateTime"`
	UpdateTime time.Time `json:"update_time" gorm:"column:update_time;autoUpdateTime"`
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
	Limit  int `form:"limit" example:"10" default:"10"`
	Offset int `form:"offset" example:"0" default:"0"`
	// Condition 分页查询条件
	Condition T
}

func (p *PageQuery[T]) SetDefault() {
	p.Limit = 10
}

// Page 分页数据
type Page[T any] struct {
	List []T `json:"list"`
	// Total 总数
	Total  int64 `json:"total"`
	Offset int   `json:"offset"`
	Limit  int   `json:"limit"`
}

// NewPage 新建分页数据
func NewPage[T any](list []T, total int64, offset, limit int) *Page[T] {
	return &Page[T]{
		List:   list,
		Total:  total,
		Offset: offset,
		Limit:  limit,
	}
}
