// Package main URL shortener API.
//
// Open API for URL shortener service
//
// Terms Of Service:
//
//     Schemes: http
//     Host: localhost:8000
//     BasePath: /api
//     Version: 1.0.0
//     License: MIT https://opensource.org/licenses/MIT
//     Contact: Pavel Bobylev<p_bobylev@bk.ru> https://github.com/phpCoder88
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package main

import (
	"log"

	"github.com/phpCoder88/golang-test/internal/config"
	"github.com/phpCoder88/golang-test/internal/ioc"
	"github.com/phpCoder88/golang-test/internal/server"
	"github.com/phpCoder88/golang-test/internal/version"

	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}

	logger = logger.With(
		zap.String("Version", version.Version),
		zap.String("BuildDate", version.BuildDate),
		zap.String("BuildCommit", version.BuildCommit),
	)

	defer func() {
		err = logger.Sync()
		if err != nil {
			log.Println(err)
		}
	}()
	slogger := logger.Sugar()

	slogger.Info("Starting the application...")
	slogger.Info("Reading configuration and initializing resources...")
	conf, err := config.GetConfig()
	if err != nil {
		slogger.Error(err)
		return
	}

	slogger.Info("Configuring the application units...")
	container := ioc.NewContainer(nil)
	apiServer := server.NewServer(slogger, conf, container)
	err = apiServer.Run()
	if err != nil {
		slogger.Error("Occurred error during stopping the API server.", "err", err)
	}

	slogger.Info("The app is calling the last defers and will be stopped.")
}
