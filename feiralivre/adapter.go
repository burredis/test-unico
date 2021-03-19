package feiralivre

import (
	"net/http"
	"strconv"
	"unico/app"

	"github.com/labstack/echo/v4"
)

type FeiraLivreAdapter struct {
	service FeiraLivreService
}

func NewAdapter(service FeiraLivreService) FeiraLivreAdapter {
	return FeiraLivreAdapter{
		service: service,
	}
}

// CreateHTTPHandler godoc
// @Description Create a new item
// @Tags API
// @Accept json
// @Param data body FeiraLivre true " "
// @Success 204
// @Failure 400 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /feiraslivres [post]
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

// SearchHTTPHandler godoc
// @Description Search item by query
// @Tags API
// @Accept json
// @Param q query string false "Search by distrito or regiao5 or nome or bairro"
// @Response 200 {object} app.Response
// @Router /feiraslivres [get]
func (a FeiraLivreAdapter) SearchHTTPHandler(c echo.Context) error {
	q := c.QueryParam("q")
	result := a.service.Search(q)
	return c.JSON(http.StatusOK, app.ResponseData(result))
}

// GetHTTPHandler godoc
// @Description Get item by id
// @Tags API
// @Accept json
// @Success 200 {object} app.Response
// @Failure 404 {object} app.Response
// @Router /feiraslivres/:id [get]
func (a FeiraLivreAdapter) GetHTTPHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	result := a.service.FindById(id)
	if result.ID != id || result.ID == 0 {
		return c.JSON(http.StatusNotFound, app.ResponseMessage("Not found"))
	}
	return c.JSON(http.StatusOK, app.ResponseData(result))
}

// UpdateHTTPHandler godoc
// @Description Update item by id
// @Tags API
// @Accept json
// @Param data body FeiraLivre true " "
// @Success 204
// @Failure 400 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /feiraslivres/:id [patch]
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

// RemoveHTTPHandler godoc
// @Description Remove item by id
// @Tags API
// @Accept json
// @Success 202
// @Failure 500 {object} app.Response
// @Router /feiraslivres/:id [delete]
func (a FeiraLivreAdapter) RemoveHTTPHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := a.service.Remove(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, app.ResponseMessage("Internal server error"))
	}
	return c.NoContent(http.StatusAccepted)
}
