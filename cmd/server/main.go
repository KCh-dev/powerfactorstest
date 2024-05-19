package main

import (
	"goapp/internal/config"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"

	goapp "goapp/internal/app/server"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags | log.Lmsgprefix | log.Lshortfile)
}

func main() {
	// Debug.
	go func() {
		log.Println(http.ListenAndServe(":6060", nil))
	}()

	// Load configuration
	cfg, err := config.LoadConfig("config/domains.json") // Ensure the path is correct
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Register signal handlers for exiting
	exitChannel := make(chan os.Signal, 1)
	signal.Notify(exitChannel, syscall.SIGINT, syscall.SIGTERM)

	// Start server with configuration
	if err := goapp.Start(cfg, exitChannel); err != nil {
		log.Fatalf("fatal: %+v\n", err)
	}
}
