package backend

import (
	"log"

	"api.xinfos.com/api"
	"api.xinfos.com/internal/service"
	"api.xinfos.com/utils/errs"
	"github.com/gin-gonic/gin"
)

type listStaffRequest struct {
	RequestID string `json:"request_id"`
	ShopID    uint64 `json:"shop_id" binding:"required"`
	Name      string `json:"name"`
	Mobile    string `json:"mobile"`
	PageNo    uint   `json:"page" binding:"required"`
	PageSize  uint   `json:"page_size"`
}

//ListStaff - 获取员工列表
func ListStaff(c *gin.Context) {
	var req listStaffRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println(err)
		api.JSON(c, errs.ErrParamVerify)
		return
	}

	//assemble query condition
	query := make(map[string]interface{}, 3)

	query["shop_id"] = req.ShopID

	if len(req.Mobile) > 0 {
		query["mobile"] = req.Mobile
	}
	if len(req.Name) > 0 {
		query["name"] = req.Name
	}
	if req.PageSize <= 0 || req.PageSize > 20 {
		req.PageSize = 20
	}

	data, err := service.NewShopStaffService().FindAll(query, "", req.PageNo, req.PageSize)
	if err != nil {
		api.JSON(c, err)
		return
	}
	type hook func(v interface{}) error

	api.JSON(c, errs.ErrSuccess, data)
	return
}
