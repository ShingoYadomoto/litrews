package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// mappingしたいstructとjsonを返すapiのエンドポイントを渡すと、mappingする
func setEncodedDataFromEndPoint(i interface{}, ep string) (err error) {
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
