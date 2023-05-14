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
	SetDefault(data.DB(context.Background()))
	return &userQueryService{log: log, query: Q}
}

// ListUser list user by keywords with supporting pagination
func (q *userQueryService) ListUser(ctx context.Context, query *biz.ListUserReq) (*types.Page[*biz.User], error) {
	user := q.query.User
	result, count, err := user.userDo.Where(user.Name.Like(query.Keywords)).FindByPage(query.Offset, query.Limit)
	if err != nil {
		q.log.Error("find user error", zap.Error(err))
		//todo
		return nil, errors.DBError
	}

	page := types.NewPage(result, count, query.Offset, query.Limit)
	return page, nil
}
