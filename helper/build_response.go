package helper

import (
	"net/http"
	"reflect"

	"github.com/labstack/echo"
)

func BuildJSON(c echo.Context, rs interface{}, err error) (error) {
	reqId := RandomString(20)
	rsLen := rs == nil
	val := reflect.ValueOf(rs)
	if val.Kind() == reflect.Ptr && val.IsNil() || rsLen {
		rs = make([]string, 0)
	}
	if err != nil {
		c.Error(err)
		return err
	} else {

		response := map[string]interface{}{
			"code":      http.StatusOK,
			"requestId": reqId,
			"results":   rs,
		}

		return c.JSON(http.StatusOK, response)
	}
}

func BuildResponse(c echo.Context, rs interface{}, total interface{}, err error) error {
	reqId := RandomString(20)
	rsLen := rs == nil
	val := reflect.ValueOf(rs)
	if val.Kind() == reflect.Ptr && val.IsNil() || rsLen {
		rs = make([]string, 0)
	}
	if err != nil {
		c.Error(err)
		return err
	} else {

		response := map[string]interface{}{
			"code":      http.StatusOK,
			"requestId": reqId,
			"results":   rs,
		}
		if total != nil {
			response["total"] = total
		}

		return c.JSON(http.StatusOK, response)
	}
}
