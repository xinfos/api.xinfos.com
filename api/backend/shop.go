package backend

import (
	"api.xinfos.com/api"
	"api.xinfos.com/internal/model"
	"api.xinfos.com/internal/service"
	"api.xinfos.com/utils/errs"
	"github.com/gin-gonic/gin"
)

type createShopRequest struct {
	RequestID string `json:"request_id"`
	SellerID  uint64 `json:"seller_id" binding:"required"`
	Name      string `json:"name" binding:"required"`
	Type      uint   `json:"type" binding:"required"`
	Location  string `json:"location"`
	Address   string `json:"address"`
	IsAgree   uint   `json:"is_agree" binding:"required"`
	Desc      string `json:"desc"`
	Logo      string `json:"logo"`
	URL       string `json:"url"`
	State     uint   `json:"state"`
}

type updateShopRequest struct {
	RequestID string `json:"request_id"`
	ShopID    uint64 `json:"shop_id" binding:"required"`
	SellerID  uint64 `json:"seller_id" binding:"required"`
	Name      string `json:"name" binding:"required"`
	Type      uint   `json:"type" binding:"required"`
	Location  string `json:"location"`
	Address   string `json:"address"`
	Desc      string `json:"desc"`
	Logo      string `json:"logo"`
	URL       string `json:"url"`
	State     uint   `json:"state"`
}

type getShopRequest struct {
	RequestID string `json:"request_id"`                   //请求ID
	SellerID  uint64 `json:"seller_id" binding:"required"` //卖家ID
	ShopID    uint64 `json:"shop_id" binding:"required"`   //店铺ID
}

type infoShopRequest struct {
	RequestID string `json:"request_id"`                   //请求ID
	SellerID  uint64 `json:"seller_id" binding:"required"` //卖家ID
}

type listShopRequest struct {
	RequestID string `json:"request_id"`
	SellerID  uint64 `json:"seller_id" binding:"required"`
	Name      string `json:"name"`
	PageNo    uint   `json:"page_no" binding:"required"`
	PageSize  uint   `json:"page_size"`
}

/**
* @api {post} /backend/shop/create 创建店铺接口
* @apiName CreateShop
* @apiGroup 后台管理接口/创建店铺
*
* @apiParam {String} request_id	请求ID
* @apiParam {Number} seller_id	卖家ID
* @apiParam {String} name		店铺名称
* @apiParam {String} desc		店铺描述
* @apiParam {String} logo		店铺Logo
* @apiParam {String} url		店铺官网
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
 */
func CreateShop(c *gin.Context) {

	var req createShopRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		api.JSON(c, errs.ErrParamVerify, nil)
		return
	}

	data, errsmsg := service.NewShopService().Create(&model.Shop{
		SellerID: req.SellerID,
		Name:     req.Name,
		Type:     req.Type,
		Location: req.Location,
		Address:  req.Address,
		IsAgree:  req.IsAgree,
		Desc:     req.Desc,
		Logo:     req.Logo,
		URL:      req.URL,
		State:    req.State,
	})
	if errsmsg != nil {
		api.JSON(c, errsmsg, nil)
		return
	}
	api.JSON(c, errs.ErrSuccess, map[string]uint64{"shop_id": data})
	return

}

/**
* @api {post} /backend/shop/delete 删除店铺接口
* @apiName DeleteShop
* @apiGroup 后台管理接口/删除店铺
*
* @apiParam {String} request_id	请求ID
* @apiParam {Number} seller_id	卖家ID
* @apiParam {Number} shop_id	店铺ID
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
 */
func DeleteShop(c *gin.Context) {

	var req getShopRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		api.JSON(c, errs.ErrParamVerify, nil)
		return
	}

	errsmsg := service.NewShopService().Delete(req.ShopID, req.SellerID)
	if errsmsg != nil {
		api.JSON(c, errsmsg, nil)
		return
	}
	api.JSON(c, errs.ErrSuccess, map[string]uint64{"shop_id": req.ShopID})
	return
}

/**
* @api {post} /backend/shop/update 更新店铺信息接口
* @apiName UpdateShop
* @apiGroup 后台管理接口/更新店铺信息
*
* @apiParam {String} request_id	请求ID
* @apiParam {Number} shop_id	店铺ID
* @apiParam {Number} seller_id	卖家ID
* @apiParam {String} name		店铺名称
* @apiParam {String} desc		店铺描述
* @apiParam {String} logo		店铺Logo
* @apiParam {String} url		店铺官网
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
 */
