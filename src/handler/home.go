package handler

import (
	"net/http"

	"github.com/ShingoYadomoto/litrews/src/api"
	"github.com/labstack/echo"
)

func Home(c echo.Context) (err error) {
	topics := api.GetAllTopics()

	return c.Render(http.StatusOK, "home", map[string]interface{}{
		"topics": topics,
	})
}
