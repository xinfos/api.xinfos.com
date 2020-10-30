package backend

import (
	"api.xinfos.com/api"
	"api.xinfos.com/internal/model"
	"api.xinfos.com/internal/service"
	"api.xinfos.com/utils/errs"
	"github.com/gin-gonic/gin"
)

type createAttrRequest struct {
	Name        string `json:"name" binding:"required"`
	FillType    uint   `json:"fill_type"`
	IsRequired  uint   `json:"is_required"`
	IsNumeric   uint   `json:"is_numeric"`
	Unit        string `json:"unit"`
	IsGeneric   uint   `json:"is_generic"`
	IsSearching uint   `json:"is_searching"`
	Segments    string `json:"segments"`
}

type updateAttrRequest struct {
	AttrID      uint64 `json:"attr_id" binding:"required"`
	Name        string `json:"name" binding:"required"`
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

/**
* @api {post} /backend/attr/create 创建属性
* @apiName CreateAttr
*
* @apiGroup 商品属性管理
*
* @apiParam (公共参数) {String} [request_id] 请求ID
*
* @apiParam (请求参数) {String} Name      	  属性名称
* @apiParam (请求参数) {Number} [FillType = 2]   属性值填充类型 [1: 选项框 | 2: 输入框]
* @apiParam (请求参数) {Number} [IsRequired = 2] 属性是否为必填项 [1:是 | 2: 否]
* @apiParam (请求参数) {Number} [IsNumeric = 2]  属性值是否为数字类型 [1:是 | 2: 否]
* @apiParam (请求参数) {String} [Unit]  	  属性单位
* @apiParam (请求参数) {Number} [IsGeneric = 2]  是否是SPU通用属性 [1:是 | 2: 否]
* @apiParam (请求参数) {Number} [IsSearching = 2]是否用于搜索过滤 [1:是 | 2: 否]
* @apiParam (请求参数) {String} [Segments]   分割
*
* @apiSuccess (响应参数) {String} request_id 返回请求的ID.
* @apiSuccess (响应参数) {Number} code 返回CODE码
* @apiSuccess (响应参数) {String} msg 返回信息
* @apiSuccess (响应参数) {Object} data 返回具体内容
* @apiSuccess (响应参数) {Number} data.attr_id 属性ID
* @apiSuccess (响应参数) {String} next ""
*
* @apiSuccessExample 响应示例:
* {
*   "request_id": "c3b51fe3-0dc6-4b69-9a66-508ecc9a8633",
*   "code": 200,
*   "msg": "请求成功"
*   "data": {
*      "attr_id": 100031
*    },
*    "next": ""
* }
*
* @apiError (错误码) {Number} 180001 抱歉，创建失败!~
* @apiError (错误码) {Number} 180002 抱歉，删除失败!~
* @apiError (错误码) {Number} 180003 抱歉，删除失败，该属性被其他属性模板引用!~
* @apiError (错误码) {Number} 180004 抱歉，更新失败!~
* @apiError (错误码) {Number} 180005 抱歉，属性名已存在，请换一个名称重试!~
* @apiError (错误码) {Number} 180006 抱歉，抱歉，没有找到相关的属性信息!~
* @apiError (错误码) {Number} 180007 抱歉，由于属性为数值类型，单位为必填项!~
*
* @apiErrorExample 异常示例:
*     {
*       "request_id": "c3b51fe3-0dc6-4b69-9a66-508ecc9a8633",
*       "code": 180001,
*       "msg": "抱歉，创建失败!~"
*       "data": null,
*       "next": ""
*     }
 */
func CreateAttr(c *gin.Context) {
	var req createAttrRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		api.JSON(c, errs.ErrParamVerify, nil)
		return
	}

	data, err := service.NewAttrService().Create(&model.SAttr{
		Name:        req.Name,
		FillType:    req.FillType,
		IsRequired:  req.IsRequired,
		IsNumeric:   req.IsNumeric,
		Unit:        req.Unit,
		IsGeneric:   req.IsGeneric,
		IsSearching: req.IsSearching,
		Segments:    req.Segments,
		IsDelete:    2,
	})

	if err != nil {
		api.JSON(c, err)
		return
	}
	api.JSON(c, errs.ErrSuccess, map[string]uint64{"attr_id": data})
	return
}

