package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"pingo/models"
)

func (a *Application) createTarget(c echo.Context) error {

	t := &models.Target{
		Source:    "",
		Frequency: 0,
		Unit:      "",
	}

	// TODO: Replace with DB
	//targets = append(targets, *t)

	return c.JSON(http.StatusCreated, t)
}

func (a *Application) getTarget(c echo.Context) error {
	// TODO: Get from DB
	return c.JSON(http.StatusNotImplemented, nil)
}

// TODO: Implement
//func (a *Application) updateTarget(c echo.Context) error {
//	// TODO Update a DB thingo
//	return c.JSON(http.StatusNotImplemented, nil)
//}

func (a *Application) deleteTarget(c echo.Context) error {

	// TODO: Delete from DB
	//targets = nil

	return c.JSON(http.StatusNotImplemented, nil)
}
