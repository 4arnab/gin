package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "this is my server",
		})
	})

	server.GET("/products/", func(context *gin.Context) {
		page := context.Query("page")
		size := context.Query("size")

		// formData := context.PostForm()

		fmt.Print()
		context.JSON(200, map[string]string{"page": page, "size": size})
	})

	fmt.Println("server listening on port 4000")
	server.Run(":4000")
}
