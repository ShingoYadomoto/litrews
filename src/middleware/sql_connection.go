package middleware

import (
	"github.com/ShingoYadomoto/litrews/src/context"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func SqlDBMiddleware(conn *sqlx.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(context.DatabasesKey, conn)
			log.Debug("set sql.Conn echo.Context.")
			if err := next(c); err != nil {
				c.Error(err)
			}
			return nil
		}
	}
}
