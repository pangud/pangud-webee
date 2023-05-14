package biz

import (
	"context"
	"time"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"

	"github.com/jinzhu/copier"
	"github.com/pangud/pangud/pkg/errors"
	"github.com/pangud/pangud/pkg/types"
)

// User user model
type User struct {
	types.IDModel
	Username      string    `json:"username" gorm:"column:username;not null;index:uqx_username,unique"`
	Realname      string    `json:"realname" gorm:"column:realname"`
	Password      string    `json:"password" gorm:"column:password;not null"`
	Nickname      string    `json:"nickname" gorm:"column:nickname"`
	Avatar        string    `json:"avatar" gorm:"column:avatar"`
	Locked        bool      `json:"locked" gorm:"column:locked"`
	LastLoginTime time.Time `json:"last_login_time" gorm:"column:last_login_time"`
	types.TimeModel
}

// TableName db table name of user model
func (u *User) TableName() string {
	return "users"
}

// UserRepository user repository
type UserRepository interface {
	types.Repository[*User]
}

type UserQueryService interface {
	ListUser(ctx context.Context, query *ListUserQuery) (page *types.Page[*User], err error)
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

/*=======================================================
*===========================创建用户======================
*======================================================*/

// CreateUserCommand 创建用户请求数据
type CreateUserCommand struct {
	Username string `json:"username" binding:"required" example:"pduser"`
	Password string `json:"password" binding:"required" example:"your password"`
	Nickname string `json:"nickname" example:"Gaga"`
	Realname string `json:"realname" example:"张三"`
}

// New 创建用户
func (u *UserUsecase) New(ctx context.Context, cmd *CreateUserCommand) error {
	var model User
	copier.Copy(&model, cmd)
	//为用户设置默认密码
	bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(model.Password), bcrypt.DefaultCost)
	if err != nil {
		u.log.Error("generate password error", zap.Error(err))
		return ErrGeneratePasswordError
	}
	(&model).Password = string(bcryptPassword)
	err = u.repo.Save(ctx, &model)
	if err != nil {
		return ErrSaveUserError
	}
	u.log.Info("save user success", zap.Any("user", model))
	return nil
}

/*=======================================================
*===========================更新======================
*======================================================*/

// CreateUserCommand 创建用户请求数据
type UpdateUserCommand struct {
	Nickname string `json:"nickname" example:"Gaga"`
	Realname string `json:"realname" example:"张三"`
	Avatar   string `json:"avatar" example:"https://xxxx.com/xb.jpg"`
	Locked   bool   `json:"locked" example:"false"`
	UserId   uint32 `json:"-"`
}

// UpdateUser update user
func (u *UserUsecase) UpdateUser(ctx context.Context, cmd *UpdateUserCommand) error {
	user, err := u.repo.FindOne(ctx, cmd.UserId)
	if err != nil {
		return errors.DBError
	}
	if cmd.Avatar != "" {
		user.Avatar = cmd.Avatar
	}
	if cmd.Locked != user.Locked {
		user.Locked = cmd.Locked
	}
	if cmd.Nickname != "" {
		user.Nickname = cmd.Nickname
	}
	if cmd.Realname != "" {
		user.Realname = cmd.Realname
	}
	if err = u.repo.Save(ctx, user); err != nil {
		return ErrSaveUserError
	}
	return nil
}

/*=======================================================
*==========================分页查询用户====================
*======================================================*/

// ListUserQuery 查询用户请求数据
type ListUserQuery struct {
	Keywords string `form:"keywords"`
	types.PageQuery[*User]
}

// ListUser 分页查询用户
func (u *UserUsecase) ListUser(ctx context.Context, query *ListUserQuery) (page *types.Page[*User], err error) {
	return u.userQrySvc.ListUser(ctx, query)
}
