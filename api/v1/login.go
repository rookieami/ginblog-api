package v1

import (
	"ginblog/middleware"
	"ginblog/model"
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	var data model.User
	c.ShouldBindJSON(&data)
	var token string
	code := model.CheckLogin(data.Username, data.Password)
	//验证正确
	if code == errmsg.SUCCESS {
		//生成token
		token, code = middleware.SetToken(data.Username)
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  code,
		"msg":   errmsg.GetErrMsg(code),
		"token": token,
	})

}
