package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/nicoabatedaga/golang_workshop/models"
	"github.com/nicoabatedaga/golang_workshop/services"
)

type UserHandler interface {
	GetUserHandler(c *gin.Context)
	PostUserHandler(c *gin.Context)
	DeleteUserHandler(c *gin.Context)
	PutUserHandler(c *gin.Context)
}

func NewUserHandler(s services.UserService) UserHandler {
	return &UserHandlerImp{s}
}

type UserHandlerImp struct {
	userService services.UserService
}

func (u *UserHandlerImp) GetUserHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{
			"error": "Bad Request, id is required",
		})
		return

	}
	user, err := u.userService.GetUser(id)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Internal Server Error",
		})
		return
	}
	c.JSON(200, user)
	return
}

func (u *UserHandlerImp) PostUserHandler(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Bad Request, invalid JSON",
		})
		return
	}
	postUser, err := u.userService.PostUser(user)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Internal Server Error",
		})
		return
	}
	c.JSON(200, postUser)
}

func (u *UserHandlerImp) DeleteUserHandler(c *gin.Context) {
	// return mock response
	c.JSON(200, gin.H{
		"message": "Delete User",
	})
	return
}

func (u *UserHandlerImp) PutUserHandler(c *gin.Context) {
	// return mock response
	c.JSON(200, gin.H{
		"message": "Put User",
	})
	return
}
