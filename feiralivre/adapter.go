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
		return c.JSON(http.StatusBadRequest, app.ResponseMessage("Bad request"))
	}
	err = a.service.Create(feiraLivre)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, app.ResponseMessage("Internal server error"))
	}
	return c.NoContent(http.StatusNoContent)
}

func (a FeiraLivreAdapter) SearchHTTPHandler(c echo.Context) error {
	q := c.QueryParam("q")
	result := a.service.Search(q)
	return c.JSON(http.StatusOK, app.ResponseData(result))
}

func (a FeiraLivreAdapter) GetHTTPHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	result := a.service.FindById(id)
	if result.ID != id || result.ID == 0 {
		return c.JSON(http.StatusNotFound, app.ResponseMessage("Not found"))	
	}
	return c.JSON(http.StatusOK, app.ResponseData(result))
}

func (a FeiraLivreAdapter) UpdateHTTPHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	feiraLivre := FeiraLivre{}
	err := c.Bind(&feiraLivre)
	if err != nil {
		return c.JSON(http.StatusBadRequest, app.ResponseMessage("Bad request"))
	}
	err = a.service.Update(id, feiraLivre)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, app.ResponseMessage("Internal server error"))
	}
	return c.NoContent(http.StatusAccepted)
}

func (a FeiraLivreAdapter) RemoveHTTPHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := a.service.Remove(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, app.ResponseMessage("Internal server error"))
	}
	return c.NoContent(http.StatusAccepted)
}