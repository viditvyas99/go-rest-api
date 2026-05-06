package main

import (
	"context"
	"example/rest-api/config"
	"example/rest-api/database"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Println("welcome")

	//load config
	config.Load()

	//connect to database
	dbUser := config.AppConfig.Database.Username
	database.Connect()
	fmt.Println("Database User:", dbUser)

	//setup routes

	router := http.NewServeMux()
	router.HandleFunc("/students", func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("welcome to student api "))

	})
	// setup server
	// http.ListenAndServe(config.AppConfig.Address, router)

	server := http.Server{
		Addr:    config.AppConfig.Address,
		Handler: router,
	}

	slog.Info("API server started", "addr", config.AppConfig.Address)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("ListenAndServe error", "error", err)
		}
	}()

	<-done // wait for signal
	slog.Info("Shutdown signal received")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Server shutdown failed", "error", err)
	} else {
		slog.Info("Server stopped gracefully")
	}

}
