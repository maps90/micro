package loket

import (
    "github.com/labstack/echo"
    "github.com/mataharimall/micro/apps/loket/handlers"
    "github.com/mataharimall/micro/middleware"
    "github.com/mataharimall/micro/service"
)

type LoketRoute struct{}

func init() {
    service.ServiceManager.Register("route.loket", &LoketRoute{})
}

func (l *LoketRoute) SetRoute(s *echo.Echo) *echo.Echo {
    s.Use(middleware.Logger())
    s.SetBinder(middleware.AppBinder{})
    s.SetHTTPErrorHandler(middleware.AppHttpErrorHandler)

    s.Get("/loket/event/list", handlers.GetEventList)
    s.Get("/loket/schedule/:scheduleID", handlers.GetSchedule)
    s.Get("/loket/tickets/:scheduleID", handlers.GetTicketsBySchedule)
    return s
}
