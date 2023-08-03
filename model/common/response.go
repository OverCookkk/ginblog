package response

import (
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func Result(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func Ok(c *gin.Context) {
	Result(errmsg.SUCCSE, map[string]interface{}{}, errmsg.GetErrMsg(errmsg.SUCCSE), c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(errmsg.SUCCSE, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(errmsg.SUCCSE, data, "查询成功", c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(errmsg.SUCCSE, data, message, c)
}

func Fail(c *gin.Context) {
	Result(errmsg.ERROR, map[string]interface{}{}, "操作失败", c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(errmsg.ERROR, map[string]interface{}{}, message, c)
}

// 返回data结构体
func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(errmsg.ERROR, data, message, c)
}

func ReturnWithDetailed(code int, data interface{}, message string, c *gin.Context) {
	Result(code, data, message, c)
}

func ReturnWithMessage(code int, message string, c *gin.Context) {
	Result(code, map[string]interface{}{}, message, c)
}
