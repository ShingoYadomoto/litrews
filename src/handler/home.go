package handler

import (
	"net/http"
	"time"

	"github.com/ShingoYadomoto/litrews/src/api"
	"github.com/ShingoYadomoto/litrews/src/context"
	"github.com/ShingoYadomoto/litrews/src/job"
	"github.com/bamzi/jobrunner"
	"github.com/labstack/echo"
)

func Home(c echo.Context) (err error) {
	conf := c.(*context.CustomContext).GetConfig()

	topics := api.GetAllTopics()
	jobrunner.Start()
	jobrunner.In(25*time.Minute, job.Curl{conf.App.URL})

	return c.Render(http.StatusOK, "home", map[string]interface{}{
		"topics": topics,
	})
}

func Jobjson(c echo.Context) (err error) {
	return c.JSON(200, jobrunner.StatusJson())
}