/**
* @api {post} /backend/attr/delete 删除属性
* @apiName DeleteAttr
*
* @apiGroup 商品属性管理
*
* @apiParam (公共参数) {String} [request_id] 请求ID
*
* @apiParam (请求参数) {Number} attr_id 属性ID

*
* @apiSuccess (响应参数) {String} request_id 返回请求的ID.
* @apiSuccess (响应参数) {Number} code 返回CODE码
* @apiSuccess (响应参数) {String} msg 返回信息
* @apiSuccess (响应参数) {Object} data 返回具体内容
* @apiSuccess (响应参数) {Number} data.attr_id 属性ID
* @apiSuccess (响应参数) {String} next ""
*
* @apiSuccessExample 响应示例:
* {
*   "request_id": "c3b51fe3-0dc6-4b69-9a66-508ecc9a8633",
*   "code": 200,
*   "msg": "请求成功"
*   "data": {
*      "attr_id": 100031
*    },
*    "next": ""
* }
*
* @apiError (错误码) {Number} 180001 抱歉，创建失败!~
* @apiError (错误码) {Number} 180002 抱歉，删除失败!~
* @apiError (错误码) {Number} 180003 抱歉，删除失败，该属性被其他属性模板引用!~
* @apiError (错误码) {Number} 180004 抱歉，更新失败!~
* @apiError (错误码) {Number} 180005 抱歉，属性名已存在，请换一个名称重试!~
* @apiError (错误码) {Number} 180006 抱歉，抱歉，没有找到相关的属性信息!~
* @apiError (错误码) {Number} 180007 抱歉，由于属性为数值类型，单位为必填项!~
*
* @apiErrorExample 异常示例:
*     {
*       "request_id": "c3b51fe3-0dc6-4b69-9a66-508ecc9a8633",
*       "code": 180001,
*       "msg": "抱歉，创建失败!~"
*       "data": null,
*       "next": ""
*     }
 */
func DeleteAttr(c *gin.Context) {
	var req getAttrRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		api.JSON(c, errs.ErrParamVerify, nil)
		return
	}

	errsmsg := service.NewAttrService().Delete(req.AttrID)
	if errsmsg != nil {
		api.JSON(c, errsmsg)
		return
	}
	api.JSON(c, errs.ErrSuccess, map[string]uint64{"attr_id": req.AttrID})
	return
}

/**
* @api {post} /backend/attr/update 更新属性
* @apiName UpdateAttr
*
* @apiGroup 商品属性管理
*
* @apiParam (公共参数) {String} [request_id] 请求ID
*
* @apiParam (请求参数) {Number} attr_id 属性ID
* @apiParam (请求参数) {String} Name      	  属性名称
* @apiParam (请求参数) {Number} [FillType = 2]   属性值填充类型 [1: 选项框 | 2: 输入框]
* @apiParam (请求参数) {Number} [IsRequired = 2] 属性是否为必填项 [1:是 | 2: 否]
* @apiParam (请求参数) {Number} [IsNumeric = 2]  属性值是否为数字类型 [1:是 | 2: 否]
* @apiParam (请求参数) {String} [Unit]  	  属性单位
* @apiParam (请求参数) {Number} [IsGeneric = 2]  是否是SPU通用属性 [1:是 | 2: 否]
* @apiParam (请求参数) {Number} [IsSearching = 2]是否用于搜索过滤 [1:是 | 2: 否]
* @apiParam (请求参数) {String} [Segments]   分割
*
* @apiSuccess (响应参数) {String} request_id 返回请求的ID.
* @apiSuccess (响应参数) {Number} code 返回CODE码
* @apiSuccess (响应参数) {String} msg 返回信息
* @apiSuccess (响应参数) {Object} data 返回具体内容
* @apiSuccess (响应参数) {Number} data.attr_id 属性ID
* @apiSuccess (响应参数) {String} next ""
*
* @apiSuccessExample 响应示例:
* {
*   "request_id": "c3b51fe3-0dc6-4b69-9a66-508ecc9a8633",
*   "code": 200,
*   "msg": "请求成功"
*   "data": {
*      "attr_id": 100031
*    },
*    "next": ""
* }
*
* @apiError (错误码) {Number} 180001 抱歉，创建失败!~
* @apiError (错误码) {Number} 180002 抱歉，删除失败!~
* @apiError (错误码) {Number} 180003 抱歉，删除失败，该属性被其他属性模板引用!~
* @apiError (错误码) {Number} 180004 抱歉，更新失败!~
* @apiError (错误码) {Number} 180005 抱歉，属性名已存在，请换一个名称重试!~
* @apiError (错误码) {Number} 180006 抱歉，抱歉，没有找到相关的属性信息!~
* @apiError (错误码) {Number} 180007 抱歉，由于属性为数值类型，单位为必填项!~
*
* @apiErrorExample 异常示例:
*     {
*       "request_id": "c3b51fe3-0dc6-4b69-9a66-508ecc9a8633",
*       "code": 180001,
*       "msg": "抱歉，创建失败!~"
*       "data": null,
*       "next": ""
*     }
 */
func UpdateAttr(c *gin.Context) {
	var req updateAttrRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		api.JSON(c, errs.ErrParamVerify, nil)
		return
	}

	err := service.NewAttrService().Update(&model.SAttr{
		ID:          req.AttrID,
		Name:        req.Name,
		FillType:    req.FillType,
		IsRequired:  req.IsRequired,
		IsNumeric:   req.IsNumeric,
		Unit:        req.Unit,
		IsGeneric:   req.IsGeneric,
		IsSearching: req.IsSearching,
		Segments:    req.Segments,
		IsDelete:    2,
	})

	if err != nil {
		api.JSON(c, err)
		return
	}
	api.JSON(c, errs.ErrSuccess, map[string]uint64{"attr_id": req.AttrID})
	return
}

