package middleware

import (
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

type ApiUserAuthFunc func(bearer, token string) (bool, error)

func ApiUserAuth(getAuth ApiUserAuthFunc) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			requestAuth := c.Request().Header().Get("Authorization")
			a := strings.Split(requestAuth, " ")
			he := echo.NewHTTPError(http.StatusUnauthorized)
			if len(a) != 2 {
				return he
			}
			if ok, _ := getAuth(a[0], a[1]); !ok {
				return he
			}
			return next(c)
		}
	}
}
