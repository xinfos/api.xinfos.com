package v1

import (
	"api.xinfos.com/api"
	"api.xinfos.com/internal/service"
	"api.xinfos.com/utils/errs"
	"github.com/gin-gonic/gin"
)

type GetBeforeCreateRequest struct {
	Request string `json:"request_id"`
	CatID   uint64 `json:"cat_id"`
}

//GetBeforeCreateByCatID - Get user info by user_id
func GetBeforeCreateByCatID(c *gin.Context) {
	var req GetBeforeCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		api.JSON(c, errs.ErrParamVerify, nil)
		return
	}
	data, _ := service.NewProductService().BeforeCreateByCatID(req.CatID)
	api.JSON(c, errs.ErrSuccess, data)
	return
}
