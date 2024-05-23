package server

import (
	"github.com/gin-gonic/gin"
	"github.com/nicoabatedaga/golang_workshop/handlers"
)

func createRoutes(userHandler handlers.UserHandler) {
	userRoutes := r.Group("user")
	userRoutes.GET("/:id", userHandler.GetUserHandler)
	userRoutes.POST("", userHandler.PostUserHandler)

	userRoutes.DELETE("/:id", authMiddleware, userHandler.DeleteUserHandler)

	userRoutesWithAuthValidation := userRoutes.Group("", authMiddleware)
	userRoutesWithAuthValidation.PUT("/:id", userHandler.PutUserHandler)

}

func authMiddleware(ctx *gin.Context) {
	authorizationHeader := ctx.GetHeader("Authorization")
	if authorizationHeader == "" {
		ctx.JSON(401, gin.H{
			"error": "Unauthorized",
		})
		ctx.Abort()
		return
	}
	ctx.Next()
}
