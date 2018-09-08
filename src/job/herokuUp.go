package job

import (
	"net/http"

	"github.com/labstack/gommon/log"
)

type Curl struct {
	URL string
}

func (e Curl) Run() {
	_, err := http.Get(e.URL)
	if err != nil {
		log.Error(err)
		return
	}
	return
}
