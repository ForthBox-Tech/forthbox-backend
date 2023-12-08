package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	router "forthboxbe/app/http"
	"forthboxbe/pkg/setting"

	"github.com/gin-gonic/gin"
)

func init() {
	os.Setenv("TZ", "Asia/Shanghai") 
}

// main configures the HTTP server and handles graceful shutdown.
func main() {
	gin.SetMode(setting.ServerSetting.RunMode)
	endPoint := fmt.Sprintf("0.0.0.0:%d", setting.ServerSetting.HttpPort)

	s := &http.Server{
		Addr:         endPoint,
		Handler:      router.GetRouter(),
		ReadTimeout:  setting.ServerSetting.ReadTimeout,
		WriteTimeout: setting.ServerSetting.WriteTimeout,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := s.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}



