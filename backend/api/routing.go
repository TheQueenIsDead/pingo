package api

import (
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(a *Application, parent *echo.Group) {

	// Bootstrap
	b := parent.Group("/bootstrap")
	b.POST("", a.createBootstrap)
	b.DELETE("", a.deleteBootstrap)

	// Target
	t := parent.Group("/target")
	t.POST("", a.createTarget)
	t.GET("", a.getTarget)
	t.GET("/:id", a.getTarget)
	t.DELETE("", a.deleteTarget)
}
