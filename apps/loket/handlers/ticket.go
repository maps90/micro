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

    loket, ok := container.Get("api.loket").(*api.Loket)
    if !ok {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }

    loket.GetAuth().Post(fmt.Sprintf("/v3/tickets/%s", c.Param("scheduleID")), "form", "")

    var m map[string]interface{}
    json.Unmarshal([]byte(loket.Body), &m)

    return helpers.BuildJSON(c, m)

}
