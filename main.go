package main

import (
	"MarvelousBlog-Backend/config"
	r "MarvelousBlog-Backend/router"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(config.AppMode)
	router := r.SetupRouter()
	router.Run(config.ServerPort)
}
