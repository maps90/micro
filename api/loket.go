package api

import (
	"encoding/json"
	"fmt"

	gr "github.com/parnurzeal/gorequest"
)

var conf map[string]string
var baseUrl string

type Loket struct {
	BaseUrl      string
	UserName     string
	Password     string
	ApiKey       string
	Token        string
	Response     LoketResponse
	Body         string
	Error        error
	TokenExpired bool
}

type LoketResponse struct {
	Code    uint16      `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Status  string      `json:"status"`
}

func getConfig(key string) string {
	if _, ok := conf[key]; ok {
		return conf[key]
	}
	return ""
}

func (l *Loket) getAuth() *Loket {
	if !l.TokenExpired {
		return l
	}
	if len(l.UserName) == 0 || len(l.Password) == 0 || len(l.ApiKey) == 0 {
		return l
	}
	var errs []error
	body := fmt.Sprintf(`{"username": "%s","password": "%s","APIKEY": "%s"}`, l.UserName, l.Password, l.ApiKey)
	_, l.Body, errs = gr.New().
		Post(SetUrl("/v3/login")).
		Type("form").
		Send(body).
		End()

	for _, err := range errs {
		l.Error = err
	}

	if err := json.Unmarshal([]byte(l.Body), &l.Response); err != nil {
		l.Error = err
	}
	l.SetToken()
	return l
}

func NewLoketApi(url, username, password, key, clientKey string) (*Loket, error) {
	baseUrl = url
	l := &Loket{
		UserName:     username,
		Password:     password,
		ApiKey:       key,
		Token:        clientKey,
		TokenExpired: true,
	}
	return l, nil
}

func GetResources() map[string]string {
	r := map[string]string{
		"get_event_list":         SetUrl("v3/event"),
		"get_ticket_groups":      SetUrl("v3/schedule/:scheduleID"),
		"get_ticket_by_schedule": SetUrl("v3/tickets/:scheduleID"),
	}
	return r
}

func SetUrl(url string) string {
	t := fmt.Sprintf("%s%s", baseUrl, url)
	return t
}

func (l *Loket) SetToken() *Loket {
	resp := struct {
		Status string `json:"status"`
		Data   *struct {
			Token string `json:"token"`
		} `json:"data"`
		Message string `json:"message"`
	}{"", nil, ""}
	byt := []byte(l.Body)

	if err := json.Unmarshal(byt, &resp); err != nil {
		return l
	}

	l.Token = resp.Data.Token
	l.TokenExpired = false
	return l
}

func (l *Loket) SetStruct(v interface{}) *Loket {
	err := json.Unmarshal([]byte(l.Body), &v)
	if err != nil {
		l.Error = err
		return l
	}
	return l
}

func (l *Loket) Post(url, t, body string) *Loket {
	var errs []error
	_, l.Body, errs = gr.New().
		Post(SetUrl(url)).
		Set("clientKey", l.Token).
		Type(t).
		Send(body).
		End()

	for _, err := range errs {
		l.Error = err
	}

	if err := json.Unmarshal([]byte(l.Body), &l.Response); err != nil {
		l.Error = err
	}

	return l
}

func (l *Loket) Get(url string) *Loket {
	var errs []error
	_, l.Body, errs = gr.New().
		Set("clientKey", l.Token).
		Get(SetUrl(url)).
		Set("token", l.Token).
		End()

	for _, err := range errs {
		l.Error = err
	}

	if err := json.Unmarshal([]byte(l.Body), &l.Response); err != nil {
		l.Error = err
	}

	return l
}
