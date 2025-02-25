package handlers

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func StatusHandler(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
		"message": "this battle station is fully operational",
	})
}

func ExhaustPortHandler(c *gin.Context) {

    c.String(http.StatusOK, "the force is strong with this one")
}