package controller

import (
    "net/http"

    "github.com/labstack/echo"
    service "github.com/mataharimall/micro/apps/loket/services"
    helper "github.com/mataharimall/micro/helpers"
)

func GetEventList(c echo.Context) error {
    r := &service.EventService{}

    err := r.GetList()
    if err != nil {
        return echo.NewHTTPError(http.StatusNotFound, "Data not found")
    }

    return helper.BuildResponse(c, r.Response.Result, nil, err)
}
