package loket

import (
	"github.com/labstack/echo"
	em "github.com/labstack/echo/middleware"
	"github.com/mataharimall/micro"
	ctrl "github.com/mataharimall/micro-api/apps/loket/controllers"
)

type LoketRoute struct{}

func init() {
	micro.RouterManager.Register("route.loket", &LoketRoute{})
}

func (l *LoketRoute) SetRoute(e *echo.Echo) *echo.Echo {
	e.Use(em.Logger())
	e.Post("/loket/event", ctrl.GetEventList)
	e.Post("/loket/schedule/:scheduleID", ctrl.GetSchedule)
	e.Post("/loket/tickets/:scheduleID", ctrl.GetTicketsBySchedule)
	return e
}
