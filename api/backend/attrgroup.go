package backend

import (
	"api.xinfos.com/api"
	"api.xinfos.com/internal/model"
	"api.xinfos.com/internal/service"
	"api.xinfos.com/utils/errs"
	"github.com/gin-gonic/gin"
)

type createAttrGroupRequest struct {
	RequestID string `json:"request_id"`
	Name      string `json:"name" binding:"required"`
	CatID     uint64 `json:"cat_id" binding:"required"`
}

//CreateAttrGroup Create a single attribute group
func CreateAttrGroup(c *gin.Context) {
	var req createAttrGroupRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		api.JSON(c, errs.ErrParamVerify, nil)
		return
	}
	sGroupID, errmsg := service.NewAttrGroupService().Create(&model.SAttrGroup{
		Name:  req.Name,
		CatID: req.CatID,
	})
	if errmsg != nil {
		api.JSON(c, errmsg)
		return
	}
	api.JSON(c, errs.ErrSuccess, map[string]uint64{"group_id": sGroupID})
	return
}
