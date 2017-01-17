package service

import (
    "fmt"

    helper "github.com/mataharimall/micro-api/commons"
    loket "github.com/mataharimall/micro-api/components/Loket"
    entity "github.com/mataharimall/micro-api/entities"
)

type ScheduleService struct {
    Response struct {
        Result interface{}
    }
}

func (self *ScheduleService) GetSchedule(ScheduleID string) (err error) {

    api := loket.New()
    api.GetAuth()
    token := fmt.Sprintf(`{"token": "%s"}`, api.Token)

    endpoint := fmt.Sprintf("schedule/%s", ScheduleID)

    api.Post("v3", endpoint, token)

    ev := new(entity.Schedules)
    api.SetStruct(ev)

    x, err := helper.JsMap(ev)
    fmt.Println(x)

    self.Response.Result = x

    return nil
}
