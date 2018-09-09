package job

import (
	"net/http"

	"github.com/bamzi/jobrunner"
	"github.com/labstack/gommon/log"
)

type Curl struct {
	URL string
}

func HerokuUp(url string) {
	jobrunner.Start()
	jobrunner.Schedule("@every 3s", Curl{url})
}

func (e Curl) Run() {
	_, err := http.Get(e.URL)
	if err != nil {
		log.Error(err)
		return
	}
	return
}
