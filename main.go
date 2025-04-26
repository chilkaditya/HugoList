package main

import (
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
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
    r.LoadHTMLGlob("frontend/*")

    // Enable CORS
    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:5500"},
        AllowMethods:     []string{"GET", "POST", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type"},
        AllowCredentials: true,
    }))

    // Routes
    r.GET("/",handlers.HomePageHandler)
    r.GET("/auth", handlers.AuthHandler)
    r.GET("/auth/callback", handlers.AuthCallback)
    r.POST("/create-playlist", handlers.CreatePlaylistHandler)

    r.Run(":8080") // Run server on port 8080
}
