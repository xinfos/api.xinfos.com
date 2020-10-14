package backend

import (
	"fmt"

	"api.xinfos.com/api"
	"api.xinfos.com/internal/model"
	"api.xinfos.com/internal/service"
	"api.xinfos.com/utils/errs"
	"github.com/gin-gonic/gin"
)

type generalAttrs struct {
	GroupID uint64   `json:"gourp_id"`
	AttrIDs []uint64 `json:"attr_ids"`
}
type createAttrTemplateRequest struct {
	RequestID    string         `json:"request_id"`
	Name         string         `json:"name" binding:"required"`
	CatID        uint64         `json:"cat_id" binding:"required"`
	State        uint8          `json:"state"`
	SpecAttrs    []uint64       `json:"spec_attrs"`
	GeneralAttrs []generalAttrs `json:"general_attrs"`
}

//CreateAttrTemplate - Create attribute template
func CreateAttrTemplate(c *gin.Context) {

	var req createAttrTemplateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		api.JSON(c, errs.ErrParamVerify, nil)
		return
	}

	generalGroupIDs := []uint64{}
	generalAttrIDs := map[uint64][]uint64{}
	if len(req.GeneralAttrs) > 0 {
		for _, v := range req.GeneralAttrs {
			generalGroupIDs = append(generalGroupIDs, v.GroupID)
			generalAttrIDs[v.GroupID] = v.AttrIDs
		}
	}
	fmt.Println(generalGroupIDs)
	fmt.Println(generalAttrIDs)

	api.JSON(c, errs.ErrSuccess, map[string]uint64{"group_id": 1})
	return

	sGroupID, errmsg := service.NewAttrTemplateService().Create(&model.SAttrTempalte{
		Name:  req.Name,
		CatID: req.CatID,
		State: req.State,
	}, req.SpecAttrs, req.SpecAttrs)

	if errmsg != nil {
		api.JSON(c, errmsg)
		return
	}
	api.JSON(c, errs.ErrSuccess, map[string]uint64{"group_id": sGroupID})
	return
}
