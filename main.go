package main

import (
	"jasaPengiriman/connection"
	"jasaPengiriman/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := connection.ConnectDB()
	rh := &router.Handlers{
		DB: db,
		R:  r,
	}
	rh.Routes()

	r.Run()
}
