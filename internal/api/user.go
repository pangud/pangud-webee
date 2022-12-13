package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"pangud.io/pangud/internal/sys/user"
)

type UserResource struct {
	log         *zap.Logger
	userUsecase *user.UserUsecase
}

func NewUserResource(logger *zap.Logger, userUsecase *user.UserUsecase) *UserResource {
	return &UserResource{
		log:         logger,
		userUsecase: userUsecase,
	}
}
func (r *UserResource) List(ctx *gin.Context) {
	r.log.Debug("list users")
	//r.userUsecase.List(ctx)
}
