package v1

import (
	"ginblog/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateCasbin(c *gin.Context) {
	var cmr model.CasbinInReceive
	err := c.ShouldBindJSON(&cmr)
	if err != nil {
		// response.FailWithMessage(err.Error(), c)
		return
	}
	err = model.UpdateCasbin(cmr.AuthorityName, cmr.CasbinInfos)

	c.JSON(http.StatusOK, gin.H{
		// TODO
		"status": 9999,
		// "message": errmsg.GetErrMsg(code),
		// "url":     url,
	})
}

func UpdateGroupCasbin(c *gin.Context) {
	var cmr model.CasbinInReceive
	err := c.ShouldBindJSON(&cmr)
	if err != nil {
		// response.FailWithMessage(err.Error(), c)
		return
	}
	err = model.UpdateGroupCasbin(cmr.AuthorityName, cmr.Group)

	c.JSON(http.StatusOK, gin.H{
		// TODO
		"status": 9999,
		// "message": errmsg.GetErrMsg(code),
		// "url":     url,
	})
}

func DeleteRoleForUser(c *gin.Context) {
	var cmr model.CasbinInReceive
	err := c.ShouldBindJSON(&cmr)
	if err != nil {
		// response.FailWithMessage(err.Error(), c)
		return
	}
	err = model.DeleteRoleForUser(cmr.AuthorityName, cmr.Group)

	c.JSON(http.StatusOK, gin.H{
		// TODO
		"status": 9999,
		// "message": errmsg.GetErrMsg(code),
		// "url":     url,
	})
}
