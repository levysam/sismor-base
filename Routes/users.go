package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	users "github.com/levysam/sismor-base/Domains/Users"
)

func Users(r *gin.Engine) {
	r.GET("/usuarios", func(ctx *gin.Context) {
		users := users.GetUsers()
		ctx.JSON(http.StatusOK, users)
	})
}
