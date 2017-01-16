package loket

import (
	"encoding/json"
	"fmt"
	"net/http"

	gr "github.com/parnurzeal/gorequest"
	c "github.com/spf13/viper"
)

type Loket struct {
	UserName     string
	Password     string
	ApiKey       string
	Token        string
	Response     *http.Response
	Body         string
	Errors       []error
	TokenExpired bool
}

var LOKET_URI = ""

func (l *Loket) GetAuth() *Loket {
	if l.TokenExpired {
		return l
	}
	if len(l.UserName) == 0 || len(l.Password) == 0 || len(l.ApiKey) == 0 {
		return l
	}
	body := fmt.Sprintf(`{"username": "%s","password": "%s","APIKEY": "%s"}`, l.UserName, l.Password, l.ApiKey)
	l.Post("v3", "login", body)
	l.SetToken()
	return l
}

func New() *Loket {
	setBaseUrl(c.GetString("loket.url"))
	l := &Loket{
		UserName: c.GetString("loket.username"),
		Password: c.GetString("loket.password"),
		ApiKey:   c.GetString("loket.key"),
		Token:    "",
	}
	return l
}

func setBaseUrl(c string) {
	LOKET_URI = c
}

func SetUrl(version, url string) string {
	t := fmt.Sprintf("%s/%s/%s", LOKET_URI, version, url)
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
	return l
}

func (l *Loket) SetStruct(v interface{}) *Loket {
	err := json.Unmarshal([]byte(l.Body), &v)
	if err != nil {
		l.Errors = append(l.Errors, err)
		return l
	}
	return l
}

func (l *Loket) Post(vr, url, body string) *Loket {
	l.Response, l.Body, l.Errors = gr.New().
		Post(SetUrl(vr, url)).
		Type("form").
		Send(body).
		End()
	return l
}

func (l *Loket) Get(vr, url, body string) *Loket {
	l.Response, l.Body, l.Errors = gr.New().
		Get(SetUrl(vr, url)).
		Send(body).
		End()
	return l
}
