package feiralivre

import (
	"net/http"

	"github.com/labstack/echo"
)

type FeiraLivreAdapter struct {
	service FeiraLivreService
}

func NewAdapter(service FeiraLivreService) FeiraLivreAdapter {
	return FeiraLivreAdapter{
		service: service,
	}
}

func (a FeiraLivreAdapter) CreateHTTPHandler(c echo.Context) error {
	feiraLivre := FeiraLivre{}
	err := c.Bind(&feiraLivre)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = a.service.Create(feiraLivre)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusNoContent, nil)
}

func (a FeiraLivreAdapter) SearchHTTPHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, a.service.List())
}

func (a FeiraLivreAdapter) GetHTTPHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, a.service.List())
}

func (a FeiraLivreAdapter) UpdateHTTPHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, a.service.List())
}

func (a FeiraLivreAdapter) RemoveHTTPHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, a.service.List())
}