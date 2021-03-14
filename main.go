package main

import (
	"net/http"
	"unico/app"
	"unico/app/repository"
	"unico/app/helper/io"
	"unico/feiralivre"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	database := "./unico.db"
	io.Removefile(database)
	db := repository.Conn(database)
	defer db.Close()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))
	e.GET("/", healthHTTPHandler)

	// feira livres
	repo := feiralivre.NewRepository(db)
	service := feiralivre.NewService(repo)
	adapter := feiralivre.NewAdapter(service)

	f := e.Group("feiraslivres")
	f.GET("", adapter.SearchHTTPHandler)
	f.POST("", adapter.CreateHTTPHandler)
	f.GET("/:id", adapter.GetHTTPHandler)
	f.PATCH("/:id", adapter.UpdateHTTPHandler)
	f.DELETE("/:id", adapter.RemoveHTTPHandler)

	e.Logger.Fatal(e.Start(":8000"))
}

//HealthHTTPHandler greatings
func healthHTTPHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, app.ResponseMessage("I'm good, tks!"))
}