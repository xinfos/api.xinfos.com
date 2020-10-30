package routers

import (
	backend "api.xinfos.com/routers/backend"
	v1 "api.xinfos.com/routers/v1"
	"github.com/gin-gonic/gin"
)

//Load 加载相关路由
func Load(g *gin.Engine) *gin.Engine {

	v1.NewV1API(g)

	backend.NewBackendAPI(g)

	return g
}
