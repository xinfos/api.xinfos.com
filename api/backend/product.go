package backend

import (
	"api.xinfos.com/api"
	"api.xinfos.com/internal/service"
	"api.xinfos.com/utils/errs"
	"github.com/gin-gonic/gin"
)

// type Attr struct {
// 	AttrID  uint64
// 	AttrVal string
// }

// type Sepc struct {
// }

// type createProductRequest struct {
// 	Title string
// 	props []*Attr
// 	binds []*Spec
// 	sale_props
// 	customer_props
// 	一口价
// }
type getProductAttrsRequest struct {
	Request string `json:"request_id"`
	CatID   uint64 `json:"cat_id"`
}

//CreateProduct - create new category
func CreateProduct(c *gin.Context) {

}

//DeleteProduct - Delete category By BrandID
func DeleteProduct(c *gin.Context) {

}

//UpdateProduct - Update category By BrandID
func UpdateProduct(c *gin.Context) {

}

//GetProduct - Get category By BrandID
func GetProduct(c *gin.Context) {

}

//ListProduct - Get all category List
func ListProduct(c *gin.Context) {

	return
}

//GetProductAttrs - Get Product attrs
func GetProductAttrs(c *gin.Context) {
	var req getProductAttrsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		api.JSON(c, errs.ErrParamVerify, nil)
		return
	}
	data, _ := service.NewProductService().BeforeCreateByCatID(req.CatID)
	api.JSON(c, errs.ErrSuccess, data)
	return
}
