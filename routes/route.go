package routes

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lavender-snow/neosvr-musicscore/app/controllers"
)

func Routes(engine *gin.Engine) {
	engine.GET("/score-creator", indexHandler)
	engine.POST("/score-creator/create", createHandler)
	engine.GET("/score-creator/score-data", scoreDataHandler)
}

func indexHandler(context *gin.Context) {
	context.HTML(http.StatusOK, "index.tmpl", gin.H{})
}

func createHandler(context *gin.Context) {
	log.Printf(context.Request.URL.Path)
	scores := context.PostFormArray("score")
	url := controllers.ScoreCreateController(scores)
	context.String(http.StatusOK, url)
}

func scoreDataHandler(context *gin.Context) {
	context.String(http.StatusOK, context.Param("v"))
}
