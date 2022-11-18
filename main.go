package main

import (
	"geinterra/config"
	mid "geinterra/middleware"
	"geinterra/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	mid.LogMiddleware(e)
	e.Logger.Fatal(e.Start(":8000"))
}
