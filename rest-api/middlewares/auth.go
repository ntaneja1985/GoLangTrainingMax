package middlewares

import (
	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "No token found"})
		return
	}

	//Verify the Token
	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err})
		return
	}

	//Attach data to the context, in this case the UserId
	context.Set("userId", userId)

	//Call the next method in the request pipeline
	context.Next()
}
