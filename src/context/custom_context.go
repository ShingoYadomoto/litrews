package context

import (
	"github.com/ShingoYadomoto/litrews/src/config"
	cdb "github.com/ShingoYadomoto/litrews/src/db"

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
