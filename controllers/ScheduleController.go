package controller

import (
    "net/http"

    "github.com/labstack/echo"
    helper "github.com/mataharimall/micro-api/commons"
    service "github.com/mataharimall/micro-api/services"
)

func GetSchedule(c echo.Context) error {
    r := &service.ScheduleService{}

    err := r.GetSchedule(c.Param("scheduleID"))
    if err != nil {
        return echo.NewHTTPError(http.StatusNotFound, "Data not found")
    }

    return helper.BuildResponse(c, r.Response.Result, nil, err)
}
