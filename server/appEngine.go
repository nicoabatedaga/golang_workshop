package server

import (
	"github.com/gin-gonic/gin"
	"github.com/nicoabatedaga/golang_workshop/handlers"
	"github.com/nicoabatedaga/golang_workshop/services"
	"github.com/nicoabatedaga/golang_workshop/storage"
)

var r *gin.Engine

func init() {
	r = gin.Default()
	r.Use(
		gin.Recovery(),
	)
	// memcache := storage.NewStorageMemcached()
	redis := storage.NewStorageRedis()
	userService := services.NewUserService(redis)
	userHandler := handlers.NewUserHandler(userService)
	createRoutes(userHandler)
}

func Start() {
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
