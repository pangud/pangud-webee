package biz

import (
	"context"

	"go.uber.org/zap"

	"github.com/pangud/pangud/pkg/types"
)

// User user model
type User struct {
	ID       uint32 `gorm:"column:uid"`
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

type ListUserReq struct {
	Keywords string `form:"keywords"`
	Offset   int    `form:"offset"`
	Limit    int    `form:"limit"`
}

type UserQueryService interface {
	ListUser(ctx context.Context, query *ListUserReq) (page *types.Page[*User], err error)
}

// UserUsecase user usecase
type UserUsecase struct {
	userQrySvc UserQueryService
	repo       UserRepository
	log        *zap.Logger
}

// NewUserUsecase new user usecase
func NewUserUsecase(userQrySvc UserQueryService, repo UserRepository, log *zap.Logger) *UserUsecase {
	return &UserUsecase{
		userQrySvc: userQrySvc,
		repo:       repo,
		log:        log,
	}
}

func (u *UserUsecase) New(ctx context.Context, model *User) error {
	err := u.repo.Save(ctx, model)
	if err != nil {
		return ErrSaveUserError

	}
	u.log.Info("save user success", zap.Any("user", model))
	return nil
}
func (u *UserUsecase) ListUser(ctx context.Context, query *ListUserReq) (page *types.Page[*User], err error) {
	return u.userQrySvc.ListUser(ctx, query)
}
