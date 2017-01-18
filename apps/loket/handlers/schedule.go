package handlers

import (
    "fmt"
    "net/http"

    "github.com/labstack/echo"
    "github.com/mataharimall/micro/api"
    "github.com/mataharimall/micro/container"
)

type schedule struct {
    Request  interface{}
    Response struct {
        Result interface{}
    }
}

func GetSchedule(c echo.Context) (err error) {

    r := &schedule{}

    if err := c.Bind(r.Request); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }

    loket := container.Get("api.loket").(*api.Loket)
    loket.GetAuth().Post(fmt.Sprintf("/v3/schedule/%s", c.Param("scheduleID")), "form", "")

    return c.JSON(http.StatusOK, loket.Body)
}
