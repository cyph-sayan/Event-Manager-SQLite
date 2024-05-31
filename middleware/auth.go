package middleware

import (
	"events-management/utility"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authorization(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Could not authorize user"})
		return
	}

	userId, err := utility.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err})
		return
	}
	context.Set("userId", userId)
	context.Next()
}
