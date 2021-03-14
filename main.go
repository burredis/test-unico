package main

import (
	"net/http"
	"unico/app"
	"unico/app/repository"
	"unico/feiralivre"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.com/swaggo/echo-swagger/example/docs"
)

// @title unico
// @description "A RESTful API made with Golang and SQLite"
// @version 1.0

// @contact.email burredis@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8000
// @BasePath /
func main() {
	db := repository.Conn()
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
	e.File("/swagger/doc.json", "./docs/swagger.json") // FIX TO LOAD CORRECT FILE JSON
	e.Logger.Fatal(e.Start(":8000"))
}

// healthHTTPHandler godoc
// @Description Check the api health
// @Tags API
// @Accept json
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Router / [get]
//HealthHTTPHandler greatings
func healthHTTPHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, app.ResponseMessage("I'm good, tks!"))
}
