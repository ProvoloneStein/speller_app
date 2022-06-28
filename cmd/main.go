package main

import (
	"Nexign/internal"
	"Nexign/internal/handler"
	"Nexign/internal/service"
	"Nexign/pkg/logger"
	"context"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx := context.Background()
	logger.GetLogger(ctx).Info("Starting Speller server")

	if err := initConfig(); err != nil {
		logger.GetLogger(ctx).Fatal("Initializing configs error: ", zap.Error(err))
	}

	services := service.NewService(&service.SpellerConfig{Url: viper.GetString("yandex_speller.url"),
		Lang: viper.GetString("yandex_speller.lang"), Format: viper.GetString("yandex_speller.format")})
	handlers := handler.NewHandler(services)
	srv := new(internal.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logger.GetLogger(ctx).Fatal("Http server running error: %s", zap.Error(err))
		}
	}()

	logger.GetLogger(ctx).Info("Speller Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logger.GetLogger(ctx).Info("Speller Shutting Down")

	if err := srv.Shutdown(ctx); err != nil {
		logger.GetLogger(ctx).Error("Shutting down error: %s", zap.Error(err))
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
