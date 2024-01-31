package main

import (
	"context"
	"errors"
	"github.com/blxsyy/gameLife/game/internal/application"
	"github.com/blxsyy/gameLife/game/internal/config"
	"log"
	"os"
)

func main() {
	ctx := context.Background()
	os.Exit(mainWithExitCode(ctx))
}

func mainWithExitCode(ctx context.Context) int {
	cfg := config.Config{
		Width:  10,
		Height: 10,
	}
	app := application.New(cfg)

	if err := app.Run(ctx); err != nil {
		switch {
		case errors.Is(err, context.Canceled):
			log.Println("Processing cancelled")
		default:
			log.Println("Application run error", err)
		}
		return 1
	}
	return 0
}
