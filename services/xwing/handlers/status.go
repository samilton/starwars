package handlers

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func StatusHandler(c *gin.Context) {
    c.String(http.StatusOK, "torpedoes ready")
}