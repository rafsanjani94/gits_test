package main

import (
	"gits/test3/config"
	"gits/test3/handler"
	"gits/test3/service/login"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config.ConnectDB()
	config.ConnectRedis()

	r.POST("/login", handler.Login)
	r.POST("/book", login.Auth, handler.CreateBook)
	r.GET("/book", login.Auth, handler.ListBook)
	r.PUT("/book/:id", login.Auth, handler.UpdateBook)
	r.DELETE("/book/:id", login.Auth, handler.DeleteBook)

	r.GET("/pdf/books", handler.PdfBook)

	r.Run(":8080")
}
