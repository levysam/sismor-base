package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Health(g *gin.Engine) {
	g.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "Ok")
	})
}
