package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/maps90/librarian"
	"github.com/mataharimall/micro/api"
	"github.com/mataharimall/micro/helper"
)

func PostInvoiceStatus(c echo.Context) error {
	loket, ok := librarian.Get("api.loket").(*api.Loket)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	url := fmt.Sprintf(`/v3/invoice/%s/paid`, c.Param("code"))
	loket.GetAuth().Post(url, "form", "")
	return helper.BuildJSON(c, loket.Response.Data, loket.Error)
}
