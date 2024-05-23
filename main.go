package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	// Graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		println("Shutting down...")
		cancel()
		os.Exit(1)
	}()

	// Start servers
	go standarLibraryApiRest()
	go muxApiRest()
	go echoApiRest()
	go ginApiRest()

	// Wait for shutdown
	<-ctx.Done()
}

func standarLibraryApiRest() {
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Response from standar library"}`))
	})
	http.ListenAndServe(":8080", nil)
}

func muxApiRest() {
	r := mux.NewRouter()
	r.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Response from mux"}`))
	}).Methods("GET")

	http.ListenAndServe(":8081", r)
}

func echoApiRest() {
	e := echo.New()
	e.GET("/test", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "Response from echo"})
	})
	e.Logger.Info(e.Start(":8082"))
}

func ginApiRest() {
	r := gin.Default()
	r.Handle("GET", "/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Response from gin",
		})
	})
	r.POST("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Response from gin",
		})
	})
	r.Run(":8083")
}
