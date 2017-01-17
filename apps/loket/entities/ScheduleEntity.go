package entity

type ScheduleData struct {
    ScheduleID     string       `json:"id_schedule" jsmap:"schedule_id"`
    GroupName      string       `json:"group_name" jsmap:"name"`
    MaxTransaction string       `json:"max_transaction" jsmap:"max_transaction"`
    TicketTypes    []TicketType `json:"ticket_types" jsmap:"ticket_types"`
}
