package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"pingo/api"
	"pingo/db"
)

func main() {
	// Echo instance
	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Database
	db := new(db.PingoDB)

	// Routes
	apiGroup := e.Group("/api")
	api.RegisterRoutes(apiGroup, db)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
