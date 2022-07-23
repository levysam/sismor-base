package routes

import (
	"github.com/gin-gonic/gin"
	auth "github.com/levysam/sismor-base/Auth"
	"github.com/levysam/sismor-base/middleware"
)

func Auth(router *gin.Engine) {
	router.POST("/login", auth.Login)
	router.POST("/logout", middleware.TokenAuthMiddleware(), auth.Logout)
}
