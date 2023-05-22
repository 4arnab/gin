package middleware

import "github.com/gin-gonic/gin"

func Authenticate(context *gin.Context) {

	// if token is missing we are sending error message to the client and we stop the request
	if !(context.Request.Header.Get("Token") == "auth") {
		context.AbortWithStatusJSON(500, gin.H{
			"message": "Token is required",
		})

		return
	}

	// else the request is valid and we are continuing the request with the other handlers
	context.Next()
}

func AddHeaders(context *gin.Context) {
	context.Request.Header.Add("key", "this is the value")
	context.Next()
}
