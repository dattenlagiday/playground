package main

import (
	bootstrap "apps/gin/internal/gin/boostrap"
	"apps/gin/internal/gin/utils"

	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

func main() {
	godotenv.Load()
	logger := utils.GetLogger().GetFxLogger()
	fx.New(bootstrap.Module, fx.Logger(logger)).Run()
}
