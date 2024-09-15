package router

import (
	"musicservice/src/api/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	r.GET("/", handlers.GetMusicList)
	r.GET("/api/v2/musics", handlers.GetMusicList)

	return r
}
