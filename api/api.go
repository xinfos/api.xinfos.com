package api

import (
	"net/http"

	"api.xinfos.com/pkg/logger"
	"api.xinfos.com/utils/errs"

	"github.com/gin-gonic/gin"
)

//Response 公共返回结构
type Response struct {
	RequestID string      `json:"request_id"`
	Code      int         `json:"code"`
	Msg       string      `json:"msg"`
	Data      interface{} `json:"data"`
	Next      string      `json:"next"`
}

//JSON - Public interface return
func JSON(c *gin.Context, err *errs.Errs, data ...interface{}) {
	if err.ErrCode != 200 {
		logger.Error(err.ErrMsg)
	}

	var responseData interface{}
	if len(data) > 0 {
		responseData = data[0]
	}

	c.JSON(http.StatusOK, Response{
		Code: err.ErrCode,
		Msg:  err.ErrMsg,
		Data: responseData,
		Next: "",
	})
	return
}
