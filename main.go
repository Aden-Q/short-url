/*
Copyright © 2023 Zecheng Qian
*/
package main

import (
	"context"
	"log"

	"github.com/Aden-Q/short-url/internal/db"
	"github.com/Aden-Q/short-url/internal/redis"
	"github.com/Aden-Q/short-url/internal/router"
	"github.com/Aden-Q/short-url/internal/settings"
)

func main() {
	configs, err := settings.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	serverCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// connect to the mysql database
	db, err := db.NewEngine(db.Config{
		MySQLDSN: configs.MySQLDSN,
	})
	if err != nil {
		// TODO: add a log
		panic(err)
	}

	// connect to redis server
	redis, err := redis.NewClient(serverCtx, redis.Config{
		Addr: configs.RedisAddr,
	})
	if err != nil {
		// TODO: add a log
		panic(err)
	}

	r := router.NewRouter(
		router.Config{
			DB:    db,
			Redis: redis,
		},
	)

	// Run is a blocking method, it only retuns when the server is shut down
	if err := r.Run(configs.ServerAddr); err != nil {
		panic(err)
	}
}
