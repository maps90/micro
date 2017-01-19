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

type EventsList struct {
	Request  interface{}
	Response interface{}
}

type SearchEventPayload struct {
	Data struct {
		Search struct {
			EventName    string `json:"event_name" query:"event_name"`
			LocationName string `json:"location_name" query:"location_name"`
			MinStartDate string `json:"min_start_date" query:"min_start_date"`
			MaxEndDate   string `json:"max_end_date" query:"max_end_date"`
		} `json:"search"`
	} `json:"data"`
	Limit  string `json:"limit"`
	Offset string `json:"offset"`
}

func GetEventList(c echo.Context) error {
	loket, ok := librarian.Get("loket").(*api.Loket)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")

	}

	loket.GetAuth().Post("/v3/event", "form", "")
	var m map[string]interface{}
	json.Unmarshal([]byte(loket.Body), &m)

	return helper.BuildJSON(c, m)

}

func SearchEvent(c echo.Context) error {
	loket, ok := librarian.Get("loket").(*api.Loket)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	s := SearchEventPayload{}

	if err := c.Bind(&s); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	search_json, _ := json.Marshal(s)
	fmt.Println(string(search_json))

	loket.GetAuth().Post("/v3/event_search", "json", string(search_json))
	var m map[string]interface{}
	json.Unmarshal([]byte(loket.Body), &m)

	return helper.BuildJSON(c, m)

}
