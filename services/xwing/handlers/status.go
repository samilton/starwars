package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func StatusHandler(c *gin.Context) {
	c.String(http.StatusOK, "torpedoes ready")
}

func LaunchTorpedoHander(c *gin.Context) {
	c.String(http.StatusOK, "Launching Torpedo")
}
