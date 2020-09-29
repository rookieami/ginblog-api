package main

import (
	"fmt"
	"ginblog/model"
	"ginblog/router"
	"ginblog/utils"
	"net/http"
)

func main() {
	//初始化数据库
	model.Init()
	//服务初始化
	r := router.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", utils.Conf.HTTPPort),
		Handler:        r,
		ReadTimeout:    utils.Conf.ReadTimeout,
		WriteTimeout:   utils.Conf.WriteTimeout,
		MaxHeaderBytes: 8 << 20,
	}
	err := s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
