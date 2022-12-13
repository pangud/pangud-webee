package user

import (
	"github.com/google/wire"
	"go.uber.org/zap"
	"pangud.io/pangud/internal/types"
)

var ProviderSet = wire.NewSet(NewUserUsecase)

type User struct {
	UID      uint8  `gorm:"column:uid"`
	Passwd   string `gorm:"column:password"`
	Name     string
	Nickname string
	Avatar   string
}

func (u *User) TableName() string {
	return "t_user"
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
