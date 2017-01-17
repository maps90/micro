package service

import (
    "fmt"

    loket "github.com/mataharimall/micro/api/loket"
    entity "github.com/mataharimall/micro/apps/loket/entities"
    helper "github.com/mataharimall/micro/helpers"
)

type TicketService struct {
    Response struct {
        Result interface{}
    }
}

type Tickets struct {
    Status  string            `json:"status" jsmap:"status"`
    Data    entity.TicketData `json:"data" jsmap:"data"`
    Message string            `json:"message" jsmap:"message"`
    Code    int               `json:"code" jsmap:"code"`
}

func (self *TicketService) GetTicketsBySchedule(ScheduleID string) (err error) {

    ev := new(Tickets)

    loket.New().GetAuth().
        Post(fmt.Sprintf("/v3/tickets/%s", ScheduleID), "form", "").
        SetStruct(ev)

    x, err := helper.JsMap(ev)

    self.Response.Result = x

    return nil
}
