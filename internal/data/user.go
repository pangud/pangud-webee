package data

import (
	"context"

	"go.uber.org/zap"

	"pangud.io/pangud/internal/apiserver/biz"
)

type userRepository struct {
	data *Data
	log  *zap.Logger
}

func NewUserRepository(data *Data, log *zap.Logger) biz.UserRepository {
	return &userRepository{data: data, log: log}
}

func (u userRepository) Save(ctx context.Context, model *biz.User) error {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) FindOne(ctx context.Context, id uint32) (*biz.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) Remove(ctx context.Context, model *biz.User) error {
	//TODO implement me
	panic("implement me")
}
