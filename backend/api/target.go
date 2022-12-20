package api

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
	"pingo/models"
)

func (a *Application) createTarget(c echo.Context) error {

	var t models.Target

	err := c.Bind(&t)
	if err != nil {
		log.Error(err)
		return echo.ErrBadRequest
	}

	err = a.db.Create(&t).Error
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusBadRequest, err)
		//return echo.ErrBadRequest
	}

	return c.JSON(http.StatusCreated, t)
}

func (a *Application) getTarget(c echo.Context) error {

	targetId := c.Param("id")

	var err error
	var t []models.Target
	if targetId == "" {
		err = a.db.Find(&t).Error
	} else {
		err = a.db.First(&t, targetId).Error
	}

	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, t)
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
