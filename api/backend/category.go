package backend

import (
	"fmt"
	"strings"

	"api.xinfos.com/api"
	"api.xinfos.com/internal/model"
	"api.xinfos.com/internal/service"
	"api.xinfos.com/pkg/logger"
	"api.xinfos.com/utils/errs"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type createCategoryRequest struct {
	RequestID string `json:"request_id"`
	PID       uint64 `json:"pid"`
	Name      string `json:"name" binding:"required"`
	Alias     string `json:"alias"`
	Desc      string `json:"desc"`
	ShowInNav uint   `json:"show_in_nav"`
	IsShow    uint   `json:"is_show"`
	IsParent  uint   `json:"is_parent"`
	State     uint   `json:"state"`
}
type categoryRequest struct {
	RequestID string `json:"request_id"`
	CatID     uint64 `json:"cat_id" binding:"required"`
}

type updateCategoryRequest struct {
	RequestID string `json:"request_id"`
	CatID     uint64 `json:"cat_id" binding:"required"`
	PID       uint64 `json:"pid"`
	Name      string `json:"name" binding:"required"`
	Alias     string `json:"alias"`
	Desc      string `json:"desc"`
	ShowInNav uint   `json:"show_in_nav"`
	IsShow    uint   `json:"is_show"`
	IsParent  uint   `json:"is_parent"`
}

type listCategoryRequest struct {
	RequestID string `json:"request_id"`
	CatID     uint64 `json:"cat_id"`
	Name      string `json:"name"`
	PageNo    uint   `json:"page_no" binding:"required"`
	PageSize  uint   `json:"page_size"`
}

//CreateCategory - create new category
func CreateCategory(c *gin.Context) {
	var req createCategoryRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		api.JSON(c, errs.ErrParamVerify, nil)
		logger.Error(err.Error())
		return
	}
	logger.Info("request params", zap.Any("params", req))
	id, err := service.NewCategoryService().Create(&model.Category{
		PID:       req.PID,
		Name:      req.Name,
		Alias:     req.Alias,
		Desc:      req.Desc,
		ShowInNav: req.ShowInNav,
		IsShow:    req.IsShow,
		IsParent:  req.IsParent,
		State:     req.State,
		IsDelete:  2,
	})
	if err != nil {
		api.JSON(c, err)
		return
	}
	api.JSON(c, errs.ErrSuccess, map[string]uint64{"cat_id": id})
	return
}

//DeleteCategory - Delete category By BrandID
func DeleteCategory(c *gin.Context) {
	var req categoryRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		api.JSON(c, errs.ErrParamVerify, nil)
		return
	}
	err := service.NewCategoryService().Delete(req.CatID)
	if err != nil {
		api.JSON(c, err)
		return
	}
	api.JSON(c, errs.ErrSuccess)
	return
}

//UpdateCategory - Update category By BrandID
func UpdateCategory(c *gin.Context) {
	var req updateCategoryRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		api.JSON(c, errs.ErrParamVerify, nil)
		return
	}
	err := service.NewCategoryService().Update(&model.Category{
		CatID:     req.CatID,
		PID:       req.PID,
		Name:      req.Name,
		Alias:     req.Alias,
		Desc:      req.Desc,
		ShowInNav: req.ShowInNav,
		IsShow:    req.IsShow,
		IsParent:  req.IsParent,
		IsDelete:  2,
	})

	if err != nil {
		api.JSON(c, err)
		return
	}
	api.JSON(c, errs.ErrSuccess)
	return
}

//GetCategory - Get category By BrandID
func GetCategory(c *gin.Context) {
	var req categoryRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		api.JSON(c, errs.ErrParamVerify, nil)
		return
	}
	data, err := service.NewCategoryService().FindByID(req.CatID)
	if err != nil {
		api.JSON(c, err)
		return
	}
	api.JSON(c, errs.ErrSuccess, data)
	return
}

//ListSubCategory - Get all subcategory List By pid
func ListSubCategory(c *gin.Context) {
	var req categoryRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		api.JSON(c, errs.ErrParamVerify, nil)
		return
	}
	data, err := service.NewCategoryService().FindAllByPID(req.CatID)
	if err != nil {
		api.JSON(c, err)
		return
	}
	api.JSON(c, errs.ErrSuccess, data)
	return
}

//ListCategory - Get all category List
func ListCategory(c *gin.Context) {
	var req listCategoryRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		api.JSON(c, errs.ErrParamVerify, nil)
		return
	}

	querys := []string{"is_delete = (?)"}
	args := []interface{}{2}
	//分类ID
	if req.CatID > 0 {
		querys = append(querys, "cat_id = (?)")
		args = append(args, req.CatID)
	}
	if len(req.Name) > 0 {
		querys = append(querys, "name = (?)")
		args = append(args, req.Name)
	}
	if req.PageSize <= 0 || req.PageSize > 20 {
		req.PageSize = 20
	}
	query := ""
	if len(querys) > 0 && len(args) > 0 {
		query = strings.Join(querys, " AND ")
	}
	fmt.Println(query)
	data, err := service.NewCategoryService().FindAll(query, args, "", req.PageNo, req.PageSize)
	if err != nil {
		api.JSON(c, err)
		return
	}
	api.JSON(c, errs.ErrSuccess, data)
	return
}

//GetCategoryAttrs - Get Product attrs
func GetCategoryAttrs(c *gin.Context) {
	var req getProductAttrsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		api.JSON(c, errs.ErrParamVerify, nil)
		return
	}
	data, _ := service.NewProductService().BeforeCreateByCatID(req.CatID)
	api.JSON(c, errs.ErrSuccess, data)
	return
}
