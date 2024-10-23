package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/pos-backend/contracts"
	"github.com/umardev500/pos-backend/database"
	"github.com/umardev500/pos-backend/di"
	"github.com/umardev500/pos-backend/router"
)

type App struct{}

func NewApp() *App {
	return &App{}
}

func (a *App) Start(ctx context.Context, r contracts.RouterInterface) error {
	fiberApp := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})
	ch := make(chan error, 1)

	// Setup router
	r.Setup(fiberApp)

	go func() {
		port := os.Getenv("PORT")
		log.Info().Str("port", port).Msg("Starting server")
		ch <- fiberApp.Listen(":" + port)
	}()

	select {
	case err := <-ch:
		return err
	case <-ctx.Done():
		fmt.Println()
		log.Info().Msg("Shutting down server")
		fiberApp.Shutdown()
	}

	return nil
}

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("err", err)
	}

	// Setup zerolog
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	ctx, cancel := signal.NotifyContext(context.Background(), os.Kill, os.Interrupt)
	defer cancel()

	db := database.NewDB(ctx)

	// Initialize DI container
	container := di.NewRegistryContainer(db)

	// Initialize router
	r := router.NewRouter(container)

	// Start server
	err = NewApp().Start(ctx, r)
	if err != nil {
		fmt.Println("err", err)
	}
}
