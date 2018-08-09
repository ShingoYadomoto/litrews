package handler

import (
	"net/http"

	"github.com/ShingoYadomoto/litrews/src/api"
	"github.com/ShingoYadomoto/litrews/src/context"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func Home(c echo.Context) (err error) {
	cc := c.(*context.CustomContext)
	conf := cc.GetConfig()

	dapi := api.DocomoApi{conf.DocomoApi}

	genres, err := dapi.GetAllGenres()
	if err != nil {
		log.Error(err)
		return c.Render(http.StatusOK, "error", err)
	}

	return c.Render(http.StatusOK, "home", map[string]interface{}{
		"genres": genres,
	})
}
