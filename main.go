/*
Copyright © 2023 Zecheng Qian
*/
package main

import "github.com/Aden-Q/short-url/internal/pkg/router"

func main() {
	r := router.NewRouter()
	r.Run("localhost:8080")
}
