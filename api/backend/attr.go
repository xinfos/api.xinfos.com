package backend

import (
	"api.xinfos.com/api"
	"api.xinfos.com/internal/model"
	"api.xinfos.com/internal/service"
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

type getAttrRequest struct {
	RequestID string `json:"request_id"`
	AttrID    uint64 `json:"attr_id" binding:"required"`
}

type listAttrRequest struct {
	RequestID string `json:"request_id"`
	Name      string `json:"attr_name"`
	PageNo    uint   `json:"page_no" binding:"required"`
	PageSize  uint   `json:"page_size"`
}

type queryAttrRequest struct {
	Search string `json:"search" binding:"required"`
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

/**
* @api {post} /backend/attr/get 获取单个属性详情
* @apiName GetAttr
* @apiGroup 后台管理接口/属性管理
*
* @apiParam {String} request_id  请求ID
* @apiParam {String} [Name]      属性名称
* @apiParam {Number} [PageNo]    页数
* @apiParam {Number} [PageSize]  每页显示个数
*
* @apiSuccess {String} firstname Firstname of the User.
* @apiSuccess {String} lastname  Lastname of the User.
*
* @apiSuccessExample Success-Response:
*     HTTP/1.1 200 OK
*     {
*       "firstname": "John",
*       "lastname": "Doe"
*     }
*
* @apiUse UserNotFoundError
 */
func GetAttr(c *gin.Context) {

	var req getAttrRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		api.JSON(c, errs.ErrParamVerify, nil)
		return
	}
	data, err := service.NewAttrService().FindByID(req.AttrID)
	if err != nil {
		api.JSON(c, err, nil)
		return
	}
	api.JSON(c, errs.ErrSuccess, data)
	return
}

/**
* @api {post} /backend/attr/list 获取属性列表
* @apiName ListAttr
* @apiGroup 后台管理接口/属性管理
*
* @apiParam {String} request_id  请求ID
* @apiParam {String} [Name]      属性名称
* @apiParam {Number} [PageNo]    页数
* @apiParam {Number} [PageSize]  每页显示个数
*
* @apiSuccess {String} firstname Firstname of the User.
* @apiSuccess {String} lastname  Lastname of the User.
*
* @apiSuccessExample Success-Response:
*     HTTP/1.1 200 OK
*     {
*       "firstname": "John",
*       "lastname": "Doe"
*     }
*
* @apiUse UserNotFoundError
 */
func ListAttr(c *gin.Context) {
	var req listAttrRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		api.JSON(c, errs.ErrParamVerify, nil)
		return
	}

	query := []string{"is_delete = (?)"}
	args := []interface{}{2}
	//属性名称
	if len(req.Name) > 0 {
		query = append(query, "name = (?)")
		args = append(args, req.Name)
	}
	if req.PageSize <= 0 || req.PageSize > 20 {
		req.PageSize = 20
	}

	data, err := service.NewAttrService().FindAll(model.QueryArrayToString(query), args, "", req.PageNo, req.PageSize)
	if err != nil {
		api.JSON(c, err, nil)
		return
	}
	api.JSON(c, errs.ErrSuccess, data)
	return
}

//QueryAttr - Query attr list
func QueryAttr(c *gin.Context) {
	var req queryAttrRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		api.JSON(c, errs.ErrParamVerify, nil)
		return
	}

	searchRes, errmsg := service.NewAttrService().Query(req.Search)
	if errmsg != nil {
		api.JSON(c, errmsg)
		return
	}
	api.JSON(c, errs.ErrSuccess, searchRes)
	return
}
