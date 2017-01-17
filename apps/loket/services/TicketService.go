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
    Status  string     `json:"status" jsmap:"status"`
    Data    TicketData `json:"data" jsmap:"data"`
    Message string     `json:"message" jsmap:"message"`
    Code    int        `json:"code" jsmap:"code"`
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
