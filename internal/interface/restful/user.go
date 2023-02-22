package restful

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"pangud.io/pangud/internal/biz"
)

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
func (r *UserResource) List(ctx *gin.Context) {
	r.log.Debug("list users")
	//r.userUsecase.List(ctx)
}
