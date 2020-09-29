package v1

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

//添加文章
func AddArt(c *gin.Context) {
	var data model.Article
	c.ShouldBindJSON(&data)
	code := model.CreateAtr(&data)
	c.JSON(http.StatusOK, gin.H{
		"cde":  code,
		"msg":  errmsg.GetErrMsg(code),
		"data": data,
	})
}

//查询分类下所有文章
func GetCateArt(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data, code, total := model.GetCateArt(id, pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"code":  code,
		"msg":   errmsg.GetErrMsg(code),
		"total": total,
		"data":  data,
	})
}

//查询单个文章信息
func GetArtInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := model.GetArtInfo(id)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errmsg.GetErrMsg(code),
		"data": data,
	})
}

//查询文章列表
func GetArt(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data, code, total := model.GetAtr(pageSize, pageNum)

	c.JSON(http.StatusOK, gin.H{
		"code":  code,
		"msg":   errmsg.GetErrMsg(code),
		"total": total,
		"data":  data,
	})
}

//编辑文章
func EditArt(c *gin.Context) {
	var data model.Article
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&data)
	code := model.EditArt(id, &data)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errmsg.GetErrMsg(code),
	})

}

//删除文章
func DeleteArt(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	//软删除,不会被读取,但数据库中依然存在
	code := model.DeleteArt(id)
	if code != errmsg.SUCCESS {
		log.Printf("删除未成功")
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errmsg.GetErrMsg(code),
	})
}
