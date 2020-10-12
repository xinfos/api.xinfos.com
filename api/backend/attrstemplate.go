package backend

import (
	"api.xinfos.com/api"
	"api.xinfos.com/internal/model"
	"api.xinfos.com/internal/service"
	"api.xinfos.com/utils/errs"
	"github.com/gin-gonic/gin"
)

type createAttrTemplateRequest struct {
	RequestID    string   `json:"request_id"`
	Name         string   `json:"name" binding:"required"`
	CatID        uint64   `json:"cat_id" binding:"required"`
	State        uint8    `json:"state"`
	AttrGroupIDs []uint64 `json:"attrgroups"`
}

//CreateAttrTemplate Create attribute template
func CreateAttrTemplate(c *gin.Context) {

	var req createAttrTemplateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		api.JSON(c, errs.ErrParamVerify, nil)
		return
	}

	sGroupID, errmsg := service.NewAttrTemplateService().Create(&model.SAttrTempalte{
		Name:  req.Name,
		CatID: req.CatID,
		State: req.State,
	}, AttrGroupIDs)

	if errmsg != nil {
		api.JSON(c, errmsg)
		return
	}
	api.JSON(c, errs.ErrSuccess, map[string]uint64{"group_id": sGroupID})
	return
}
