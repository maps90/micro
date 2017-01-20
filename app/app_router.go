package app

import (
	"github.com/facebookgo/grace/gracehttp"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/mataharimall/micro/handler"
	"github.com/mataharimall/micro/middleware"
	c "github.com/spf13/viper"
)

func initRouter() error {
	e := echo.New()
	e.SetDebug(c.GetBool("app.debug"))
	e.Use(middleware.Logger())

	e.SetBinder(middleware.AppBinder{})
	e.SetHTTPErrorHandler(middleware.AppHttpErrorHandler)

	e.Get("/event", handler.GetEventList)
	e.Get("/event/search", handler.SearchEvent)
	e.Get("/schedule/:schedule_id", handler.GetScheduleById)
	e.Get("/tickets/:schedule_id", handler.GetTicketsBySchedule)
	e.Post("/invoice/paid/:code", handler.PostInvoicePaid)
	e.Post("/invoice/create", handler.CreateInvoice)
	e.Get("/invoice/:invoice_code", handler.GetInvoiceStatus)
	e.Post("/invoice/list/:invoice_code/attendee", handler.FetchInvoiceListAttendee)

	std := standard.New(":" + c.GetString("app.port"))
	std.SetHandler(e)

	err := gracehttp.Serve(std.Server)
	return err
}
