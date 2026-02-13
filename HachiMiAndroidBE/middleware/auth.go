package utils

import (
	"ProjectAndroidTest/pkg"
	"context"
	"net/http"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
)

func LoggerMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		//检查请求头
		authHeader := c.Request.Header.Get("Authorization")
		if len(authHeader) == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, pkg.FinalResponse{
				Status: "401",
				Info:   "unauthorized",
				Data:   nil,
			})
			return
		}
		//验证格式
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, pkg.FinalResponse{
				Status: "401",
				Info:   "not a auth format",
				Data:   nil,
			})
			return
		}
		//验证Access有效性
		tokenString := parts[1]
		claims, err := pkg.VerifyAccessToken(tokenString)
		if err == nil {
			//信息存入Gin Context，相信后人的智慧
			c.Set("userAccount", claims.Account)
			c.Set("userRole", claims.Role)
			c.Next(ctx)
			return
		}
		//检验refresh有效性
		refreshToken := c.Cookie("refresh_token")
		if string(refreshToken) == "" {
			//你根本不是refresh，你到底在哪
			c.AbortWithStatusJSON(http.StatusUnauthorized, pkg.FinalResponse{
				Status: "401",
				Info:   "refreshToken required",
				Data:   nil,
			})
			return
		}
		claims, err = pkg.VerifyRefreshToken(string(refreshToken))
		if err != nil {
			//过期了
			c.AbortWithStatusJSON(http.StatusUnauthorized, pkg.FinalResponse{
				Status: "401",
				Info:   "invalid refreshToken",
				Data:   nil,
			})
			return
		}
		newAccessToken, err := pkg.CreateAccessToken(claims.Account, claims.Role)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, pkg.FinalResponse{
				Status: "500",
				Info:   "fail to create new refreshToken",
				Data:   nil,
			})
			return
		}
		c.Set("userAccount", claims.Account)
		c.Set("userRole", claims.Role)
		c.Set("newAccessToken", newAccessToken)
		c.Header("New-Access-Token", newAccessToken)
		c.Next(ctx)
		return
	}
}
