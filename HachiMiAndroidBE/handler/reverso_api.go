package handler

import (
	"ProjectAndroidTest/model"
	"ProjectAndroidTest/pkg"
	"ProjectAndroidTest/service"
	"context"

	"github.com/cloudwego/hertz/pkg/app"
)

func LogonHandler(ctx context.Context, c *app.RequestContext) {
	var logonDTO model.LogonDTO
	err0 := c.BindJSON(&logonDTO)
	if err0 != nil {
		c.JSON(400, pkg.FinalResponse{
			Status: "400",
			Info:   "fail to receive json",
			Data:   nil,
		})
		return
	}
	response, err1 := service.LogonHandler(logonDTO)
	if err1.Status != "200" {
		c.JSON(400, pkg.FinalResponse{
			Status: "400",
			Info:   "fail to logon",
			Data:   nil,
		})
		return
	}
	c.JSON(200, pkg.FinalResponse{Status: "200", Info: "success", Data: response})
}
func RegisterHandler(ctx context.Context, c *app.RequestContext) {
	var registerDTO model.RegisterDTO
	err0 := c.BindJSON(&registerDTO)
	if err0 != nil {
		c.JSON(400, pkg.FinalResponse{
			Status: "400",
			Info:   "fail to receive json",
			Data:   nil,
		})
		return
	}

	err1 := service.RegisterHandler(registerDTO)
	if err1.Status != "200" {
		c.JSON(400, pkg.FinalResponse{
			Status: "400",
			Info:   "fail to register",
			Data:   nil,
		})
		return
	}
	c.JSON(200, pkg.FinalResponse{Status: "200", Info: "success", Data: nil})
}
func DeregisterHandler(ctx context.Context, c *app.RequestContext) {
	var logonDTO model.LogonDTO
	err0 := c.BindJSON(&logonDTO)
	if err0 != nil {
		c.JSON(400, pkg.FinalResponse{
			Status: "400",
			Info:   "fail to receiver json",
			Data:   nil,
		})
		return
	}

	err1 := service.DeregisterHandler(logonDTO)
	if err1.Status != "200" {
		c.JSON(400, pkg.FinalResponse{
			Status: "400",
			Info:   "Fail to deregister",
			Data:   nil,
		})
		return
	}
	c.JSON(200, pkg.FinalResponse{Status: "200", Info: "success", Data: nil})
}
func ChangeProfileHandler(ctx context.Context, c *app.RequestContext) {
	var changeProfileDto model.ChangeProfileDTO
	err0 := c.BindJSON(&changeProfileDto)
	if err0 != nil {
		c.JSON(400, pkg.FinalResponse{
			Status: "400",
			Info:   "fail to receiver json",
			Data:   nil,
		})
		return
	}
	err1 := service.ChangeProfileHandler(changeProfileDto)
	if err1.Status != "200" {
		c.JSON(400, pkg.FinalResponse{
			Status: "400",
			Info:   err1.Info,
			Data:   nil,
		})
		return
	}
	c.JSON(200, pkg.FinalResponse{Status: "200", Info: "success", Data: nil})

}
func JWTtestHandler(ctx context.Context, c *app.RequestContext) {
	c.JSON(200, pkg.FinalResponse{Status: "200", Info: "success", Data: nil})
}
