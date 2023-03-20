package resource

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"pangud.io/pangud/internal/account/biz"
)

// @BasePath /api/v1

type UserResource struct {
	log         *zap.Logger
	userUsecase *biz.UserUsecase
}

func NewUserResource(logger *zap.Logger, userUsecase *biz.UserUsecase) *UserResource {
	return &UserResource{
		log:         logger,
		userUsecase: userUsecase,
	}
}

// List 查询用户列表
// @Summary 查询用户列表
// @Schemes
// @Description 分页查询用户列表
// @Tags 用户
// @Accept json
// @Produce json
// @Success 200 {string} string	"ok"
// @Router /users [get]
func (r *UserResource) List(ctx *gin.Context) {
	r.log.Debug("list users")
	//r.userUsecase.List(ctx)
}
