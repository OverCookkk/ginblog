package v1

import (
	"ginblog/middleware"
	"ginblog/model"
	"ginblog/utils/errmsg"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var data model.User

	c.ShouldBindJSON(&data)

	// 验证用户账号密码
	var code int
	code = model.CheckLogin(data.Username, data.Password)

	// 账号密码正确，就生成token返回给前端
	var token string
	if code == errmsg.SUCCSE {
		token, code = middleware.BuildToken(data.Username)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"token":   token,
	})
}
