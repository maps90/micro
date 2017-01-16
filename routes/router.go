package routes

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	controller "github.com/mataharimall/micro-api/controllers"
	appMiddleware "github.com/mataharimall/micro-api/routes/middleware"
	config "github.com/spf13/viper"
)

func SetRoute() *echo.Echo {

	//auth := appMiddleware.Auth()
	/*start echo*/
	e := echo.New()
	// Set MiddleWare
	if config.GetBool("debug") {
		e.Use(appMiddleware.Logger())
	}
	e.Use(middleware.Recover(), middleware.Gzip())
	e.SetBinder(appMiddleware.AppBinder{})
	e.SetHTTPErrorHandler(appMiddleware.AppHttpErrorHandler)

	//routing goes here
	event := e.Group("/event")
	//event.use(auth)
	event.Post("/list", controller.GetEventList)

	return e
}
