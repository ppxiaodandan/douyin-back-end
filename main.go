package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ppxiaodandan/douyin-back-end/dao"
	"github.com/ppxiaodandan/douyin-back-end/service"
)

func main() {
	go service.RunMessageServer()

	r := gin.Default()

	initRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	err := dao.InitMysql()
	if err != nil {
		panic(err)
	}

	// 延时关闭连接
	defer dao.DB.Close()
}
