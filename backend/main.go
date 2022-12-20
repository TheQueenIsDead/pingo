package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"pingo/api"
	"pingo/models"
)

func main() {

	// Echo instance
	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)

	// DB
	db, err := gorm.Open(sqlite.Open("pingo.db"), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	err = db.AutoMigrate(
		&models.Target{},
		&models.Poll{})
	if err != nil {
		e.Logger.Fatal(err)
	}

	// Middleware
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))

	// Routes
	api.RegisterRoutes(api.NewApplication(db), e.Group("/api"))

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
