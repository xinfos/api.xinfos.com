package backend

import (
	"api.xinfos.com/api"
	"api.xinfos.com/internal/model"
	"api.xinfos.com/internal/service"
	"api.xinfos.com/utils/errs"
	"github.com/gin-gonic/gin"
)

type listStaffRequest struct {
	RequestID string `json:"request_id"`
	SellerID  uint64 `json:"seller_id" binding:"required"`
	ShopID    uint64 `json:"shop_id"`
	Name      string `json:"name"`
	Mobile    string `json:"mobile"`
	PageNo    uint   `json:"page" binding:"required"`
	PageSize  uint   `json:"page_size"`
}

//ListStaff - 获取员工列表
func ListStaff(c *gin.Context) {
	var req listStaffRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		api.JSON(c, errs.ErrParamVerify, nil)
		return
	}

	query := []string{"seller_id = (?)", "is_delete = (?)"}
	args := []interface{}{req.SellerID, 2}

	if len(req.Mobile) > 0 {
		query = append(query, "mobile = (?)")
		args = append(args, req.Mobile)
	}

	if len(req.Name) > 0 {
		query = append(query, "name = (?)")
		args = append(args, req.Name)
	}

	if req.PageSize <= 0 || req.PageSize > 20 {
		req.PageSize = 20
	}

	data, err := service.NewShopStaffService().FindAll(model.QueryArrayToString(query), args, "", req.PageNo, req.PageSize)
	if err != nil {
		api.JSON(c, err, nil)
		return
	}
	api.JSON(c, errs.ErrSuccess, data)
	return
}
