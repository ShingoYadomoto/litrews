package api

import (
	"strconv"

	"github.com/ShingoYadomoto/litrews/src/config"
	"github.com/pkg/errors"
)

type DocomoApi struct {
	config.DocomoApi
}

type GenreData struct {
	Genres []Genre `json:"genre"`
}

type Genre struct {
	ID          int    `json:"genreId"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type ArticleData struct {
	Total        int        `json:"totalResults"`
	StartIndex   int        `json:"startIndex"`
	CountPerPage int        `json:"itemsPerPage"`
	SetTime      string     `json:"issueDate"` // ToDo: time.Time的な型にしたい
	Articles     []DArticle `json:"articleContents"`
}

type DArticle struct {
	ID                       int    `json:"contentId"`
	ContentType              int    `json:"contentType"`
	GenreID                  int    `json:"genreId"`
	RelatedArticleDataEPoint string `json:"relatedContents"`
	Data                     struct {
		Title      string `json:"title"`
		Body       string `json:"body"`
		CreateTime string `json:"createdDate"` // ToDo: time.Time的な型にしたい
		SrcDomain  string `json:"sourceDomain"`
		SrcName    string `json:"sourceName"`
		SrcURL     string `json:"linkUrl"`
		ImageURL   string `json:"imageUrl"`
		ImageSize  struct {
			Width  int `json:"width"`
			Height int `json:"height"`
		} `json:"imageSize"`
	} `json:"contentData"`
}

func (da *DocomoApi) getEndPoint(kind string, params string) (ep string) {
	ep = da.BaseEndPoint + kind + "?APIKEY=" + da.Key + params + da.CommonParams
	return
}

func (da *DocomoApi) SetEncodedDataFromEndPoint(i interface{}, ep string) (err error) {
	err = setEncodedDataFromEndPoint(i, ep)
	return
}

func (da *DocomoApi) GetAllGenres() ([]Genre, error) {
	data := new(GenreData)

	endPoint := da.getEndPoint("genre", "")

	err := da.SetEncodedDataFromEndPoint(&data, endPoint)
	if err != nil {
		err = errors.Wrap(err, "ジャンルの取得に失敗しました。")
	}
	return data.Genres, err
}

func (da *DocomoApi) GetArticles(genreID int, count int) ([]DArticle, error) {
	data := new(ArticleData)

	queryParams := "&genreId=" + strconv.Itoa(genreID) + "&n=" + strconv.Itoa(count)
	endPoint := da.getEndPoint("contents", queryParams)

	err := da.SetEncodedDataFromEndPoint(&data, endPoint)
	if err != nil {
		err = errors.Wrap(err, "記事の取得に失敗しました。")
	}
	return data.Articles, err
}
