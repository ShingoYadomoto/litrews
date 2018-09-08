package context

import (
	"bytes"

	"github.com/ShingoYadomoto/litrews/src/config"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func CustomContextMiddleware() echo.MiddlewareFunc {
	return func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &CustomContext{c}
			return h(cc)
		}
	}
}

type CustomContext struct {
	echo.Context
}

const (
	ConfigKey = "__CONFIG__"
)

func (c *CustomContext) GetConfig() *config.Conf {
	conf, ok := c.Get(ConfigKey).(*config.Conf)
	if !ok {
		log.Panic("*config.Conf assertion error")
	}
	return conf
}

// echo.Context.Renderをオーバーライド
func (c *CustomContext) Render(code int, name string, data interface{}) (err error) {
	if c.Echo().Renderer == nil {
		return echo.ErrRendererNotRegistered
	}

	// render時に各view共通で渡す変数を格納
	if d, isMap := data.(map[string]interface{}); isMap {
		conf := c.GetConfig()
		d["googleAnalyticsID"] = conf.GoogleApi.AnalyticsID
	}

	buf := new(bytes.Buffer)
	if err = c.Echo().Renderer.Render(buf, name, data, c); err != nil {
		return
	}
	return c.HTMLBlob(code, buf.Bytes())
}
