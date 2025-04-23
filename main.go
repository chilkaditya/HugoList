package main

import (
    "github.com/gin-gonic/gin"
    "youtube-playlist-automation/handlers"
	"youtube-playlist-automation/utils"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
	utils.InitOAuthConfig()

    r := gin.Default()

    // Routes
    r.GET("/auth", handlers.AuthHandler)
    r.GET("/auth/callback", handlers.AuthCallback)
    r.POST("/create-playlist", handlers.CreatePlaylistHandler)

    r.Run(":8080") // Run server on port 8080
}
