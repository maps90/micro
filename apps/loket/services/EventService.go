package service

import (
    "fmt"

    loket "github.com/mataharimall/micro/api/loket"
    entity "github.com/mataharimall/micro/apps/loket/entities"
    helper "github.com/mataharimall/micro/helpers"
)

type EventService struct {
    Response struct {
        Result interface{}
    }
}

type Events struct {
    Status  string      `json:"status" jsmap:"status"`
    Data    []EventData `json:"data" jsmap:"data"`
    Message string      `json:"message" jsmap:"message"`
    Code    int         `json:"code" jsmap:"code"`
}

func (self *EventService) GetList() (err error) {

    api := loket.New()
    api.GetAuth()
    token := fmt.Sprintf(`{"token": "%s"}`, api.Token)
    api.Post("v3/event", "json", token)

    ev := new(entity.Events)
    api.SetStruct(ev)

    x, err := helper.JsMap(ev)
    fmt.Println(x)

    self.Response.Result = x

    return nil
}
