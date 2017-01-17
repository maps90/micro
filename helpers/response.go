package helpers

import (
    "fmt"
    "net/http"
    "reflect"
    "time"

    "github.com/labstack/echo"
    config "github.com/spf13/viper"
)

var StartTime time.Time

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
        if config.GetBool("debug") {
            stop := time.Now()
            ss := uint32(stop.Sub(StartTime) / time.Millisecond)
            exec_time := fmt.Sprintf("%dms", ss)
            response["execution_time"] = exec_time
        }

        return c.JSON(http.StatusOK, response)
    }
}
