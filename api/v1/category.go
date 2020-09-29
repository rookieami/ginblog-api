package v1

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

//添加分类
func AddCategory(c *gin.Context) {
	var data model.Category
	c.ShouldBindJSON(&data)
	code := model.CheckCategory(data.Name)
	if code == errmsg.SUCCESS {
		model.CreateCategory(&data)
	}
	if code == errmsg.ERROR_CATEGORYNAME_USED {
		code = errmsg.ERROR_CATEGORYNAME_USED
	}
	c.JSON(http.StatusOK, gin.H{
		"cde":  code,
		"msg":  errmsg.GetErrMsg(code),
		"data": data,
	})
}

//查询分类列表
func GetCategory(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data, code, total := model.GetCategory(pageSize, pageNum)
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

//编辑分类
func EditCategory(c *gin.Context) {
	//修改不允许用户重名
	//查询是否重名
	var data model.Category
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&data)

	code := model.CheckCategory(data.Name)
	if code == errmsg.SUCCESS {
		//可以使用
		code = model.EditCategory(id, &data)
	} else if code == errmsg.ERROR_CATEGORYNAME_USED {
		//重名
		code = errmsg.ERROR_CATEGORYNAME_USED
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errmsg.GetErrMsg(code),
	})

}

//删除分类
func DeleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	//软删除,不会被读取,但数据库中依然存在
	code := model.DeleteCategory(id)
	if code != errmsg.SUCCESS {
		log.Printf("删除未成功")
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errmsg.GetErrMsg(code),
	})
}
