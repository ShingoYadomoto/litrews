package main

import (
	"html/template"
	"io"

	"os"

	"strconv"

	"github.com/ShingoYadomoto/litrews/src/config"
	"github.com/ShingoYadomoto/litrews/src/context"
	"github.com/ShingoYadomoto/litrews/src/handler"
	"github.com/ShingoYadomoto/litrews/src/middleware"
	"github.com/labstack/echo"
	echo_middleware "github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

func main() {

	conf := config.GetConfig()

	e := initEcho(&conf)

	e.Debug = true

	e.GET("/", handler.Home)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = strconv.Itoa(conf.App.Port)
	}
	address := ":" + port
	e.Logger.Fatal(e.Start(address))
}

func initEcho(conf *config.Conf) *echo.Echo {
	// Setup
	e := echo.New()

	e.Logger.SetLevel(conf.Log.Level)
	log.SetLevel(conf.Log.Level)

	e.Static("/static", "resources/assets")

	e.Use(context.CustomContextMiddleware())
	e.Use(middleware.ConfigMiddleware(conf))
	e.Use(echo_middleware.Logger())
	e.Use(echo_middleware.Recover())

	e.Renderer = &Template{
		templates: template.Must(template.ParseGlob("resources/views/**/*.html")),
	}

	return e
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
