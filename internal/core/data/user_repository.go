package data

import (
	"context"

	"go.uber.org/zap"

	"github.com/pangud/pangud/internal/core/biz"
	"github.com/pangud/pangud/internal/pkg/data"
)

// userRepository implement user repository
type userRepository struct {
	data *data.Data
	log  *zap.Logger
	user user
}

// NewUserRepository create user repository
func NewUserRepository(data *data.Data, log *zap.Logger) biz.UserRepository {
	return &userRepository{data: data, log: log, user: newUser(data.DB(context.Background()))}
}

func (u userRepository) Save(ctx context.Context, model *biz.User) error {
	user := u.user.clone(u.data.DB(ctx))
	err := user.userDo.Save(model)
	if err != nil {
		u.log.Error("save user error", zap.Error(err))
		return err
	}
	u.log.Info("save user", zap.Any("user", model))
	return nil
}

func (u userRepository) FindOne(ctx context.Context, id uint32) (userModel *biz.User, err error) {
	user := u.user.clone(u.data.DB(ctx))
	userModel, err = user.userDo.Where(user.ID.Eq(id)).First()
	if err != nil {
		u.log.Error("find user error", zap.Error(err))
	}
	u.log.Info("find user", zap.Any("user", userModel))
	return

}

func (u userRepository) Remove(ctx context.Context, model *biz.User) error {
	user := u.user.clone(u.data.DB(ctx))
	reslut, err := user.userDo.Delete(model)
	if err != nil {
		return err
	}
	u.log.Info("delete user", zap.Any("user", model), zap.Any("reslut", reslut))
	return nil
}
