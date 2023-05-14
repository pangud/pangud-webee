package resource

import (
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
	"github.com/pangud/pangud/internal/pkg/data"
	"github.com/pangud/pangud/pkg/errors"
	"github.com/pangud/pangud/pkg/types"

	"github.com/pangud/pangud/internal/core/biz"
)

type UserResource struct {
	log *zap.Logger

	data        *data.Data
	userUsecase *biz.UserUsecase
}

func NewUserResource(log *zap.Logger, data *data.Data, userUsecase *biz.UserUsecase) *UserResource {
	return &UserResource{log: log, data: data, userUsecase: userUsecase}
}

// newUser 新增用户
// @Summary 新增用户
// @Schemes
// @Description 新增用户
// @Tags User
// @Param createUserReqData body biz.User true "user info"
// @Accept json
// @Produce json
// @Success 200 {string} string	"ok"
// @Router /users [post]
// @Security ApiKeyAuth
func (u *UserResource) newUser(ctx *gin.Context) {
	var user biz.User
	err := ctx.BindJSON(&user)
	if err != nil {
		u.log.Error("fail to get user data", zap.Error(err))
	}
	u.userUsecase.New(ctx, &user)
	ctx.JSON(200, []string{"OK"})
}

type ListUserResp struct {
	errors.Error
	Data types.Page[*biz.User]
}

// list 分页查询用户
// @Summary 分页查询用户
// @Schemes
// @Description 分页查询用户
// @Tags User
// @Param "query params" query biz.ListUserReq true "query params"
// @Accept json
// @Produce json
// @Success 200 {object}  biz.ListUserReq
// @Router /users [get]
// @Security ApiKeyAuth
func (u *UserResource) listUser(ctx *gin.Context) {
	var queryParams biz.ListUserReq
	ctx.BindQuery(&queryParams)
	u.log.Debug("query params", zap.Any("params", queryParams))

	page, err := u.userUsecase.ListUser(ctx, &queryParams)
	if err != nil {
		u.log.Error("fail to get user data", zap.Error(err))
		return
	}
	ctx.JSON(200, ListUserResp{Data: *page, Error: errors.OK()})
}
