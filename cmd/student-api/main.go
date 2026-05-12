package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Go-REST-API/database"
	"github.com/Go-REST-API/internal/http/handles/student"

	"github.com/Go-REST-API/config"
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
	router.HandleFunc("POST /api/students", student.New())
	router.HandleFunc("GET  /api/students/{id}", student.GetStudent())

	router.HandleFunc("GET /api/students", student.GetList())

	router.HandleFunc("PUT /api/students/{id}", student.UpdateStudent())

	router.HandleFunc("DELETE /api/students/{id}", student.DeleteStudent())

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
