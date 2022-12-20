package api

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
	"net/http"
	"pingo/models"
)

// TODO: Implement
func (a *Application) createPoll(c echo.Context) error {

	var p models.Poll

	err := c.Bind(&p)
	if err != nil {
		log.Error(err)
		return echo.ErrBadRequest
	}

	// Inflate the target ids into structs
	for id := range p.TargetIds {
		p.Targets = append(p.Targets, models.Target{
			Model: gorm.Model{
				ID: uint(id),
			},
		})
	}

	// Create poll
	if err := a.db.Create(&p).Error; err != nil {
		return err
	}

	//a.db.Begin()
	//err = a.db.Create(&p).Error
	//if err != nil {
	//	log.Error(err)
	//	return c.JSON(http.StatusBadRequest, err)
	//}

	return c.JSON(http.StatusCreated, p)

}

// TODO: Implement
func (a *Application) getPoll(c echo.Context) error {

	id := c.Param("id")

	var err error
	var p []models.Poll
	if id == "" {
		err = a.db.Find(&p).Error
	} else {
		err = a.db.First(&p, id).Error
	}

	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, p)

}
