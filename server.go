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

	type USER struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	newUser := USER{
		Name: "Ahmed Arnab",
		Age:  19,
	}
	server.POST("/products/", func(context *gin.Context) {
		page := context.Query("page")
		size := context.Query("size")

		// formData := context.PostForm()
		// jsonData, _ := ioutil.ReadAll(context.Request.Body)

		var user USER
		context.BindJSON(&user)

		fmt.Print(user)
		// context.JSON(200, map[string]string{"page": page, "size": size})
		context.JSON(200, gin.H{"user": newUser, "page": page, "size": size, "bodyData": user})
	})

	fmt.Println("server listening on port 4000")
	server.Run(":4000")
}
