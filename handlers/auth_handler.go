package handlers

import (
    "fmt"
    "net/http"
    "youtube-playlist-automation/utils"

    "github.com/gin-gonic/gin"
)

func AuthHandler(c *gin.Context) {
    url := utils.GetAuthURL()
    fmt.Printf("Redirect url: %s\n",url)
    c.Redirect(http.StatusTemporaryRedirect, url)
}

func AuthCallback(c *gin.Context) {
    code := c.Query("code")
    token, err := utils.ExchangeCode(code)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    fmt.Printf("Token: %s\n",token)
    utils.SaveToken(token) // save token for later use (in memory for now)

    c.Redirect(http.StatusFound, "http://localhost:5500/frontend/form.html")
}
