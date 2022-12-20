package api

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
	"net/http"
	"pingo/models"
	"time"
)

func (a *Application) createBootstrap(c echo.Context) error {

	// Create targets
	target := []*models.Target{
		{Source: "https://google.com", Frequency: 1, Unit: "SECOND"},
		{Source: "https://github.com", Frequency: 5, Unit: "SECOND"},
		{Source: "https://example.com", Frequency: 3, Unit: "MINUTE"},
	}
	err := a.db.Create(&target).Error
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	// Create poll
	poll := models.Poll{
		Targets: []models.Target{
			{Model: gorm.Model{ID: 0}},
			{Model: gorm.Model{ID: 2}},
			{Model: gorm.Model{ID: 3}},
		},
		DurationSeconds: 60,
		StartedAt:       time.Time{},
	}
	err = a.db.Create(&poll).Error
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, HttpEmptyJSONResponse)
}

func (a *Application) deleteBootstrap(c echo.Context) error {

	// Remove all models
	var ml []interface{}
	ml = append(ml, &models.Poll{})
	ml = append(ml, &models.Target{})

	for _, m := range ml {
		err := a.db.Unscoped().Where("1 = 1").Delete(m).Error
		if err != nil {
			log.Error(err)
			return c.JSON(http.StatusInternalServerError, err)
		}
	}

	// Remove all relationships
	// TODO: Not working, investigate
	//goland:noinspection SqlWithoutWhere
	err := a.db.Raw("DELETE FROM poll_targets;").Error
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, err)

	}

	return c.JSON(http.StatusOK, HttpEmptyJSONResponse)
}
