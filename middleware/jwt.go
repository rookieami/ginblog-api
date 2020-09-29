package middleware

import (
	"ginblog/utils"
	"ginblog/utils/errmsg"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

//认证中间件

var JwtKey = []byte(utils.Conf.JwtSecret)

type MyClaims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

//生成token

func SetToken(username string) (string, int) {
	expireTime := time.Now().Add(4 * time.Hour)
	SetClaims := MyClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), //有效时间
			Issuer:    "ginblon",         //签发人
		},
	}
	reqClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaims)

	token, err := reqClaims.SignedString(JwtKey)
	if err != nil {
		return "", errmsg.ERROR_USER_NOT_EXIST
	}
	return token, errmsg.SUCCESS

}

//验证token
func CheckToken(token string) (*MyClaims, int) {
	setToken, _ := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if Key, code := setToken.Claims.(*MyClaims); code && setToken.Valid {
		return Key, errmsg.SUCCESS
	} else {
		return nil, errmsg.ERROR
	}
}

//jwt中间件
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHerder := c.Request.Header.Get("Authorization")
		code := errmsg.ERROR
		if tokenHerder == "" {
			code = errmsg.ERROR_TOKEN_EXIST //不存在
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		checkToken := strings.SplitN(tokenHerder, " ", 2)
		if len(checkToken) != 2 && checkToken[0] != "Bearer" {
			code = errmsg.ERROR_TOKEN_TYPE_WRONG //格式错误
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		key, ctCode := CheckToken(checkToken[1])
		if ctCode == errmsg.ERROR {
			code = errmsg.ERROR_TOKEN_WRONG
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		if time.Now().Unix() > key.ExpiresAt {
			code = errmsg.ERROR_TOKEN_RUNTIME //已过期
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		c.Set("username", key.Username)
		c.Next()
	}
}
