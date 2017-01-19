package handler

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
	"github.com/maps90/librarian"
	"github.com/mataharimall/micro/api"
	"github.com/mataharimall/micro/helper"
)

type CreateInvoiceRequestResponse struct {
	Request  CreateInvoiceRequest
	Response interface{}
}

type CreateInvoiceRequest struct {
	Data Data `json:"data"`
}

type Data struct {
	Tickets        []Ticket `json:"tickets"`
	Attendee       Attendee `json:"attendee"`
	OrderId        string   `json:"order_id"`
	ExpirationType string   `json:"expiration_type"`
	Notes          string   `json:"notes"`
}

type Ticket struct {
	IdTicket string `json:"id_ticket"`
	Quantity string `json:"qty"`
}

type Attendee struct {
	FirstName  string `json:"firstname"`
	LastName   string `json:"lastname"`
	IdentityId string `json:"identity_id"`
	Dob        string `json:"dob"`
	Gender     string `json:"gender"`
	Email      string `json:"email"`
	Telephone  string `json:"telephone"`
}

func CreateInvoice(c echo.Context) error {
	r := CreateInvoiceRequestResponse{}

	if err := c.Bind(&r.Request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	loket, ok := librarian.Get("loket").(*api.Loket)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	jbyte, err := json.Marshal(r.Request)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	loket.GetAuth().Post("/v3/invoice/create", "json", string(jbyte))
	return helper.BuildJSON(c, loket.Response.Data, loket.Error)
}
