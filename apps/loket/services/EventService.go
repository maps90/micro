package service

import (
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
    Status  string             `json:"status" jsmap:"status"`
    Data    []entity.EventData `json:"data" jsmap:"data"`
    Message string             `json:"message" jsmap:"message"`
    Code    int                `json:"code" jsmap:"code"`
}

func (self *EventService) GetList() (err error) {

    ev := new(Events)

    loket.New().GetAuth().
        Post("/v3/event", "form", "").
        SetStruct(ev)

    x, err := helper.JsMap(ev)

    self.Response.Result = x

    return nil
}
