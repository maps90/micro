package handlers

import (
	"github.com/labstack/echo"
	"github.com/mataharimall/micro/container"
	"github.com/mataharimall/micro/api"
	"github.com/mataharimall/micro/helpers"
	"encoding/json"
	"net/http"
	"fmt"
)

func PaidInvoice(c echo.Context) error {
	loket, ok := container.Get("api.loket").(*api.Loket)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	url := fmt.Sprintf(`/v3/invoice/%s/paid`, c.Param("code"))
	loket.GetAuth().Post(url, "form", "")
	var m map[string]interface{}
	json.Unmarshal([]byte(loket.Body), &m)

	return helpers.BuildJSON(c, m)
}