/**
* @api {post} /backend/attr/get 获取属性信息
* @apiName GetAttr
*
* @apiGroup 商品属性管理
*
* @apiParam (公共参数) {String} [request_id] 请求ID
*
* @apiParam (请求参数) {Number} attr_id 属性ID
*
* @apiSuccess (响应参数) {String} request_id 返回请求的ID.
* @apiSuccess (响应参数) {Number} code 返回CODE码
* @apiSuccess (响应参数) {String} msg 返回信息
* @apiSuccess (响应参数) {Object} data 返回具体内容
* @apiSuccess (响应参数) {Number} data.attr_id 属性ID
* @apiSuccess (响应参数) {String} next ""
*
* @apiSuccessExample 响应示例:
* {
*   "request_id": "c3b51fe3-0dc6-4b69-9a66-508ecc9a8633",
*   "code": 200,
*   "msg": "请求成功"
*   "data": {
*      "attr_id": 100031,
*    },
*    "next": ""
* }
*
* @apiError (错误码) {Number} 180001 抱歉，创建失败!~
* @apiError (错误码) {Number} 180002 抱歉，删除失败!~
* @apiError (错误码) {Number} 180003 抱歉，删除失败，该属性被其他属性模板引用!~
* @apiError (错误码) {Number} 180004 抱歉，更新失败!~
* @apiError (错误码) {Number} 180005 抱歉，属性名已存在，请换一个名称重试!~
* @apiError (错误码) {Number} 180006 抱歉，抱歉，没有找到相关的属性信息!~
* @apiError (错误码) {Number} 180007 抱歉，由于属性为数值类型，单位为必填项!~
*
* @apiErrorExample 异常示例:
*     {
*       "request_id": "c3b51fe3-0dc6-4b69-9a66-508ecc9a8633",
*       "code": 180001,
*       "msg": "抱歉，创建失败!~"
*       "data": null,
*       "next": ""
*     }
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
* @api {post} /backend/attr/list 获取属性列表信息
* @apiName ListAttr
*
* @apiGroup 商品属性管理
*
* @apiParam (公共参数) {String} [request_id] 请求ID
*
* @apiParam (请求参数) {Number} attr_id 属性ID
* @apiParam (请求参数) {String} Name      属性名称
* @apiParam (请求参数) {Number} page_no   页数
* @apiParam (请求参数) {Number} page_size 每页显示个数
*
* @apiSuccess (响应参数) {String} request_id 返回请求的ID.
* @apiSuccess (响应参数) {Number} code 返回CODE码
* @apiSuccess (响应参数) {String} msg 返回信息
* @apiSuccess (响应参数) {Object} data 返回具体内容
* @apiSuccess (响应参数) {Object[]} data.list 属性列表
* @apiSuccess (响应参数) {Object[]} data.list.id 属性ID
* @apiSuccess (响应参数) {Object[]} data.list.name 属性名称
* @apiSuccess (响应参数) {Object[]} data.list.fill_type 属性值填充类型 [1: 选项框 | 2: 输入框]
* @apiSuccess (响应参数) {Object[]} data.list.is_required 属性是否为必填项 [1:是 | 2: 否]
* @apiSuccess (响应参数) {Object[]} data.list.is_numeric 属性值是否为数字类型 [1:是 | 2: 否]
* @apiSuccess (响应参数) {Object[]} data.list.unit 属性单位
* @apiSuccess (响应参数) {Object[]} data.list.is_generic 是否是SPU通用属性 [1:是 | 2: 否]
* @apiSuccess (响应参数) {Object[]} data.list.is_searching 是否用于搜索过滤 [1:是 | 2: 否]
* @apiSuccess (响应参数) {Object[]} data.list.segments 分割
* @apiSuccess (响应参数) {Number} data.current_page_no 当前页数
* @apiSuccess (响应参数) {Number} data.current_page_size 每页显示个数
* @apiSuccess (响应参数) {Number} data.total_count 列表总个数
* @apiSuccess (响应参数) {String} next ""
*
* @apiSuccessExample 响应示例:
* {
*   "request_id": "c3b51fe3-0dc6-4b69-9a66-508ecc9a8633",
*   "code": 200,
*   "msg": "请求成功"
*   "data": {
*      "attr_id": 100031,
*    },
*    "next": ""
* }
*
* @apiError (错误码) {Number} 180001 抱歉，创建失败!~
* @apiError (错误码) {Number} 180002 抱歉，删除失败!~
* @apiError (错误码) {Number} 180003 抱歉，删除失败，该属性被其他属性模板引用!~
* @apiError (错误码) {Number} 180004 抱歉，更新失败!~
* @apiError (错误码) {Number} 180005 抱歉，属性名已存在，请换一个名称重试!~
* @apiError (错误码) {Number} 180006 抱歉，抱歉，没有找到相关的属性信息!~
* @apiError (错误码) {Number} 180007 抱歉，由于属性为数值类型，单位为必填项!~
*
* @apiErrorExample 异常示例:
*     {
*       "request_id": "c3b51fe3-0dc6-4b69-9a66-508ecc9a8633",
*       "code": 180001,
*       "msg": "抱歉，创建失败!~"
*       "data": null,
*       "next": ""
*     }
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
