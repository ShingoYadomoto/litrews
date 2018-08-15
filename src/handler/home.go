package handler

import (
	"net/http"
	"net/url"

	"html/template"

	"github.com/ShingoYadomoto/litrews/src/api"
	"github.com/ShingoYadomoto/litrews/src/context"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func Home(c echo.Context) (err error) {
	newsLink := map[string]string{}
	for topicJa, topicEn := range api.GoogleNewsTopic {
		newsLink[topicEn] = "/google_news/" + topicJa
	}

	return c.Render(http.StatusOK, "home", map[string]interface{}{
		"newsLink": newsLink,
	})
}

func GoogleNews(c echo.Context) (err error) {
	cc := c.(*context.CustomContext)
	conf := cc.GetConfig()

	topic := c.Param("topic")
	googleNewsEndPoint := api.GetGoogleNewsEndPoint(topic)
	googleNewsEndPoint = url.QueryEscape(googleNewsEndPoint)

	articlesData := new(api.RssData)

	rss2JsonApi := api.Rss2JsonApi{conf.Rss2JsonApi}
	rss2JsonEndPoint := rss2JsonApi.GetEndPoint(googleNewsEndPoint)

	rss2JsonApi.SetEncodedDataFromEndPoint(articlesData, rss2JsonEndPoint)

	for i, article := range articlesData.Articles {
		articlesData.Articles[i].Body = template.HTML(article.Description)
	}

	return c.Render(http.StatusOK, "googleNews", map[string]interface{}{
		"articlesData": articlesData,
	})
}

func Docomo(c echo.Context) (err error) {
	cc := c.(*context.CustomContext)
	conf := cc.GetConfig()

	dapi := api.DocomoApi{conf.DocomoApi}

	genres, err := dapi.GetAllGenres()
	if err != nil {
		log.Error(err)
		return c.Render(http.StatusOK, "error", err)
	}

	favGenre := genres[0]
	articles, err := dapi.GetArticles(favGenre.ID, 10)
	if err != nil {
		log.Error(err)
		return c.Render(http.StatusOK, "error", err)
	}

	return c.Render(http.StatusOK, "docomo", map[string]interface{}{
		"genres":   genres,
		"articles": articles,
	})
}
