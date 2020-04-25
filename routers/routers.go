package routers

import (
	"api.xinfos.com/api/backend"
	v1 "api.xinfos.com/api/v1"

	"github.com/gin-gonic/gin"
)

func Load(g *gin.Engine) *gin.Engine {

	v1Api := g.Group("/v1")
	{
		v1Api.POST("/user/get", v1.GetUserInfoByID)
		v1Api.POST("/user/create", v1.CreateUser)

		v1Api.POST("/product/get", v1.GetBeforeCreateByCatID)
	}

	//backendAPI - Background service management interface
	backendAPI := g.Group("/backend")
	{
		backendAPI.POST("/brand/create", backend.CreateBrand)
		backendAPI.POST("/brand/delete", backend.DeleteBrand)
		backendAPI.POST("/brand/update", backend.UpdateBrand)
		backendAPI.POST("/brand/get", backend.GetBrand)
		backendAPI.POST("/brand/list", backend.ListBrand)
	}

	return g
}
