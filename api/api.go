package api

import (
	"net/http"

	"api.xinfos.com/pkg/logger"
	"api.xinfos.com/utils/errs"

	"github.com/gin-gonic/gin"
)

type response struct {
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
	if data == nil {
		data = make([]interface{}, 0)
	}
	c.JSON(http.StatusOK, response{
		Code: err.ErrCode,
		Msg:  err.ErrMsg,
		Data: data,
		Next: "",
	})
	return
}
