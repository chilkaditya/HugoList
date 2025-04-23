package handlers

import (
    "net/http"
    "youtube-playlist-automation/services"

    "github.com/gin-gonic/gin"
)

type CreatePlaylistRequest struct {
    Title       string   `json:"title"`
    Description string   `json:"description"`
    Songs       []string `json:"songs"`
}

func CreatePlaylistHandler(c *gin.Context) {
    var req CreatePlaylistRequest
    if err := c.BindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    err := services.CreatePlaylistAndAddSongs(req.Title, req.Description, req.Songs)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Playlist created successfully!"})
}
