package main

import (
	"context"

	"github.com/erminson/tasker/internal/app"
	"github.com/erminson/tasker/pkg/logger"
)

func main() {
	ctx := context.Background()
	log := logger.New()

	a, err := app.New(ctx, log)
	if err != nil {
		log.Error("failed to init app", logger.Err(err))
		return
	}

	err = a.Run(ctx)
	if err != nil {
		log.Error("failed to run app", logger.Err(err))
		return
	}
}
