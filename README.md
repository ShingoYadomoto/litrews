# Litrews
Litrews using docomo and line APIs

### API仕様

#### [docomo トレンド記事抽出API](https://dev.smt.docomo.ne.jp/?p=docs.api.page&api_name=trend_article_extraction&p_name=api_usage_scenario)

 - ジャンルの取得
   - エンドポイント:https://api.apigw.smt.docomo.ne.jp/webCuration/v3/genre

`リクエストクエリパラメータ`

キー    | 必須 | 説明
------ | ---  | -----------------------------------------------------------------------------------
APIKEY | ○    | APIにアクセスするアプリの認証に利用する。
lang   | -    | 記事およびジャンル情報の言語は、下記のいずれかを指定。<br>ja : 日本語(デフォルト)<br>  en : 英語

`レスポンスJSON例`
```
{
"genre": [
{
      "genreId": 1,
      "title": "スポーツ",
      "description" : "スポーツニュースをお届けします！"

},
{
      "genreId": 2,
      "title": "グルメ"
},
...
  ]
}
```

- 記事の取得
  - エンドポイント:https://api.apigw.smt.docomo.ne.jp/webCuration/v3/contents

`リクエストクエリパラメータ`

キー     | 必須 | 説明                                                                                            
------- | ---- | -------------------------------------------------------------------------------------
APIKEY  | ○    | APIにアクセスするアプリの認証に利用する。
genreId | ○    | 表示設定されているジャンルIDを指定。                                                                           
lang    | -    | 記事およびジャンル情報の言語は、下記のいずれかを指定。<br>ja : 日本語(デフォルト) <br>en : 英語
s       | -    | 記事一覧の開始番号を指定(1以上999999以下の整数)。<br>デフォルト : 1
n       | -    | カテゴリ内の記事一覧の取得件数を指定(0以上50以下の整数)。<br>デフォルト : 10

`レスポンスJSON例`
```
{
"totalResults" : 10,
"startIndex" : 1,
"itemsPerPage" : 10,
"issueDate" : "2013-05-01T11:11:11+0900",
"articleContents" :[
  {
    "contentId": 1127632,
    "contentType": 10,
    "genreId": 1,
    "contentData": {
      "title": "ビットコインETFの次の試練は今週末までに来る　ヴァンエック・ソリッドエックス・ビットコイン・トラストに米国証券取引委員会が意見を表明する予定 ",
      "body": "今週末までに米国証券取引委員会（SEC）はヴァンエック・ソリッドエックス・ビットコイン・トラストというビットコインETFの上場申請に対して回答をする予定です。 承認されるかもしれないし、されないかも…",
      "createdDate": "2018-08-08T03:23:00+0900",
      "sourceDomain": "markethack.net",
      "sourceName": "Market Hack",
      "linkUrl": "http://markethack.net/archives/52084939.html"
    }
  },
….
  ]
}
```








1
