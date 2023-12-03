/*
Copyright © 2023 Zecheng Qian
*/
package main

import (
	"context"
	"log"
	"os"

	"github.com/Aden-Q/short-url/internal/cache"
	"github.com/Aden-Q/short-url/internal/db"
	"github.com/Aden-Q/short-url/internal/logger"
	"github.com/Aden-Q/short-url/internal/redis"
	"github.com/Aden-Q/short-url/internal/router"
	"github.com/Aden-Q/short-url/internal/setting"
)

func main() {
	// load global settings
	configs, err := setting.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// initialize a logger instance
	logger := logger.New(os.Stdout)

	logger.Info().Msg("Starting server...")

	serverCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// attch the logger to the context
	serverCtx = logger.WithContext(serverCtx)

	// connect to the mysql database server
	dbClient, err := db.NewEngine(db.Config{
		MySQLDSN: configs.MySQLDSN,
	})
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to connect to database")
		panic(err)
	}

	// connect to redis server
	redisClient, err := redis.NewClient(serverCtx, redis.Config{
		Addr: configs.RedisAddr,
	})
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to connect to redis")
		panic(err)
	}

	// create a cache client
	redisCache := cache.New(cache.Config{
		Redis: redisClient,
	})

	r := router.New(router.Config{
		DB:    dbClient,
		Redis: redisClient,
		Cache: redisCache,
	},
	)

	// Run is a blocking method, it only retuns when the server is shut down
	if err := r.Run(configs.ServerAddr); err != nil {
		logger.Fatal().Err(err).Msg("Server stopped")
		panic(err)
	}

	logger.Info().Msg("Server stopped")
}
