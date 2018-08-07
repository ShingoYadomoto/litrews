package middleware

import (
	"github.com/ShingoYadomoto/Litrews/src/config"
	"github.com/ShingoYadomoto/Litrews/src/context"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func ConfigMiddleware(config *config.Conf) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(context.ConfigKey, config)
			log.Debug("set config to echo.Context.")
			if err := next(c); err != nil {
				c.Error(err)
			}
			return nil
		}
	}
}
