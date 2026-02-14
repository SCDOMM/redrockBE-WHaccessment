package handler

import (
	"ProjectAndroidTest/model"
	"ProjectAndroidTest/pkg"
	"ProjectAndroidTest/service"
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
)

func AuthorityCheckHandler(c *app.RequestContext) int {
	roleVal, exist := c.Get("userRole")
	role, ok := roleVal.(uint)
	if !exist || !ok {
		c.JSON(400, pkg.FinalResponse{
			Status: "400",
			Info:   "JSON解析错误!",
			Data:   nil,
		})
		return 0
	}
	if role != 1 {
		c.JSON(400, pkg.FinalResponse{
			Status: "400",
			Info:   "您没有权限进行此操作!",
			Data:   nil,
		})
		return 0
	}
	return 1
}

func DeleteHomeHandler(ctx context.Context, c *app.RequestContext) {
	code := AuthorityCheckHandler(c)
	if code == 0 {
		return
	}

	var homeDto model.HomeDTO
	err0 := c.BindJSON(&homeDto)
	fmt.Println(homeDto)
	if err0 != nil {
		c.JSON(400, pkg.FinalResponse{
			Status: "400",
			Info:   err0.Error(),
			Data:   nil,
		})
		return
	}
	err1 := service.DeleteHomeHandler(homeDto)
	if err1.Status != "200" {
		c.JSON(400, pkg.FinalResponse{
			Status: "400",
			Info:   err1.Error(),
			Data:   nil,
		})
		return
	}
	c.JSON(200, pkg.FinalResponse{
		Status: "200",
		Info:   "success!",
		Data:   nil,
	})
}
func HomeUploadHandler(ctx context.Context, c *app.RequestContext) {
	code := AuthorityCheckHandler(c)
	if code == 0 {
		return
	}

	receiver := model.HomeDTO{}
	err := c.BindJSON(&receiver)
	if err != nil {
		c.JSON(400, pkg.FinalResponse{
			Status: "400",
			Info:   "fail to bind data",
			Data:   nil,
		})
		return
	}
	err1 := service.HomeUploadHandler(receiver)
	if err1.Status != "200" {
		c.JSON(500, pkg.FinalResponse{
			Status: "500",
			Info:   "fail to upload data",
			Data:   nil,
		})
		return
	}
	c.JSON(200, pkg.FinalResponse{
		Status: "200",
		Info:   "success",
		Data:   nil,
	})
}
func DeleteChatHandler(ctx context.Context, c *app.RequestContext) {
	code := AuthorityCheckHandler(c)
	if code == 0 {
		return
	}

	var dynamicDto model.DynamicDTO
	err0 := c.BindJSON(&dynamicDto)
	if err0 != nil {
		c.JSON(400, pkg.FinalResponse{
			Status: "400",
			Info:   err0.Error(),
			Data:   nil,
		})
		return
	}
	err1 := service.DeleteChatHandler(dynamicDto)
	if err1.Status != "200" {
		c.JSON(400, pkg.FinalResponse{
			Status: "400",
			Info:   err1.Error(),
			Data:   nil,
		})
		return
	}
	c.JSON(200, pkg.FinalResponse{
		Status: "200",
		Info:   "success!",
		Data:   nil,
	})
}
