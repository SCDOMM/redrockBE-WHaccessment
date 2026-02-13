package router

import (
	"ProjectAndroidTest/handler"
	"ProjectAndroidTest/middleware"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func InitRouter(h *server.Hertz) {
	g1 := h.Group("/home", utils.MiddleHandler())
	g1.GET("/homepage", handler.HomePageHandler)
	g1.GET("/search/:keyWords", handler.SearchHandler)
	g1.POST("/upload", handler.HomeUploadHandler)

	const customBodyLimit = 10 * 1024 * 1024
	g2 := h.Group("/chat", utils.MiddleHandler())
	g2.GET("/mainpage", handler.ChatPageHandler)
	g2.GET("/search/:keyWords", handler.ChatSearchHandler)
	g2.POST("/upload", utils.LoggerMiddleware(), handler.ChatUploadHandler)

	g3 := h.Group("/reverso", utils.MiddleHandler())
	g3.POST("/login", handler.LogonHandler)
	g3.POST("/register", handler.RegisterHandler)
	g3.POST("/deregister", utils.LoggerMiddleware(), handler.DeregisterHandler)
	g3.POST("/change", utils.LoggerMiddleware(), handler.ChangeProfileHandler)
	g3.GET("/test", utils.LoggerMiddleware(), handler.JWTtestHandler)

	g4 := h.Group("/admin", utils.MiddleHandler(), utils.LoggerMiddleware())
	g4.POST("/deleteHome", handler.DeleteHomeHandler)
	g4.POST("/uploadHome", handler.HomeUploadHandler)
	g4.POST("/deleteChat", handler.DeleteChatHandler)
}
