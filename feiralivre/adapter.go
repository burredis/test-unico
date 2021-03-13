package feiralivre

import (
	"net/http"
	"strconv"
	"unico/app"
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
		errCode := http.StatusBadRequest
		return c.JSON(errCode, app.ResponseError("Bad request", errCode))
	}
	err = a.service.Create(feiraLivre)
	if err != nil {
		errCode := http.StatusInternalServerError
		return c.JSON(errCode, app.ResponseError("Internal server error", errCode))
	}
	return c.NoContent(http.StatusNoContent)
}

func (a FeiraLivreAdapter) SearchHTTPHandler(c echo.Context) error {
	result := a.service.Search()

	return c.JSON(http.StatusOK, app.ResponseSuccess(result))
}

func (a FeiraLivreAdapter) GetHTTPHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	result := a.service.FindById(id)

	if result.ID != id || id == 0 {
		errCode := http.StatusNotFound
		return c.JSON(errCode, app.ResponseError("Not found", errCode))	
	}

	return c.JSON(http.StatusOK, app.ResponseSuccess(a.service.FindById(id)))
}

func (a FeiraLivreAdapter) UpdateHTTPHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	feiraLivre := FeiraLivre{}
	err := c.Bind(&feiraLivre)
	if err != nil {
		errCode := http.StatusBadRequest
		return c.JSON(errCode, app.ResponseError("Bad request", errCode))
	}
	err = a.service.Update(id, feiraLivre)
	if err != nil {
		errCode := http.StatusInternalServerError
		return c.JSON(errCode, app.ResponseError("Internal server error", errCode))
	}
	return c.NoContent(http.StatusAccepted)
}

func (a FeiraLivreAdapter) RemoveHTTPHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := a.service.Remove(id)
	if err != nil {
		errCode := http.StatusInternalServerError
		return c.JSON(errCode, app.ResponseError("Internal server error", errCode))
	}
	return c.NoContent(http.StatusAccepted)
}