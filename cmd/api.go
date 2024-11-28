package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	log "log/slog"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/rs/cors"

	"github.com/dariusigna/binarysearch/config"
	"github.com/dariusigna/binarysearch/internal/app"
	"github.com/dariusigna/binarysearch/internal/index"
)

func main() {
	if err := run(); err != nil {
		log.Error("Applications exists with", "error", err)
		os.Exit(1)
	}
}

func run() error {
	log.Info("Server is starting...")
	// Setup cancellation context
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		return err
	}

	// Set log level
	var slogLevel log.Level
	err = slogLevel.UnmarshalText([]byte(cfg.LogLevel))
	if err != nil {
		return fmt.Errorf("invalid log level: %v", err)
	}
	log.SetLogLoggerLevel(slogLevel)

	// Read input file
	numbers, err := readInputFile(cfg.InputFile)
	if err != nil {
		return err
	}

	// Create server
	indexFinder := index.NewFinder(numbers, cfg.ConformationLevel)
	srv := app.NewServer(indexFinder)

	// Configure CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.Port),
		Handler:      c.Handler(srv),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	// Start the server
	go func() {
		log.Info("Server is listening on", "port", cfg.Port)
		if err = server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error("Could not listen on %s: %v\n", cfg.Port, err)
			os.Exit(1)
		}
	}()

	// Graceful shutdown
	<-ctx.Done()
	stop()

	log.Info("Server is shutting down...")
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	if err = server.Shutdown(ctx); err != nil {
		log.Error("Could not gracefully shutdown the server", "error", err)
		return err
	}

	log.Info("Server stopped")
	return nil
}

func readInputFile(fileName string) ([]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	var numbers []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		value, err := strconv.Atoi(line)
		if err != nil {
			return nil, fmt.Errorf("failed to convert line to number: %v", err)
		}
		numbers = append(numbers, value)
	}

	if err = scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}

	return numbers, nil
}
