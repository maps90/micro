package app

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/fasthttp"
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
	e.Post("/invoice/paid/:code", handler.PostInvoiceStatus)
	e.Post("/invoice/create", handler.CreateInvoice)
	e.Post("/invoice/list/:invoice_code/attendee", handler.FetchInvoiceListAttendee)

	err := e.Run(fasthttp.New(":" + c.GetString("app.port")))
	return err
}
