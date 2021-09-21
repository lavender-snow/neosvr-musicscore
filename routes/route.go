package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routes(engine *gin.Engine) {
	engine.GET("/", indexHandler)
}

func indexHandler(context *gin.Context) {
	context.HTML(http.StatusOK, "index.tmpl", gin.H{})
}
