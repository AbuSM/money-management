package routes

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"tajguideapi/middleware"
	"tajguideapi/utils"
	"time"

	"github.com/gin-gonic/gin"
)

// Start routes working
func Start() {
	r := gin.Default()
	r.Use(middleware.Cors())
	r.GET("/ping", Ping)
	r.GET("/categories", getAllCategories)
	r.GET("/companies", getAllCompanies)
	r.GET("/company/:id", getCompanyByID)
	r.GET("/itemstat/:id", setStat)
	r.GET("/stats", getStatistic)
	
	r.GET("/companycategoryies/:id", getCompanyByCategory)
	r.GET("/news", getAllNews)
	r.GET("/newsbyid/:id", getNewsByID)
	r.GET("/rates", getTopRatedPlace)

	srv := &http.Server{
		Addr:           fmt.Sprintf(":%d", 8000),
		WriteTimeout:   60 * time.Second,
		ReadTimeout:    40 * time.Second,
		MaxHeaderBytes: 1 << 20,
		Handler:        r,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 20 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")
}
