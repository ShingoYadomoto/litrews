package handler

import (
	"net/http"

	"strconv"

	"github.com/ShingoYadomoto/litrews/src/context"
	"github.com/ShingoYadomoto/litrews/src/model"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func Home(c echo.Context) (err error) {
	cc := c.(*context.CustomContext)
	db := cc.GetDB()

	topicModel := model.NewTopicModel(db)

	topics, err := topicModel.GetAllTopics()
	if err != nil {
		log.Error(err)
		return c.Render(http.StatusOK, "error", err)
	}

	newsLink := map[string]string{}
	for _, topic := range topics {
		newsLink[topic.NameJa] = "/google_news/" + strconv.Itoa((topic.ID))
	}

	return c.Render(http.StatusOK, "home", map[string]interface{}{
		"newsLink": newsLink,
	})
}
