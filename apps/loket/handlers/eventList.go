package handlers

import (
	"encoding/json"
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

	loket, ok := container.Get("api.loket").(*api.Loket)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	loket.GetAuth().Post("/v3/event", "form", "")

	var out interface{}
	json.Unmarshal([]byte(loket.Body), &out)
	return c.JSON(http.StatusOK, out)

}
