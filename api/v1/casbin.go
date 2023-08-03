package v1

import (
	"ginblog/model"
	response "ginblog/model/common"
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func AddPolicy(c *gin.Context) {
	var cmr model.CasbinInReceive
	err := c.ShouldBindJSON(&cmr)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	code := errmsg.SUCCSE
	err = model.AddPolicy(cmr.AuthorityName, cmr.CasbinInfos)
	if err != nil {
		logrus.Infof("AddPolicy faild, err :%v", err)
		code = errmsg.ERROR_ADD_POLICY
	}
	response.ReturnWithMessage(code, errmsg.GetErrMsg(code), c)
}

func AddGroupingPolicy(c *gin.Context) {
	var cmr model.CasbinInReceive
	err := c.ShouldBindJSON(&cmr)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	code := errmsg.SUCCSE
	err = model.AddGroupingPolicy(cmr.AuthorityName, cmr.Group)
	if err != nil {
		logrus.Infof("AddGroupingPolicy faild, err :%v", err)
		code = errmsg.ERROR_ADD_GROUPING_POLICY
	}
	response.ReturnWithMessage(code, errmsg.GetErrMsg(code), c)
}

func DeleteRoleForUser(c *gin.Context) {
	var cmr model.CasbinInReceive
	err := c.ShouldBindJSON(&cmr)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	code := errmsg.SUCCSE
	err = model.DeleteRoleForUser(cmr.AuthorityName, cmr.Group)
	if err != nil {
		logrus.Infof("AddGroupingPolicy faild, err :%v", err)
		code = errmsg.ERROR_ADD_GROUPING_POLICY
	}
	response.ReturnWithMessage(code, errmsg.GetErrMsg(code), c)
}
