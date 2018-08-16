package context

import (
	"bytes"

	"github.com/ShingoYadomoto/litrews/src/config"
	cdb "github.com/ShingoYadomoto/litrews/src/db"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"

	"github.com/jmoiron/sqlx"
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
	ConfigKey    = "__CONFIG__"
	DatabasesKey = "__DATABASES__"
)

func (c *CustomContext) GetConfig() *config.Conf {
	conf, ok := c.Get(ConfigKey).(*config.Conf)
	if !ok {
		log.Panic("*config.Conf assertion error")
	}
	return conf
}

func (c *CustomContext) GetDB() cdb.AbstractDB {
	db, ok := c.Get(DatabasesKey).(*sqlx.DB)
	if !ok {
		log.Panic("sql.Conn assertion error")
	}

	return cdb.NewSchemaContext(db).DB()
}

func (c *CustomContext) GetSession() *sessions.Session {
	s, err := session.Get("sess", c)
	if err != nil {
		log.Panic("session error")
	}

	return s
}

func (c *CustomContext) SaveValidationErrors(err error) {
	//errs := cvalidator.CustomErrorMessages(err)
	s := c.GetSession()

	//ToDo: flashメッセージにstructを保存して正常に動くか調査
	//for _, err := range errs {
	//	s.AddFlash(err, "error")
	//}
	s.AddFlash(err.Error(), "error")
	s.Save(c.Request(), c.Response())
}

// echo.Context.Renderをオーバーライド
func (c *CustomContext) Render(code int, name string, data interface{}) (err error) {
	if c.Echo().Renderer == nil {
		return echo.ErrRendererNotRegistered
	}

	// render時に各view共通で渡す変数を格納
	if d, isMap := data.(map[string]interface{}); isMap {
		s := c.GetSession()
		d["errors"], d["flashes"] = s.Flashes("error"), s.Flashes()
		s.Save(c.Request(), c.Response())
	}

	buf := new(bytes.Buffer)
	if err = c.Echo().Renderer.Render(buf, name, data, c); err != nil {
		return
	}
	return c.HTMLBlob(code, buf.Bytes())
}
