package handler

import (
	"net/http"
	"time"

	"github.com/ShingoYadomoto/litrews/src/api"
	"github.com/ShingoYadomoto/litrews/src/job"
	"github.com/bamzi/jobrunner"
	"github.com/labstack/echo"
)

func Home(c echo.Context) (err error) {
	topics := api.GetAllTopics()
	jobrunner.Start()
	jobrunner.In(28*time.Minute, job.Curl{})

	return c.Render(http.StatusOK, "home", map[string]interface{}{
		"topics": topics,
	})
}

func Jobjson(c echo.Context) (err error) {
	return c.JSON(200, jobrunner.StatusJson())
}
