/*
Copyright © 2023 Zecheng Qian
*/
package main

import (
	"log"

	"github.com/Aden-Q/short-url/internal/db"
	"github.com/Aden-Q/short-url/internal/router"
	"github.com/Aden-Q/short-url/internal/settings"
)

func main() {
	appConfigs, err := settings.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := db.NewDBEngine()
	if err != nil {
		panic(err)
	}

	r := router.NewRouter(
		router.RouterConfig{
			DB: db,
		},
	)
	r.Run(appConfigs.ServerAddr)
}
