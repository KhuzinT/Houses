package main

import (
	"Houses/internal/db/sqlite"
	"Houses/internal/handler/rest"
	"Houses/internal/repo"
	"Houses/internal/service"
	"Houses/internal/utils"
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

const dbFilename = "houses.db"
const secretKey = "https://youtu.be/dQw4w9WgXcQ"

func main() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", 15*time.Second,
		"The duration for which the server gracefully wait for existing connections to finish")
	flag.Parse()

	// Clear database
	err := os.Remove(dbFilename)
	if err != nil {
		log.Println("Failed to remove database")
	} else {
		log.Println("Database is removed")
	}

	// Init database
	db, err := sqlite.NewSQLite(dbFilename)
	if err != nil {
		log.Panicf("Failed to initialize database: %v\n", err)
	} else {
		log.Println("Database is initialized")
	}

	r := repo.NewRepo(db)
	s := service.NewService(r)
	a := utils.NewAuthManager([]byte(secretKey))
	h := rest.NewHandler(s, a)

	srv := &http.Server{
		Addr: "127.0.0.1:8000",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      h.NewRouter(), // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("Server Listen and Serve: %v\n", err)
		}
	}()

	c := make(chan os.Signal, 1)

	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("Server Shutdown: %v\n", err)
	}

	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("Shutting down server")
	os.Exit(0)
}
