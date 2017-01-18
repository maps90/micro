package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/mataharimall/micro/api"
	"github.com/mataharimall/micro/container"
)

type eventsList struct {
	Request  interface{}
	Response struct{}
}

func GetEventList(c echo.Context) error {
	/*r := &eventsList{}

	if err := c.Bind(r.Request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}*/

	loket := container.Get("api.loket").(*api.Loket)
	loket.GetAuth().Post("/v3/event", "form", "")

	return c.JSON(http.StatusOK, loket.Body)
}
