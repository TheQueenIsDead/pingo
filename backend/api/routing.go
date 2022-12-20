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
	t.GET("", a.getTarget)
	t.GET("/:id", a.getTarget)
	t.POST("", a.createTarget)
	t.DELETE("", a.deleteTarget)
	t.DELETE("/:id", a.deleteTarget)

	// Poll
	p := parent.Group("/poll")
	p.GET("/", a.getPoll)
	p.GET("/:id", a.getPoll)
	p.POST("", a.createPoll)
}
