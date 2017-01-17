package entity

type TicketData struct {
    GroupID        string       `json:"id_group" jsmap:"group_id"`
    MaxTransaction string       `json:"max_transaction" jsmap:"max_transaction"`
    ScheduleID     string       `json:"id_schedule" jsmap:"schedule_id"`
    TicketTypes    []TicketType `json:"ticket_types" jsmap:"ticket_types"`
}

type TicketType struct {
    TicketID  string `json:"id_ticket" jsmap:"ticket_id"`
    Type      string `json:"ticket_type" jsmap:"type"`
    Price     string `json:"price" jsmap:"price"`
    Booked    string `json:"booked" jsmap:"booked"`
    Sold      string `json:"sold" jsmap:"sold"`
    Available string `json:"available" jsmap:"available"`
    Quantity  string `json:"quantity" jsmap:"quantity"`
}
