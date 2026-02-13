package main

import (
	"ProjectAndroidTest/dao"
	"ProjectAndroidTest/router"

	"github.com/cloudwego/hertz/pkg/app/server"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	dao.InitDataBase()
	//dao.ChatAddTestData()
	//dao.HomeAddTestData()
	//dao.ReversoAddTestData()
	host := "0.0.0.0"
	port := "8080"
	h := server.Default(server.WithHostPorts(host+":"+port), server.WithMaxRequestBodySize(20*1024*1024))
	router.InitRouter(h)
	h.Spin()
}
