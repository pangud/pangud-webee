package resource

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewCoreAPI, NewUserResource)

type CoreAPI struct {
	engine       *gin.Engine
	userResource *UserResource
}

// NewSSLCertPI new an account api and return.
func NewCoreAPI(engine *gin.Engine, userResource *UserResource) *CoreAPI {
	return &CoreAPI{
		engine:       engine,
		userResource: userResource,
	}
}

func (api *CoreAPI) Register() {
	apiv1 := api.engine.Group("/api/v1")

	apiv1.POST("users", api.userResource.newUser)
	apiv1.GET("users", api.userResource.listUser)
}
