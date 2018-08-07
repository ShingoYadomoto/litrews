package main

import (
	"flag"
	"html/template"
	"io"
	"os"
	"strconv"

	"github.com/ShingoYadomoto/Litrews/src/config"
	"github.com/ShingoYadomoto/Litrews/src/context"
	"github.com/ShingoYadomoto/Litrews/src/handler"
	"github.com/ShingoYadomoto/Litrews/src/middleware"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	echo_middleware "github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	var (
		confPath string
	)

	workingDirPath, _ := os.Getwd()
	defaultConfigDirPath := workingDirPath + "/../env/config.yml"

	flag.StringVar(&confPath, "conf", defaultConfigDirPath, "config file path")
	flag.StringVar(&confPath, "c", defaultConfigDirPath, "config file path")
	flag.Parse()

	conf := config.Load(confPath)

	mysqlConf := mysql.Config{
		User:      conf.Database.User,
		Passwd:    conf.Database.Password,
		Net:       conf.Database.Net,
		Addr:      conf.Database.Addr,
		DBName:    conf.Database.DBName,
		ParseTime: true,
	}
	dsn := mysqlConf.FormatDSN()

	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	e := initEcho(conf, db)

	e.Debug = true

	e.GET("/", handler.Home)

	// Start server
	address := ":" + strconv.Itoa(conf.App.Port)
	e.Logger.Fatal(e.Start(address))
}

func initEcho(conf *config.Conf, db *sqlx.DB) *echo.Echo {
	// Setup
	e := echo.New()

	e.Logger.SetLevel(conf.Log.Level)
	log.SetLevel(conf.Log.Level)

	e.Static("/static", "../resources/assets")

	e.Use(context.CustomContextMiddleware())
	e.Use(middleware.ConfigMiddleware(conf))
	e.Use(middleware.SqlDBMiddleware(db))
	e.Use(echo_middleware.Logger())
	e.Use(echo_middleware.Recover())

	e.Renderer = &Template{
		templates: template.Must(template.ParseGlob("../resources/views/**/*.html")),
	}

	return e
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
