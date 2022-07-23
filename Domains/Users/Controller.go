package users

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func List(ctx *gin.Context) {
	users, err := GetUsers()
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func Detail(ctx *gin.Context) {
	type ids struct {
		Id int64 `uri:"id" binding:"required,uuid"`
	}
	var stringId ids
	ctx.ShouldBindUri(&stringId)
	user, err := GetUser(stringId.Id)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, user)
}
