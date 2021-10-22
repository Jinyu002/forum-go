package main

import (
	"forum1/controller"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.POST("/api/auth/ban", controller.Ban)
	r.POST("/api/auth/unban", controller.Unban)
	r.POST("/api/auth/number", controller.Number)
	r.POST("/api/auth/listUsers", controller.ListUsers)
	r.POST("/api/auth/publishPost", controller.PublishPost)

	return r
}