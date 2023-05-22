package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
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
	router.POST("/products/", func(context *gin.Context) {
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
	// server.Run(":4000")

	// custom configuration
	server := &http.Server{
		Addr:         ":4000", // the port to listen
		Handler:      router,  // the routes
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server.ListenAndServe()
}
