package api

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
	"pingo/models"
)

func (a *Application) createBootstrap(c echo.Context) error {

	target := &models.Target{Source: "https://google.com", Frequency: 1, Unit: "SECOND"}
	err := a.db.Create(target).Error
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, target)
}

func (a *Application) deleteBootstrap(c echo.Context) error {

	// Remove all targets
	err := a.db.Unscoped().Where("1 = 1").Delete(&models.Target{}).Error

	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, HttpEmptyJSONResponse)
}
