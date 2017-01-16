package service

import (
    "fmt"
    "reflect"
    "strings"

    loket "github.com/mataharimall/micro-api/components/Loket"
)

type Events struct {
    Status  string      `json:"status"`
    Data    []EventData `json:"data"`
    Message string      `json:"message"`
    Code    int         `json:"code"`
}

type EventData struct {
    EventID    string          `json:"id_event"`
    Name       string          `json:"event_name"`
    Banner     string          `json:"event_banner"`
    CustomForm bool            `json:"custom_form"`
    Schedules  []EventSchedule `json:"schedules"`
}

type EventSchedule struct {
    ScheduleID   string       `json:"id_schedule"`
    Name         string       `json:"name"`
    LocationName string       `json:"location_name"`
    Address      string       `json:"address"`
    Province     string       `json:"province_name"`
    Region       string       `json:"region_name"`
    District     string       `json:"district_name"`
    Latitude     string       `json:"latitude"`
    Longitude    string       `json:"longitude"`
    StartDate    string       `json:"start_date"`
    EndDate      string       `json:"end_date"`
    TicketTypes  []TicketType `json:"ticket_types"`
}

type TicketType struct {
    TicketID  string `json:"id_ticket"`
    Type      string `json:"ticket_type"`
    Price     string `json:"price"`
    Booked    string `json:"booked"`
    Sold      string `json:"sold"`
    Available string `json:"available"`
}

type EventService struct {
    Response struct {
        Result interface{}
    }
}

func ToMap(in interface{}, tag string) (map[string]interface{}, error) {
    out := make(map[string]interface{})

    v := reflect.ValueOf(in)

    if v.Kind() == reflect.Ptr {
        v = v.Elem()
    }

    if v.Kind() != reflect.Struct {
        return nil, fmt.Errorf("ToMap only accepts structs; got %T", v)
    }

    for i := 0; i < v.NumField(); i++ {
        // gets us a StructField
        fi := v.Type().Field(i)
        valueField := v.Field(i)

        tagv := fi.Tag.Get(tag)
        tagv = strings.ToLower(fi.Name)

        if tagv != "" {
            // set key of map to value in struct field
            if valueField.Kind() == reflect.Slice {
                var xx []interface{}

                s := reflect.ValueOf(valueField.Interface())

                for i := 0; i < s.Len(); i++ {
                    x, err := ToMap(s.Index(i).Interface(), "json")
                    if err != nil {
                        fmt.Println("error boy")
                    }
                    xx = append(xx, x)
                }

                out[tagv] = xx
            } else if valueField.Kind() == reflect.Struct {
                x, err := ToMap(valueField, "json")
                if err != nil {
                    fmt.Println("error boy")
                }
                out[tagv] = x
            } else {
                out[tagv] = v.Field(i).Interface()
            }
        }
    }
    return out, nil
}

func (self *EventService) List() (err error) {

    api := loket.New()
    api.GetAuth()
    token := fmt.Sprintf(`{"token": "%s"}`, api.Token)
    api.Post("v3", "event", token)
    ev := new(Events)
    api.SetStruct(ev)

    x, err := ToMap(ev, "json")
    fmt.Println(x)

    self.Response.Result = x

    return nil
}
