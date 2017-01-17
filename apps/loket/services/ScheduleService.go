package service

import (
    "fmt"

    loket "github.com/mataharimall/micro/api/loket"
    entity "github.com/mataharimall/micro/apps/loket/entities"
    helper "github.com/mataharimall/micro/helpers"
)

type ScheduleService struct {
    Response struct {
        Result interface{}
    }
}

type Schedules struct {
    Status  string                `json:"status" jsmap:"status"`
    Data    []entity.ScheduleData `json:"data" jsmap:"data"`
    Message string                `json:"message" jsmap:"message"`
    Code    int                   `json:"code" jsmap:"code"`
}

func (self *ScheduleService) GetSchedule(ScheduleID string) (err error) {

    ev := new(Schedules)

    loket.New().GetAuth().
        Post(fmt.Sprintf("/v3/schedule/%s", ScheduleID), "form", "").
        SetStruct(ev)

    x, err := helper.JsMap(ev)

    self.Response.Result = x

    return nil
}
