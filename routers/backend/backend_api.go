package backend

import (
	"api.xinfos.com/api/backend"
	"github.com/gin-gonic/gin"
)

//NewBackendAPI  初始化后台管理接口
func NewBackendAPI(g *gin.Engine) *gin.Engine {

	//backendAPI - 后台管理接口
	backendAPI := g.Group("/backend")
	{
		//店铺管理接口
		backendAPI.POST("/shop/create", backend.CreateShop)
		backendAPI.POST("/shop/delete", backend.DeleteShop)
		backendAPI.POST("/shop/update", backend.UpdateShop)
		backendAPI.POST("/shop/dashboard", backend.DashboardShop)
		backendAPI.POST("/shop/get", backend.GetShop)
		backendAPI.POST("/shop/info", backend.InfoShop)
		backendAPI.POST("/shop/list", backend.ListShop)

		//店铺员工管理接口
		backendAPI.POST("/staff/list", backend.ListStaff)

		//品牌管理接口
		backendAPI.POST("/brand/create", backend.CreateBrand)
		backendAPI.POST("/brand/delete", backend.DeleteBrand)
		backendAPI.POST("/brand/update", backend.UpdateBrand)
		backendAPI.POST("/brand/get", backend.GetBrand)
		backendAPI.POST("/brand/list", backend.ListBrand)

		//商品分类管理接口
		backendAPI.POST("/category/create", backend.CreateCategory)
		backendAPI.POST("/category/delete", backend.DeleteCategory)
		backendAPI.POST("/category/update", backend.UpdateCategory)
		backendAPI.POST("/category/get", backend.GetCategory)
		backendAPI.POST("/category/sub", backend.ListSubCategory)
		backendAPI.POST("/category/attrs/get", backend.GetCategoryAttrs)
		backendAPI.POST("/category/list", backend.ListCategory)

		//商品分组管理接口
		backendAPI.POST("/attrgroup/create", backend.CreateAttrGroup)

		//商品属性管理接口
		backendAPI.POST("/attr/create", backend.CreateAttr)
		backendAPI.POST("/attr/delete", backend.DeleteAttr)
		backendAPI.POST("/attr/update", backend.UpdateAttr)
		backendAPI.POST("/attr/get", backend.GetAttr)
		backendAPI.POST("/attr/list", backend.ListAttr)
		backendAPI.POST("/attr/query", backend.QueryAttr)

		//商品属性值管理接口
		backendAPI.POST("/attrval/create", backend.CreateAttr)
		backendAPI.POST("/attrval/delete", backend.CreateAttr)
		backendAPI.POST("/attrval/update", backend.DeleteAttr)
		backendAPI.POST("/attrval/list", backend.DeleteAttr)

		//商品属性模板管理接口
		backendAPI.POST("/attr/template/create", backend.CreateAttrTemplate)

		//商品管理接口
		backendAPI.POST("/product/create", backend.CreateProduct)
		backendAPI.POST("/product/delete", backend.DeleteProduct)
		backendAPI.POST("/product/update", backend.UpdateProduct)
		backendAPI.POST("/product/get", backend.GetProduct)
		backendAPI.POST("/product/list", backend.ListProduct)
	}

	return g
}
