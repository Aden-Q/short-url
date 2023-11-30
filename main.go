/*
Copyright © 2023 Zecheng Qian
*/
package main

import (
	"log"

	"github.com/Aden-Q/short-url/internal/pkg/db"
	"github.com/Aden-Q/short-url/internal/pkg/model"
	"github.com/Aden-Q/short-url/internal/pkg/router"
	"github.com/Aden-Q/short-url/internal/pkg/settings"
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

	db.AutoMigrate(&model.URL{})

	r := router.NewRouter()
	r.Run(appConfigs.ServerAddr)
}
