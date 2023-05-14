package resource

import (
	"net/http"
	"strconv"

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
// @Description 根据用户名、密码、昵称、姓名创建新用户, 其中用户名、密码为必填、其他为可选项
// @Tags User
// @Param CreateUserCommand body biz.CreateUserCommand true "用户信息"
// @Accept json
// @Produce json
// @Success 201 {object} errors.Error "success"
// @Router /users [post]
// @Security ApiKeyAuth
func (u *UserResource) newUser(ctx *gin.Context) {
	var cmd biz.CreateUserCommand
	err := ctx.BindJSON(&cmd)
	if err != nil {
		u.log.Error("fail to get user data", zap.Error(err))
	}
	u.userUsecase.New(ctx, &cmd)
	ctx.JSON(http.StatusCreated, errors.OK())
}

// newUser 更新用户
// @Summary 更新用户
// @Schemes
// @Description 更新用户信息 姓名、昵称、头像
// @Tags User
// @Param id path uint32 true "use id"
// @Param UpdateUserCommand body biz.UpdateUserCommand true "用户信息"
// @Accept json
// @Produce json
// @Success 200 {object} errors.Error "success"
// @Router /users/{id} [patch]
// @Security ApiKeyAuth
func (u *UserResource) updateUser(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(400, biz.ErrUserIdInvalid)
	}
	var cmd biz.UpdateUserCommand
	err = ctx.BindJSON(&cmd)
	if err != nil {
		u.log.Error("fail to get user data", zap.Error(err))
	}
	(&cmd).UserId = uint32(userId)
	u.userUsecase.UpdateUser(ctx, &cmd)
	ctx.JSON(http.StatusCreated, errors.OK())
}

// ListUserResp 分页查询用户响应数据
type ListUserResp struct {
	errors.Error
	Data types.Page[*biz.User]
}

// list 分页查询用户
// @Summary 分页查询用户
// @Schemes
// @Description 分页查询用户
// @Tags User
// @Accept json
// @Produce json
// @Param "query params" query biz.ListUserQuery true "query params"
// @Success 200 {object}  types.ResponseEntity[types.Page[biz.User]]
// @Router /users [get]
// @Security ApiKeyAuth
func (u *UserResource) listUser(ctx *gin.Context) {
	var queryParams biz.ListUserQuery
	// default offset = 0, limit = 10
	// queryParams.SetDefault()
	ctx.BindQuery(&queryParams)
	if queryParams.Limit == 0 {
		(&queryParams).SetDefault()
	}
	u.log.Debug("query params", zap.Any("params", queryParams))

	page, err := u.userUsecase.ListUser(ctx, &queryParams)
	if err != nil {
		u.log.Error("fail to get user data", zap.Error(err))
		return
	}
	ctx.JSON(200, types.NewResponseEntity(errors.OK(), page))
}
