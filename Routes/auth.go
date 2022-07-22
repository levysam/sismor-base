package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func Auth(r *gin.Engine) {
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		c.JSON(
			http.StatusOK,
			gin.H{"nome": user},
		)
	})
}
