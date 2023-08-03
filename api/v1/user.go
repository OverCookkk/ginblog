package v1

import (
	"ginblog/model"
	response "ginblog/model/common"
	"ginblog/utils/errmsg"
	"ginblog/utils/validator"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// var code int

// 查询用户是否存在
func UserExist(c *gin.Context) {

}

// 添加用户
func AddUser(c *gin.Context) {
	var data model.User
	err := c.ShouldBindJSON(&data)
	if err != nil {
		return
	}
	msg, code := validator.Validate(&data)
	if code != errmsg.SUCCSE {
		response.ReturnWithMessage(code, msg, c)
		return
	}
	code = model.CheckUser(data.Username)
	if code == errmsg.SUCCSE {
		code = model.CreateUser(&data)
	} else if code == errmsg.ERROR_USERNAME_USER {
		code = errmsg.ERROR_USERNAME_USER
	}
	response.ReturnWithDetailed(code, data, errmsg.GetErrMsg(code), c)
}

// 查询单个用户

// 查询用户列表
func GetUsers(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data, total := model.GetUsers(pageSize, pageNum)
	code := errmsg.SUCCSE
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

// 编辑用户
func EditUser(c *gin.Context) {
	var data model.User
	ID, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&data)
	code := model.CheckUser(data.Username)
	if code == errmsg.SUCCSE {
		code = model.EditUser(ID, &data)
	}
	if code == errmsg.ERROR_PASSWORD_WRONG {
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 删除用户
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.DeleteUser(id)
	response.ReturnWithMessage(code, errmsg.GetErrMsg(code), c)
}
