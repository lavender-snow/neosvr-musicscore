package main

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/lavender-snow/neosvr-musicscore/config"
	"github.com/lavender-snow/neosvr-musicscore/routes"
	"github.com/lavender-snow/neosvr-musicscore/utils"
)

func main() {
	// ログ設定
	utils.LoggingSettings("musicScore")

	// gin初期化
	engine := gin.Default()
	store := sessions.NewCookieStore([]byte(config.Config.SecretKey))
	store.Options(sessions.Options{MaxAge: 3600})
	engine.Use(sessions.Sessions("musicscore", store))
	engine.Static("/assets", "./assets")
	engine.LoadHTMLGlob("app/views/templates/*.tmpl")
	routes.Routes(engine)

	// 開始
	if config.Config.RunMode == "https" {
		engine.RunTLS(config.Config.PortNo, config.Config.CertFile, config.Config.KeyFile)
	} else {
		engine.Run(config.Config.PortNo)
	}

}
