package api

import (
	"gjob-admin/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func hello(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": service.Hello(),
	})

}
