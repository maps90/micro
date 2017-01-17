package service

import (
    "fmt"

    helper "github.com/mataharimall/micro-api/commons"
    loket "github.com/mataharimall/micro-api/components/Loket"
    entity "github.com/mataharimall/micro-api/entities"
)

type TicketService struct {
    Response struct {
        Result interface{}
    }
}

func (self *TicketService) GetTicketsBySchedule(ScheduleID string) (err error) {

    api := loket.New()
    api.GetAuth()
    token := fmt.Sprintf(`{"token": "%s"}`, api.Token)

    endpoint := fmt.Sprintf("tickets/%s", ScheduleID)
    api.Post("v3", endpoint, token)

    ev := new(entity.Tickets)

    api.SetStruct(ev)

    fmt.Println(ev)

    x, err := helper.JsMap(ev)
    fmt.Println(x)

    self.Response.Result = x

    return nil
}
