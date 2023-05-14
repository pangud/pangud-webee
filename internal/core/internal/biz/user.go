package biz

import (
	"go.uber.org/zap"

	"github.com/pangud/pangud/pkg/types"
)

// User user model
type User struct {
	UID      uint8  `gorm:"column:uid"`
	Passwd   string `gorm:"column:password"`
	Name     string
	Nickname string
	Avatar   string
}

// TableName db table name of user model
func (u *User) TableName() string {
	return "users"
}

// UserRepository user repository
type UserRepository interface {
	types.Repository[*User]
}

// UserUsecase user usecase
type UserUsecase struct {
	repo UserRepository
	log  *zap.Logger
}

// NewUserUsecase new user usecase
func NewUserUsecase(repo UserRepository, log *zap.Logger) *UserUsecase {
	return &UserUsecase{
		repo: repo,
		log:  log,
	}
}
