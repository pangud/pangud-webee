package biz

import (
	"go.uber.org/zap"

	"pangud.io/pangud/pkg/types"
)

type User struct {
	UID      uint8  `gorm:"column:uid"`
	Passwd   string `gorm:"column:password"`
	Name     string
	Nickname string
	Avatar   string
}

func (u *User) TableName() string {
	return "users"
}

// UserRepository user repository
type UserRepository interface {
	types.Repository[*User]
}

type UserUsecase struct {
	repo UserRepository
	log  *zap.Logger
}

func NewUserUsecase(repo UserRepository, log *zap.Logger) *UserUsecase {
	return &UserUsecase{
		repo: repo,
		log:  log,
	}
}
