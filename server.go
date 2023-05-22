package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/4arnab/gin/middleware"
	"github.com/gin-gonic/gin"
)

type USER struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	router := gin.Default()

	// using middleware in the app level
	// router.Use(middleware.Authenticate)

	router.GET("/products", middleware.Authenticate, middleware.AddHeaders, Products) // adding a middleware function in to a single route also chaining middleware functions

	// auth := gin.BasicAuth(gin.Accounts{
	// 	"user": "pass",
	// })

	// route grouping
	users := router.Group("/users")
	{
		users.GET("productsTwo", Products)
		users.POST("/new", func(context *gin.Context) {
			// this context parameter contains all the request and response information or functions and meta data
			var user USER
			context.BindJSON(&user)

			context.JSON(200, user)
		})
	}

	// custom configuration
	server := &http.Server{
		Addr:         ":4000", // the port to listen
		Handler:      router,  // the routes
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("server listening on port 4000")
	server.ListenAndServe()
	// server.Run(":4000")

}
func Products(context *gin.Context) {
	fmt.Println(context.Request.Header.Get("key"), "THIS IS THE HEADERS 💡")
	page := context.Query("page")
	size := context.Query("size")

	// formData := context.PostForm()
	// jsonData, _ := ioutil.ReadAll(context.Request.Body)

	var user USER
	context.BindJSON(&user)

	fmt.Print(user)
	// context.JSON(200, map[string]string{"page": page, "size": size})
	context.JSON(200, gin.H{"page": page, "size": size, "bodyData": user})
}
