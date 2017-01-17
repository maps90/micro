package loket

import (
	"github.com/labstack/echo"
	"github.com/mataharimall/micro"
	ctrl "github.com/mataharimall/micro/apps/loket/controllers"
	mm "github.com/mataharimall/micro/middleware"
)

type LoketRoute struct{}

func init() {
	micro.RouterManager.Register("route.loket", &LoketRoute{})
}

func (l *LoketRoute) SetRoute(e *echo.Echo) *echo.Echo {
	e.Use(mm.Logger())
	e.Post("/loket/event/list", ctrl.GetEventList)
	e.Post("/loket/schedule/:scheduleID", ctrl.GetSchedule)
	e.Post("/loket/tickets/:scheduleID", ctrl.GetTicketsBySchedule)
	return e
}
