package main

import (
	"context"
	"github.com/blxsyy/gameLife/game/internal/application"
	"os"
)

func main() {
	ctx := context.Background()
	// Exit приводит к завершению программы с заданным кодом.
	os.Exit(mainWithExitCode(ctx))
}

func mainWithExitCode(ctx context.Context) int {
	cfg := application.Config{
		Width:  20,
		Height: 20,
	}
	app := application.New(cfg)

	return app.Run(ctx)
}
