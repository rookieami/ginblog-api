package v1

import (
	"fmt"
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

//上传文件

func Upload(c *gin.Context) {
	code := errmsg.INVALID_PARAMS
	file, err := c.FormFile("file")
	if err != nil {
		fmt.Println("取得时候就出了错", err)
		code = errmsg.ERROR
	}
	dst := fmt.Sprintf("./upload/%s", file.Filename)

	err = c.SaveUploadedFile(file, dst)
	if err != nil {
		log.Printf("err%v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": code,
			"msg":  errmsg.GetErrMsg(code),
		})
	} else {
		code = errmsg.SUCCESS
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  errmsg.GetErrMsg(code),
		})
	}

}
