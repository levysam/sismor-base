package main

import (
	"github.com/gin-gonic/gin"
	routes "github.com/levysam/sismor-base/Routes"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	routes.Health(r)
	routes.Auth(r)
	routes.Users(r)

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
