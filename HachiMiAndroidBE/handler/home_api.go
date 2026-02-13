package handler

import (
	"ProjectAndroidTest/model"
	"ProjectAndroidTest/pkg"
	"ProjectAndroidTest/service"
	"context"

	"github.com/cloudwego/hertz/pkg/app"
)

func HomePageHandler(ctx context.Context, c *app.RequestContext) {
	data, err := service.HomePageHandler()
	//获得主页数据出错
	if err.Status != "200" {
		c.JSON(500, pkg.FinalResponse{
			Status: "500",
			Info:   "fail to access home data",
			Data:   []model.HomeDTO{},
		})
		return
	}
	//获得主页数据
	c.JSON(200, pkg.FinalResponse{
		Status: "200",
		Info:   "success",
		Data:   data,
	})
}
func SearchHandler(ctx context.Context, c *app.RequestContext) {
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
	result, err := service.HomeSearchHandler(searchRequest)
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
	if len(result) == 0 {
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
		Data:   result,
	})
}
