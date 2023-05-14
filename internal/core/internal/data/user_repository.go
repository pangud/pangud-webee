package data

import (
	"context"

	"go.uber.org/zap"

	"github.com/pangud/pangud/internal/core/internal/biz"
	"github.com/pangud/pangud/internal/core/internal/data/internal"
	"github.com/pangud/pangud/internal/pkg/data"
)

type userRepository struct {
	data  *data.Data
	log   *zap.Logger
	query *internal.Query
}

func NewUserRepository(data *data.Data, log *zap.Logger) biz.UserRepository {
	q := internal.Use(data.DB(context.Background()))
	return &userRepository{data: data, log: log, query: q}
}

func (u userRepository) Save(ctx context.Context, model *biz.User) error {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) FindOne(ctx context.Context, id uint32) (*biz.User, error) {
	u.query.User.WithContext(ctx)
	//TODO implement me
	panic("implement me")
}

func (u userRepository) Remove(ctx context.Context, model *biz.User) error {
	//TODO implement me
	panic("implement me")
}
