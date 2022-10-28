package api

import (
	"context"
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

	g.POST(apiEndpoint, h.create)
	g.GET(apiEndpoint, h.all)
	//g.GET(apiEndpoint, h.all) // TODO: Get by ID
	g.DELETE(apiEndpoint, h.delete)

}

// DB Interface

type handler struct {
	db targetDB
}

type targetDB interface {
	CreateTarget(ctx context.Context, t *models.Target) (models.Target, error)
	ReadTarget(ctx context.Context, ID int) (string, error)
	UpdateTarget(ctx context.Context, ID int, t models.Target) (models.Target, error)
	DeleteTarget(ctx context.Context, ID int) (string, error)
}

// Handlers

func (h *handler) create(c echo.Context) error {

	t := new(models.Target)

	if err := c.Bind(t); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// TODO: Enable validation
	//if err := c.Validate(u); err != nil {
	//	return err
	//}

	result, err := h.db.CreateTarget(c.Request().Context(), t)
	if err != nil {
		return err // TODO: HTTP Error?
	}

	return c.JSON(http.StatusCreated, result)
}

func (h *handler) all(c echo.Context) error {

	var id int

	result, err := h.db.ReadTarget(c.Request().Context(), id)
	if err != nil {
		return err
	}

	return c.String(http.StatusNotImplemented, result)
}

func (h *handler) delete(c echo.Context) error {

	var id int

	result, err := h.db.DeleteTarget(c.Request().Context(), id)
	if err != nil {
		return err
	}

	return c.String(http.StatusNotImplemented, result)
}
