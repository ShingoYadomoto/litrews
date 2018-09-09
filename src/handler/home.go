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

	topics := api.GetAllTopics()

	return c.Render(http.StatusOK, "home", map[string]interface{}{
		"topics": topics,
	})
}

func Jobjson(c echo.Context) (err error) {
	conf := c.(*context.CustomContext).GetConfig()

	jobrunner.Start()
	jobrunner.In(25*time.Minute, job.Curl{conf.App.URL})

	return c.JSON(200, jobrunner.StatusJson())
}
