package main

import (
	"log"

	"github.com/codescalersinternships/todoapp-omar/internal"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := internal.SetupDB(); err != nil {
		log.Fatal(err)
	}
	defer internal.Client.Close()

	router := gin.Default()
	internal.Routes(router)
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
