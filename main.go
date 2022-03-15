package main

import (
	"log"

	auth "tusharhow/auth/handlers"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func main() {
	router.POST("/login", auth.Login)
	router.POST("/register", auth.Register)
	log.Fatal(router.Run(":8080"))
}
