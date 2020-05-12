package backend

import (
	"api.xinfos.com/api"
	"api.xinfos.com/utils/errs"
	"github.com/gin-gonic/gin"
)

type createAttrRequest struct {
	SGroupID    uint64 `json:"s_group_id"`
	CatID       uint64 `json:"cat_id"`
	Name        string `json:"name"`
	FillType    uint   `json:"fill_type"`
	IsRequired  uint   `json:"is_required"`
	IsNumeric   uint   `json:"is_numeric"`
	Unit        string `json:"unit"`
	IsGeneric   uint   `json:"is_generic"`
	IsSearching uint   `json:"is_searching"`
	Segments    string `json:"segments"`
}

//CreateAttr - create new category
func CreateAttr(c *gin.Context) {
	var req createAttrRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		api.JSON(c, errs.ErrParamVerify, nil)
		return
	}
	// id, err := service.NewAttrService().Create(&model.SAttr{
	// 	SGroupID:    req.SGroupID,
	// 	CatID:       req.SGroupID,
	// 	Name:        req.Name,
	// 	FillType:    req.FillType,
	// 	IsRequired:  req.IsRequired,
	// 	IsNumeric:   req.IsNumeric,
	// 	Unit:        req.Unit,
	// 	IsGeneric:   req.IsGeneric,
	// 	IsSearching: req.IsSearching,
	// 	Segments:    req.Segments,
	// })
	// if err != nil {
	// 	api.JSON(c, err)
	// 	return
	// }
	api.JSON(c, errs.ErrSuccess, map[string]uint64{"cat_id": 0})
	return
}

//DeleteAttr - Delete category By BrandID
func DeleteAttr(c *gin.Context) {

}

//UpdateAttr - Update category attr By BrandID
func UpdateAttr(c *gin.Context) {

}

//GetAttr - Get category By BrandID
func GetAttr(c *gin.Context) {

}

//ListAttr - Get all category List
func ListAttr(c *gin.Context) {

}
