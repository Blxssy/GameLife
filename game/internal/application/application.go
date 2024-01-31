package application

import (
	"context"
	"fmt"
	"github.com/blxsyy/gameLife/game/internal/config"
	"github.com/blxsyy/gameLife/game/pkg/life"
	"time"
)

type Application struct {
	cfg config.Config
}

func New(config config.Config) *Application {
	return &Application{cfg: config}
}

func (a *Application) Run(ctx context.Context) error {
	currentWorld, _ := life.NewWorld(a.cfg.Height, a.cfg.Width)
	nextWorld, _ := life.NewWorld(a.cfg.Height, a.cfg.Width)
	currentWorld.RandInit(30)
	for {
		fmt.Println(currentWorld.String())
		life.NextState(currentWorld, nextWorld)
		currentWorld = nextWorld
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			time.Sleep(100 * time.Millisecond)
			break
		}
		fmt.Print("\033[H\033[2J")
	}
}
