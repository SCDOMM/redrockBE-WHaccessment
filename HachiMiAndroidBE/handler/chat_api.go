package handler

import (
	"ProjectAndroidTest/model"
	"ProjectAndroidTest/pkg"
	"ProjectAndroidTest/service"
	"context"
	"log"

	"github.com/cloudwego/hertz/pkg/app"
)

func ChatPageHandler(ctx context.Context, c *app.RequestContext) {
	dynamic, err := service.ChatPageHandler()
	//获得主页数据出错
	if err.Status != "200" {
		c.JSON(500, pkg.FinalResponse{
			Status: "500",
			Info:   "fail to access dynamic data",
			Data:   nil,
		})
		return
	}
	//获得主页数据
	c.JSON(200, pkg.FinalResponse{
		Status: "200",
		Info:   "success",
		Data:   dynamic,
	})
}
func ChatSearchHandler(ctx context.Context, c *app.RequestContext) {
	keyWords := c.Param("keyWords")
	if len(keyWords) == 0 {
		c.JSON(200, pkg.FinalResponse{
			Status: "200",
			Info:   "quit search",
			Data:   nil,
		})
		return
	}
	searchRequest := model.SearchDTO{
		Content: keyWords,
	}
	dynamic, err := service.ChatSearchHandler(searchRequest)
	//搜索出错
	if err.Status != "200" {
		c.JSON(500, pkg.FinalResponse{
			Status: "500",
			Info:   "fail to access search data",
			Data:   nil,
		})
		return
	}
	//搜索结果为空
	if len(dynamic) == 0 {
		c.JSON(200, pkg.FinalResponse{
			Status: "200",
			Info:   "empty search result",
			Data:   nil,
		})
		return
	}
	//有搜索结果
	c.JSON(200, pkg.FinalResponse{
		Status: "200",
		Info:   "success",
		Data:   dynamic,
	})
}
func ChatUploadHandler(ctx context.Context, c *app.RequestContext) {
	receiver := model.DynamicUploadDTO{}
	err := c.BindJSON(&receiver)
	if err != nil {
		c.JSON(400, pkg.FinalResponse{
			Status: "400",
			Info:   "fail to bind data",
			Data:   nil,
		})
		log.Println(err)
		return
	}
	err1 := service.ChatUploadHandler(receiver)
	if err1.Status != "200" {
		c.JSON(500, pkg.FinalResponse{
			Status: "500",
			Info:   "fail to upload data",
			Data:   nil,
		})
		log.Println(err1)
		return
	}
	c.JSON(200, pkg.FinalResponse{
		Status: "200",
		Info:   "success",
		Data:   nil,
	})
}
