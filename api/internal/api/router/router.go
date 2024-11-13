package router

import (
	"github.com/circuit-shell/playlist-builder-back/internal/api/handler"
	"github.com/circuit-shell/playlist-builder-back/internal/repository/sqlite"
	"github.com/circuit-shell/playlist-builder-back/internal/service"
	"github.com/circuit-shell/playlist-builder-back/pkg/database"
	"github.com/gin-gonic/gin"
)

func SetupRouter() (*gin.Engine, error) {
	router := gin.Default()

	// Setup database
	db, err := database.NewSQLiteDB("storage/playlist.db")
	if err != nil {
		return nil, err
	}

	// Setup repositories
	songRepo := sqlite.NewSongRepository(db)

	// Setup services
	songService := service.NewSongService(songRepo)

	// Setup handlers
	greetingHandler := handler.NewGreetingHandler()
	songHandler := handler.NewSongHandler(songService)

	// Setup routes
	v1 := router.Group("/api/v1")
	{
		v1.GET("/greeting", greetingHandler.GetGreeting)
		v1.POST("/songs", songHandler.CreateSong)
		v1.GET("/songs", songHandler.GetAllSongs)
	}

	return router, nil
}
