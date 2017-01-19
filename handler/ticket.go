package handler

import (
    "encoding/json"
    "fmt"
    "net/http"

    "github.com/labstack/echo"
    "github.com/maps90/librarian"
    "github.com/mataharimall/micro/api"
    "github.com/mataharimall/micro/helper"
)

func GetTicketsBySchedule(c echo.Context) (err error) {

    loket, ok := librarian.Get("loket").(*api.Loket)
    if !ok {
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }

    url := fmt.Sprintf(`/v3/tickets/%s`, c.Param("schedule_id"))
    loket.GetAuth().Post(url, "form", "")

    var m map[string]interface{}
    json.Unmarshal([]byte(loket.Body), &m)

    return helper.BuildJSON(c, m)

}
