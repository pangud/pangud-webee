package data

import (
	"context"

	"go.uber.org/zap"

	"github.com/pangud/pangud/internal/core/biz"
	"github.com/pangud/pangud/internal/pkg/data"
	"github.com/pangud/pangud/pkg/errors"
	"github.com/pangud/pangud/pkg/types"
)

// userQueryService user query service implement
type userQueryService struct {
	log   *zap.Logger
	query *Query
}

// NewUserQueryService create user query service
func NewUserQueryService(log *zap.Logger, data *data.Data) biz.UserQueryService {
	return &userQueryService{log: log, query: Q}
}

// ListUser list user by keywords with supporting pagination
func (q *userQueryService) ListUser(ctx context.Context, query *biz.ListUserQuery) (*types.Page[*biz.User], error) {
	user := q.query.User

	queryStr := "%" + query.Keywords + "%"
	result, count, err := user.userDo.
		Where(user.Username.Like(queryStr)).
		Or(user.Realname.Like(queryStr)).
		Or(user.Nickname.Like(queryStr)).
		FindByPage(query.Offset, query.Limit)
	if err != nil {
		q.log.Error("find user error", zap.Error(err))
		//todo
		return nil, errors.DBError
	}
	q.log.Info("find user", zap.Any("user", result))
	page := types.NewPage(result, count, query.Offset, query.Limit)
	return page, nil
}
