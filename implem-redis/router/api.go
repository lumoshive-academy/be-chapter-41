package router

import (
	"context"
	"golang-chapter-41/implem-redis/infra"
	"golang-chapter-41/implem-redis/middleware"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func SetupReouter(ctx infra.Context, middleware middleware.Middleware) {
	router := gin.Default()

	v1 := router.Group("/v1")
	v1.Use(middleware.Logger())

	// shipping routes
	customer := v1.Group("/shipping")
	{
		customer.Use(middleware.Authentication())
		customer.GET("/", ctx.Handler.ShippingHandler.GetAllShipping)
		customer.GET("/cost", ctx.Handler.ShippingHandler.ShippingCost)
	}

	srv := &http.Server{
		Addr:    ":" + ctx.Config.Port,
		Handler: router.Handler(),
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctxt); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")
}
