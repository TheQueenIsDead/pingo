package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"pingo/api"
)

func main() {
	// Echo instance
	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)

	// Middleware
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))

	//// Database
	//db := new(db.PingoDB)
	//db.Init()

	// Routes
	apiGroup := e.Group("/api")
	//api.RegisterRoutes(apiGroup, db)
	api.RegisterRoutes(apiGroup, nil)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
