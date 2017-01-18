package handlers

import (
    "encoding/json"
    "fmt"
    "net/http"

    "github.com/labstack/echo"
    "github.com/mataharimall/micro/api"
    "github.com/mataharimall/micro/container"
)

type ticket struct {
    Request  interface{}
    Response struct {
        Result interface{}
    }
}

func GetTicketsBySchedule(c echo.Context) (err error) {

    // r := &ticket{}

    // if err := c.Bind(r.Request); err != nil {
    //     return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    // }

    loket, ok := container.Get("api.loket").(*api.Loket)
    if !ok {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }

    loket.GetAuth().Post(fmt.Sprintf("/v3/tickets/%s", c.Param("scheduleID")), "form", "")

    var out interface{}
    json.Unmarshal([]byte(loket.Body), &out)

    return c.JSON(http.StatusOK, out)

}
