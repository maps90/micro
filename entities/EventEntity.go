package entity

type Events struct {
    Status  string      `json:"status" jsmap:"status"`
    Data    []EventData `json:"data" jsmap:"data"`
    Message string      `json:"message" jsmap:"message"`
    Code    int         `json:"code" jsmap:"code"`
}

type EventData struct {
    EventID   string          `json:"id_event" jsmap:"event_id"`
    Name      string          `json:"event_name" jsmap:"name"`
    Schedules []EventSchedule `json:"schedules" jsmap:"schedules"`
}

type EventSchedule struct {
    ScheduleID   string       `json:"id_schedule" jsmap:"schedule_id"`
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
