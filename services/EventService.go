package service

import (
    "fmt"
    "reflect"

    loket "github.com/mataharimall/micro-api/components/Loket"
)

type Events struct {
    Status  string      `json:"status" jsmap:"status"`
    Data    []EventData `json:"data" jsmap:"data"`
    Message string      `json:"message" jsmap:"message"`
    Code    int         `json:"code" jsmap:"code"`
}

type EventData struct {
    EventID   string          `json:"id_event" jsmap:"id"`
    Name      string          `json:"event_name" jsmap:"name"`
    Schedules []EventSchedule `json:"schedules" jsmap:"schedules"`
}

type EventSchedule struct {
    ScheduleID   string       `json:"id_schedule" jsmap:"id"`
    Name         string       `json:"name" jsmap:"name"`
    LocationName string       `json:"location_name" jsmap:"location_name"`
    Address      string       `json:"address" jsmap:"address"`
    Province     string       `json:"province_name" jsmap:"province"`
    Region       string       `json:"region_name" jsmap:"region"`
    District     string       `json:"district_name" jsmap:"district"`
    Latitude     string       `json:"latitude" jsmap:"latitude"`
    Longitude    string       `json:"longitude" jsmap:"longitude"`
    StartDate    string       `json:"start_date" jsmap:"start_date"`
    EndDate      string       `json:"end_date" jsmap:"end_date"`
    TicketTypes  []TicketType `json:"ticket_types" jsmap:"ticket_types"`
}

type TicketType struct {
    TicketID  string `json:"id_ticket" jsmap:"id"`
    Type      string `json:"ticket_type" jsmap:"type"`
    Price     string `json:"price" jsmap:"price"`
    Booked    string `json:"booked" jsmap:"booked"`
    Sold      string `json:"sold" jsmap:"sold"`
    Available string `json:"available" jsmap:"available"`
}

type EventService struct {
    Response struct {
        Result interface{}
    }
}

func JsMap(in interface{}) (map[string]interface{}, error) {
    out := make(map[string]interface{})

    v := reflect.ValueOf(in)

    if v.Kind() == reflect.Ptr {
        v = v.Elem()
    }

    if v.Kind() != reflect.Struct {
        return nil, fmt.Errorf("JsMap only accepts structs; got %T", v)
    }

    for i := 0; i < v.NumField(); i++ {
        // gets us a StructField
        fi := v.Type().Field(i)
        valueField := v.Field(i)

        tagv := fi.Tag.Get("jsmap")
        //tagv = strings.ToLower(fi.Name)

        if tagv != "" {
            // set key of map to value in struct field
            if valueField.Kind() == reflect.Slice {
                var xx []interface{}

                s := reflect.ValueOf(valueField.Interface())

                for i := 0; i < s.Len(); i++ {
                    fmt.Println(s.Index(i).Interface())
                    x, err := JsMap(s.Index(i).Interface())
                    if err != nil {
                        fmt.Println("error bray")
                    }
                    xx = append(xx, x)
                }

                fmt.Println(xx)

                out[tagv] = xx
            } else if valueField.Kind() == reflect.Struct {
                x, err := JsMap(valueField)
                if err != nil {
                    fmt.Println("error bray")
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

    x, err := JsMap(ev)
    fmt.Println(x)

    self.Response.Result = x

    return nil
}
