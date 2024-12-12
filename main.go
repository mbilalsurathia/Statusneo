package main

import (
	"context"
	"fmt"
	"log"
	"maker-checker/rest"
	"maker-checker/service"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	serviceContainer := service.NewServiceContainer()
	server := rest.StartServer(serviceContainer)

	fmt.Println("========== Rest Server Started ============")

	// Create a channel to listen for termination signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Block until a signal is received
	<-quit
	fmt.Println("Shutting down server...")

	// Create a context with a timeout to allow for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Attempt to gracefully shut down the server
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	fmt.Println("Server exiting")

}
