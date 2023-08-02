package middleware

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func CasbinHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// tokenHerder := c.Request.Header.Get("Authorization")
		tokenHerder := ctx.Request.Header.Get("Authorization")
		// code = errmsg.SUCCSE
		if tokenHerder == "" {
			code = errmsg.ERROR_TOKEN_EXIST
			ctx.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			ctx.Abort()
			return
		}
		checkToken := strings.SplitN(tokenHerder, " ", 2)
		if len(checkToken) != 2 && checkToken[0] != "Bearer" {
			code = errmsg.ERROR_TOKEN_TYPE_WRONG
			ctx.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			ctx.Abort()
			return
		}
		key, Tcode := CheckToken(checkToken[1])
		if Tcode == errmsg.ERROR {
			code = errmsg.ERROR_TOKEN_WRONG
			ctx.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			ctx.Abort()
			return
		}

		sub := key.Username
		obj := ctx.Request.URL.Path // 如果Path有前缀，还需要处理
		act := ctx.Request.Method
		e := model.Casbin()
		success, _ := e.Enforce(sub, obj, act)
		if !success {
			// 权限不足
			// fmt.Printf("err : %s", err)
			ctx.JSON(http.StatusOK, gin.H{
				"code":    errmsg.ERROR_USER_NO_RIGHT,
				"message": errmsg.GetErrMsg(errmsg.ERROR_USER_NO_RIGHT),
			})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
