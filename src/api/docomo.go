package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

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

type Article struct {
	ID int `json:"id"`
}

func (da *DocomoApi) getEndPoint(kind string, params string) (ep string) {
	ep = da.BaseEndPoint + kind + "?APIKEY=" + da.Key + params + da.CommonParams
	return
}

func (da *DocomoApi) SetEncodedDataFromEndPoint(i interface{}, ep string) (err error) {
	// URLを叩いてデータを取得
	resp, err := http.Get(ep)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	// 取得したデータをJSONデコード
	err = json.Unmarshal(body, i)
	if err != nil {
		return
	}
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
