package resource

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewAccountAPI, NewUserResource)

type AccountAPI struct {
	engine       *gin.Engine
	userResource *UserResource
}

// NewAccountAPI new an account api and return.
func NewAccountAPI(engine *gin.Engine, userResource *UserResource) *AccountAPI {
	return &AccountAPI{
		engine:       engine,
		userResource: userResource,
	}
}

func (api *AccountAPI) Register() {
	apiv1 := api.engine.Group("/api/v1")
	apiv1.Group("users").GET("", api.userResource.List)
}
