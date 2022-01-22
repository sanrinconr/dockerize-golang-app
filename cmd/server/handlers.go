package server

import (
	"dockerize/cmd/middlewares"
	"github.com/gin-gonic/gin"
)

func setRoutes(router *gin.Engine) {
	resolver := resolver()
	public := router.Group("")
	public.GET("/ping", middlewares.Bridge(resolver.Ping()))
	public.GET("/pingDB", middlewares.Bridge(resolver.PingDB()))
	public.GET("/error", middlewares.Bridge(resolver.Error()))
}
