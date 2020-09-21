package backend

import (
	"api.xinfos.com/api"
	"api.xinfos.com/internal/model"
	"api.xinfos.com/internal/service"
	"api.xinfos.com/utils/errs"
	"github.com/gin-gonic/gin"
)

type createBrandRequest struct {
	RequestID string `json:"request_id"`
	BrandName string `json:"name" binding:"required"`
	BrandLogo string `json:"logo" binding:"required"`
	BrandDesc string `json:"desc"`
	CatID     uint64 `json:"cat_id" binding:"required"`
	SortOrder uint32 `json:"displayorder"`
	IsShow    uint   `json:"is_show"`
}

type BrandRequest struct {
	RequestID string `json:"request_id"`
	BrandID   uint64 `json:"brand_id" binding:"required"`
}

type updateBrandRequest struct {
	RequestID string `json:"request_id"`
	BrandID   uint64 `json:"brand_id" binding:"required"`
	BrandName string `json:"name" binding:"required"`
	BrandLogo string `json:"logo" binding:"required"`
	BrandDesc string `json:"desc"`
	SortOrder uint32 `json:"displayorder"`
	IsShow    uint   `json:"is_show"`
}

type listBrandRequest struct {
	RequestID string `json:"request_id"`
	CatID     uint64 `json:"cat_id"`
	BrandName string `json:"name"`
	PageNo    uint   `json:"page_no"`
	PageSize  uint   `json:"page_size"`
}

//CreateBrand - Create a brand
func CreateBrand(c *gin.Context) {
	var req createBrandRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		api.JSON(c, errs.ErrParamVerify, nil)
		return
	}
	id, err := service.NewBrandService().Create(&model.Brand{
		BrandName: req.BrandName,
		BrandLogo: req.BrandLogo,
		BrandDesc: req.BrandDesc,
		CatID:     req.CatID,
		SortOrder: req.SortOrder,
		IsShow:    req.IsShow,
		IsDelete:  2,
	})
	if err != nil {
		api.JSON(c, err)
		return
	}
	api.JSON(c, errs.ErrSuccess, map[string]uint64{"brand_id": id})
	return
}

//DeleteBrand - Delete Brand By BrandID
func DeleteBrand(c *gin.Context) {
	var req BrandRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		api.JSON(c, errs.ErrParamVerify, nil)
		return
	}
	err := service.NewBrandService().Delete(req.BrandID)
	if err != nil {
		api.JSON(c, err)
		return
	}
	api.JSON(c, errs.ErrSuccess)
	return
}

//UpdateBrand - Update Brand By BrandID
func UpdateBrand(c *gin.Context) {
	var req updateBrandRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		api.JSON(c, errs.ErrParamVerify, nil)
		return
	}
	err := service.NewBrandService().Update(&model.Brand{
		BrandID:   req.BrandID,
		BrandName: req.BrandName,
		BrandLogo: req.BrandLogo,
		BrandDesc: req.BrandDesc,
		SortOrder: req.SortOrder,
		IsShow:    req.IsShow,
		IsDelete:  2,
	})

	if err != nil {
		api.JSON(c, err)
		return
	}
	api.JSON(c, errs.ErrSuccess)
	return
}

//GetBrand - Get Brand By BrandID
func GetBrand(c *gin.Context) {
	var req BrandRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		api.JSON(c, errs.ErrParamVerify, nil)
		return
	}
	data, err := service.NewBrandService().FindByID(req.BrandID)
	if err != nil {
		api.JSON(c, err)
		return
	}
	api.JSON(c, errs.ErrSuccess, data)
	return
}

//ListBrand - Get Brand List By CatID
func ListBrand(c *gin.Context) {
	var req listBrandRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		api.JSON(c, errs.ErrParamVerify, nil)
		return
	}

	//assemble query condition
	query := make(map[string]interface{}, 2)
	if req.CatID > 0 {
		query["cat_id"] = req.CatID
	}
	if len(req.BrandName) > 0 {
		query["brand_name"] = req.BrandName
	}

	if req.PageNo <= 0 {
		req.PageNo = 1
	}

	if req.PageSize <= 0 || req.PageSize > 20 {
		req.PageSize = 20
	}

	data, err := service.NewBrandService().FindAll(query, "", req.PageNo, req.PageSize)
	if err != nil {
		api.JSON(c, err)
		return
	}
	api.JSON(c, errs.ErrSuccess, data)
	return
}
