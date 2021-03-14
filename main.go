package main

import (
	"net/http"
	"unico/app"
	"unico/app/repository"
	"unico/app/helper/io"
	"unico/feiralivre"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.com/swaggo/echo-swagger/example/docs"
)

// @title RESTful API em Golang e SQLite
// @version 1.0

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8000
// @BasePath /
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

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Logger.Fatal(e.Start(":8000"))
}

//HealthHTTPHandler greatings
func healthHTTPHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, app.ResponseMessage("I'm good, tks!"))
}