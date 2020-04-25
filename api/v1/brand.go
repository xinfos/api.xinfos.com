package v1

import (
	"api.xinfos.com/api"
	"api.xinfos.com/utils/errs"
	"github.com/gin-gonic/gin"
)

type CreateBrandRequest struct {
	RequestID string `json:"request_id"`
	BrandName string `json:"brand_name" binding:"required"`
	BrandLogo string `json:"brand_logo" binding:"required"`
	BrandDesc string `json:"brand_desc"`
	CatID     uint64 `json:"cat_id" binding:"required"`
	SortOrder uint32 `json:"sort_order"`
	IsShow    uint   `json:"is_show"`
}

//CreateBrand -
func CreateBrand(c *gin.Context) {
	var req CreateBrandRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		api.JSON(c, errs.ErrParamVerify, nil)
		return
	}

}
