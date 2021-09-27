package routes

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lavender-snow/neosvr-musicscore/app/controllers"
)

// Routes ルート情報を設定
func Routes(engine *gin.Engine) {
	engine.GET("/score-creator", indexHandler)
	engine.POST("/score-creator/create", createHandler)
	engine.GET("/score-creator/score", scoreDataHandler)
}

func indexHandler(context *gin.Context) {
	context.HTML(http.StatusOK, "index.tmpl", gin.H{})
}

func createHandler(context *gin.Context) {
	log.Printf(context.Request.URL.Path)
	scores := context.PostFormArray("score")
	hash, err := controllers.CreateScoreController(scores)
	if err != nil {
		context.String(http.StatusServiceUnavailable, "譜面生成に失敗しました.\r\nFailed to generate score.")
	}
	context.String(http.StatusOK, hash)
}

func scoreDataHandler(context *gin.Context) {

	score, err := controllers.GetScoreController(context.Query("v"))

	if err != nil {
		log.Printf(err.Error())
		context.String(http.StatusBadRequest, "BadRequest")
	}

	context.String(http.StatusOK, score)
}
