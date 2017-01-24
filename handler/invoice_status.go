package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/maps90/librarian"
	"github.com/mataharimall/micro/api"
	"github.com/mataharimall/micro/helper"
)

func GetInvoiceStatus(c echo.Context) error {
	loket, ok := librarian.Get("loket").(*api.Loket)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	url := fmt.Sprintf(`/v3/invoice/%s`, c.Param("invoice_code"))
	loket.Post(url, "form", "")
	return helper.BuildJSON(c, loket.Response.Data, loket.Error)
}

func PostInvoicePaid(c echo.Context) error {
	loket, ok := librarian.Get("loket").(*api.Loket)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	url := fmt.Sprintf(`/v3/invoice/%s/paid`, c.Param("code"))
	loket.Post(url, "form", "")
	return helper.BuildJSON(c, loket.Response.Data, loket.Error)
}
