package main

import (
	"context"
	"os"

	"github.com/RuGoffer/gofr/internal/application"
)

func main() {
	ctx := context.Background()
	// Exit приводит к завершению программы с заданным кодом.
	os.Exit(mainWithExitCode(ctx))
}

func mainWithExitCode(ctx context.Context) int {
	cfg := application.Config{
		Width:  10,
		Height: 10,
	}
	app := application.New(cfg)

	return app.Run(ctx)
}
