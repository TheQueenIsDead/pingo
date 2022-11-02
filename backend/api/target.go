package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"pingo/db"
	"pingo/models"
)

const apiEndpoint = "/target"

// Routes

func RegisterRoutes(parent *echo.Group, db *db.PingoDB) {
	h := handler{
		db: db,
	}

	g := parent.Group("/target")

	g.POST("", h.create)
	g.GET("", h.all)
	//g.GET("/:id", h.get) // TODO: Get by ID
	g.DELETE("", h.delete)

}

// DB Interface

type handler struct {
	db targetDB
}

type targetDB interface {
	//CreateTarget(ctx context.Context) (models.Target, error)
	//ReadTarget(ctx context.Context) (string, error)
	//UpdateTarget(ctx context.Context) (models.Target, error)
	// DeleteTarget(ctx context.Context) (string, error)
}

// Handlers

var targets []models.Target

func (h *handler) create(c echo.Context) error {

	t := &models.Target{
		Id:        0,
		Type:      "",
		Source:    "",
		Frequency: 0,
		Unit:      "",
	}

	targets = append(targets, *t)

	return c.JSON(http.StatusCreated, t)
}

func (h *handler) all(c echo.Context) error {
	return c.JSON(http.StatusOK, targets)
}

func (h *handler) delete(c echo.Context) error {

	targets = nil

	return c.JSON(http.StatusOK, targets)
}
