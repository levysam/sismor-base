package routes

import (
	"github.com/gin-gonic/gin"
	users "github.com/levysam/sismor-base/Domains/Users"
)

func Users(r *gin.Engine) {
	r.GET("/usuarios", users.List)
	r.GET("/user/:id", users.Detail)
}
