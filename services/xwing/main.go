package main

import (
    "github.com/gin-gonic/gin"
    handlers "monorepo/services/xwing/handlers"

)

func main() {
    router := gin.Default()
    router.GET("/status", handlers.StatusHandler)
    router.Run(":8080")
}
