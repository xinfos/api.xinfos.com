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
		//brand management api interface
		backendAPI.POST("/brand/create", backend.CreateBrand)
		backendAPI.POST("/brand/delete", backend.DeleteBrand)
		backendAPI.POST("/brand/update", backend.UpdateBrand)
		backendAPI.POST("/brand/get", backend.GetBrand)
		backendAPI.POST("/brand/list", backend.ListBrand)

		//category management api interface
		backendAPI.POST("/category/create", backend.CreateCategory)
		backendAPI.POST("/category/delete", backend.DeleteCategory)
		backendAPI.POST("/category/update", backend.UpdateCategory)
		backendAPI.POST("/category/get", backend.GetCategory)
		backendAPI.POST("/category/sub", backend.ListSubCategory)
		backendAPI.POST("/category/attrs/get", backend.GetCategoryAttrs)

		//Attrs management api interface
		backendAPI.POST("/attr/create", backend.CreateAttr)
		backendAPI.POST("/attr/delete", backend.DeleteAttr)
		backendAPI.POST("/attr/update", backend.UpdateAttr)
		backendAPI.POST("/attr/get", backend.GetAttr)
		backendAPI.POST("/attr/list", backend.ListAttr)

		//AttrValues management api interface
		backendAPI.POST("/attrval/create", backend.CreateAttr)
		backendAPI.POST("/attrval/delete", backend.CreateAttr)
		backendAPI.POST("/attrval/update", backend.DeleteAttr)
		backendAPI.POST("/attrval/list", backend.DeleteAttr)

		//product management api interface
		backendAPI.POST("/product/create", backend.CreateProduct)
		backendAPI.POST("/product/delete", backend.DeleteProduct)
		backendAPI.POST("/product/update", backend.UpdateProduct)
		backendAPI.POST("/product/get", backend.GetProduct)
		backendAPI.POST("/product/list", backend.ListProduct)

		//staff management api interface
		backendAPI.POST("/staff/list", backend.ListStaff)
	}

	return g
}
