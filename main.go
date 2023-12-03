/*
Copyright © 2023 Zecheng Qian
*/
package main

import (
	"context"
	"os"
	"time"

	"github.com/Aden-Q/short-url/internal/cache"
	"github.com/Aden-Q/short-url/internal/db"
	"github.com/Aden-Q/short-url/internal/logger"
	"github.com/Aden-Q/short-url/internal/redis"
	"github.com/Aden-Q/short-url/internal/router"
	"github.com/Aden-Q/short-url/internal/setting"
)

var (
	// configs is a global instance of the app/server settings
	configs *setting.Setting
	// log is a global logger instance
	log *logger.Logger
)

func initSetting() error {
	var err error
	configs, err = setting.Load()

	return err
}

func initLogger() {
	log = logger.New(os.Stdout)
}

func init() {
	initLogger()
	if err := initSetting(); err != nil {
		log.Fatal().Err(err).Msg("Failed to load configs")
	}
}

// @title short URL
// @version 2.0
// @description A fast URL shortener service written in Go.
// @termsOfService https://github.com/Aden-Q/short-url
func main() {
	log.Info().Msg("Starting server...")

	serverCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// attch the logger to the context
	serverCtx = log.WithContext(serverCtx)

	// connect to the mysql database server
	dbClient, err := db.NewEngine(db.Config{
		MySQLDSN: configs.MySQLDSN,
	})
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
		panic(err)
	}

	// connect to redis server
	redisClient, err := redis.NewClient(serverCtx, redis.Config{
		Addr: configs.RedisAddr,
	})
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to redis")
		panic(err)
	}

	// create a cache client
	redisCache := cache.New(cache.Config{
		Redis: redisClient,
	})

	r := router.New(router.Config{
		RequestTimeout: time.Duration(configs.RequestTimeout),
		DB:             dbClient,
		Redis:          redisClient,
		Cache:          redisCache,
	},
	)

	// Run is a blocking method, it only retuns when the server is shut down
	if err := r.Run(configs.ServerAddr); err != nil {
		log.Fatal().Err(err).Msg("Server stopped")
		panic(err)
	}

	log.Info().Msg("Server stopped")
}
