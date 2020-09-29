package v1

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"ginblog/utils/validate"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

//查询用户是否存在
func UserExist(c *gin.Context) {

}

//添加用户
func AddUser(c *gin.Context) {
	var data model.User
	var msg string
	_ = c.ShouldBindJSON(&data)
	msg, code := validate.Validate(&data)
	if code != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  msg,
		})
		return
	}

	code = model.CheckUser(data.Username)
	if code == errmsg.SUCCESS {
		model.CreateUser(&data)
	}
	if code == errmsg.ERROR_USERNAME_USED {
		code = errmsg.ERROR_USERNAME_USED
	}
	c.JSON(http.StatusOK, gin.H{
		"cde":  code,
		"msg":  errmsg.GetErrMsg(code),
		"data": data,
	})
}

//查询单个用户

//查询多个用户
func GetUsers(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data, code, total := model.GetUsers(pageSize, pageNum)
	if code == errmsg.INVALID_PARAMS {
		code = errmsg.INVALID_PARAMS
	}
	if code == errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"code":  code,
			"msg":   errmsg.GetErrMsg(code),
			"total": total,
			"data":  data,
		})
	}

}

//编辑用户
func EditUser(c *gin.Context) {
	//修改不允许用户重名
	//查询是否重名
	var data model.User
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&data)

	code := model.CheckUser(data.Username)
	if code == errmsg.SUCCESS {
		//可以使用
		code = model.EditUser(id, &data)
	} else if code == errmsg.ERROR_USERNAME_USED {
		//重名
		code = errmsg.ERROR_USERNAME_USED
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errmsg.GetErrMsg(code),
	})

}

//删除用户
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	//软删除,不会被读取,但数据库中依然存在
	code := model.DeleteUser(id)
	if code != errmsg.SUCCESS {
		log.Printf("删除未成功")
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errmsg.GetErrMsg(code),
	})
}
