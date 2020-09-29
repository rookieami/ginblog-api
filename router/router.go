package router

import (
	"ginblog/api/v1"
	"ginblog/middleware"
	"ginblog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	//级别
	gin.SetMode(utils.Conf.RunMode)
	r := gin.New()
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())
	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		//用户模块的路由接口
		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("user/:id", v1.DeleteUser)
		//分类模块的路由接口
		auth.POST("category/add", v1.AddCategory)
		auth.PUT("category/:id", v1.EditCategory)
		auth.DELETE("category/:id", v1.DeleteCategory)
		//文章模块的路由接口
		auth.POST("article/add", v1.AddArt)
		auth.PUT("article/:id", v1.EditArt)
		auth.DELETE("article/:id", v1.DeleteArt)

		//上传文件
		auth.POST("/upload", v1.Upload)
	}
	router := r.Group("api/v1")
	{
		router.POST("login", v1.Login)
		router.POST("user/add", v1.AddUser)
		router.GET("user", v1.GetUsers)

		router.GET("category", v1.GetCategory)
		router.GET("article", v1.GetArt)
		router.GET("article/info/:id", v1.GetArtInfo)
		router.GET("article/list/:id", v1.GetCateArt)
	}
	return r
}