func UpdateShop(c *gin.Context) {
	var req updateShopRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		api.JSON(c, errs.ErrParamVerify, nil)
		return
	}

	errsmsg := service.NewShopService().Update(&model.Shop{
		ID:       req.ShopID,
		SellerID: req.SellerID,
		Name:     req.Name,
		Type:     req.Type,
		Location: req.Location,
		Address:  req.Address,
		Desc:     req.Desc,
		Logo:     req.Logo,
		URL:      req.URL,
		State:    req.State,
	})
	if errsmsg != nil {
		api.JSON(c, errsmsg, nil)
		return
	}
	api.JSON(c, errs.ErrSuccess, map[string]uint64{"shop_id": req.ShopID})
	return
}

/**
* @api {post} /backend/shop/get 获取指定店铺详情接口
* @apiName GetShop
* @apiGroup 后台管理接口/单个店铺详情
*
* @apiParam {String} request_id	请求ID
* @apiParam {Number} seller_id	卖家ID
* @apiParam {Number} shop_id	店铺ID
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
 */
func DashboardShop(c *gin.Context) {

	var req infoShopRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		api.JSON(c, errs.ErrParamVerify, nil)
		return
	}

	data, errsmsg := service.NewShopService().Dashboard(req.SellerID)
	if errsmsg != nil {
		api.JSON(c, errsmsg, nil)
		return
	}
	api.JSON(c, errs.ErrSuccess, data)
	return
}

/**
* @api {post} /backend/shop/get 获取指定店铺详情接口
* @apiName GetShop
* @apiGroup 后台管理接口/单个店铺详情
*
* @apiParam {String} request_id	请求ID
* @apiParam {Number} seller_id	卖家ID
* @apiParam {Number} shop_id	店铺ID
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
 */
func GetShop(c *gin.Context) {

	var req getShopRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		api.JSON(c, errs.ErrParamVerify, nil)
		return
	}

	data, errsmsg := service.NewShopService().FindByID(req.ShopID, req.SellerID)
	if errsmsg != nil {
		api.JSON(c, errsmsg, nil)
		return
	}
	api.JSON(c, errs.ErrSuccess, data)
	return
}

/**
* @api {post} /backend/shop/get 获取指定店铺详情接口
* @apiName GetShop
* @apiGroup 后台管理接口/单个店铺详情
*
* @apiParam {String} request_id	请求ID
* @apiParam {Number} seller_id	卖家ID
* @apiParam {Number} shop_id	店铺ID
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
 */

func InfoShop(c *gin.Context) {

	var req infoShopRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		api.JSON(c, errs.ErrParamVerify, nil)
		return
	}

	data, errsmsg := service.NewShopService().FindBySellerID(req.SellerID)
	if errsmsg != nil {
		api.JSON(c, errsmsg, nil)
		return
	}
	api.JSON(c, errs.ErrSuccess, data)
	return
}

/**
* @api {post} /backend/shop/list 获取用户店铺列表接口
* @apiName ListShop
* @apiGroup 后台管理接口/单个店铺详情
*
* @apiParam {String} request_id	请求ID
* @apiParam {Number} seller_id	卖家ID
* @apiParam {String} [name]		店铺名称
* @apiParam {Number} page_no	页数
* @apiParam {Number} page_size	每页显示个数
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
 */
func ListShop(c *gin.Context) {
	var req listShopRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		api.JSON(c, errs.ErrParamVerify, nil)
		return
	}

	query := []string{"seller_id = (?)", "is_delete = (?)"}
	args := []interface{}{req.SellerID, 2}

	//店铺名称
	if len(req.Name) > 0 {
		query = append(query, "name = (?)")
		args = append(args, req.Name)
	}
	if req.PageSize <= 0 || req.PageSize > 20 {
		req.PageSize = 20
	}

	data, err := service.NewShopService().FindAll(model.QueryArrayToString(query), args, "", req.PageNo, req.PageSize)
	if err != nil {
		api.JSON(c, err, nil)
		return
	}
	api.JSON(c, errs.ErrSuccess, data)
	return
}
