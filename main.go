package main

import (
	"MarvelousBlog-Backend/config"
	r "MarvelousBlog-Backend/router"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(config.AppMode)
	router := gin.Default()
	r.LoadVisitorRouters(router)
	err := router.Run(config.ServerPort)
	if err != nil {
		fmt.Println("Service run failed!")
	}
}
