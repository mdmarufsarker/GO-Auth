package main

import (
	"log"

	auth "tusharhow/auth/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/login", auth.Login)
	r.POST("/register", auth.Register)
	log.Fatal(r.Run(":8080"))
}