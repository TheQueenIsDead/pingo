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

	var tl []models.Target
	for _, id := range p.TargetIds {

		tl = append(tl, models.Target{
			Model: gorm.Model{ID: uint(id)},
		})
	}
	p.Targets = tl

	// Create poll
	if err := a.db.Create(&p).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, p)
}

// TODO: Implement
func (a *Application) getPoll(c echo.Context) error {

	//id := c.Param("id")

	var err error
	var pl []models.Poll

	//err = a.db.Find(&pl).Association("Targets").Error
	err = a.db.Preload("Targets").Find(&pl).Error

	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusNotFound, err)
	}

	// Build Ids into JSON response
	for pidx, p := range pl {
		var tids []int
		for _, t := range p.Targets {
			tids = append(tids, int(t.ID))
		}
		p.TargetIds = tids
		pl[pidx] = p
	}

	return c.JSON(http.StatusOK, pl)

}
