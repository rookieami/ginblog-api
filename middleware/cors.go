package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

//跨域中间件
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		cors.New(cors.Config{
			//AllowAllOrigins: true,//允许所有跨域
			AllowOrigins:  []string{"*"}, //允许跨域域名
			AllowMethods:  []string{"*"}, //跨域方法
			AllowHeaders:  []string{"Origin"},
			ExposeHeaders: []string{"Content-Length", "Authorization"},
			//AllowCredentials: true,
			//AllowOriginFunc: func(origin string) bool {
			//	return origin == "https://github.com"
			//},
			MaxAge: 12 * time.Hour,
		})

	}
}
