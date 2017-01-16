package controller

import (
    "net/http"

    "github.com/labstack/echo"
    helper "github.com/mataharimall/micro-api/commons"
    service "github.com/mataharimall/micro-api/services"
)

func GetEventList(c echo.Context) error {
    r := &service.EventService{}

    err := r.List()
    if err != nil {
        return echo.NewHTTPError(http.StatusNotFound, "Data not found")
    }

    return helper.BuildResponse(c, r.Response.Result, nil, err)
}
