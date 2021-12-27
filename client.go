package hakuna

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Hakuna struct {
	SubDomain string
	Token     string
	Client    http.Client
}

type Request struct {
	Method   string
	Endpoint string
	Body     []byte
}

type Response struct {
	StatusCode int
	Body       []byte
}

func New(subDomain string, token string, client http.Client) Hakuna {
	return Hakuna{SubDomain: subDomain, Token: token, Client: client}
}

func (h Hakuna) Ping() Pong {
	req := Request{Method: "GET", Endpoint: "/ping"}
	res := h.request(req)

	var pong Pong
	if err := json.Unmarshal(res.Body, &pong); err != nil {
		log.Fatal("Error decoding pong response")
	}

	return pong
}

func (h Hakuna) StartTimer(data StartTimerReq) (Timer, error) {
	reqBody, err := json.Marshal(&data)
	if err != nil {
		log.Fatal("Error creating request body")
	}

	req := Request{Method: "POST", Endpoint: "/timer", Body: reqBody}
	res := h.request(req)

	var timer Timer

	if err := getResponeError(res); err != nil {
		return timer, err
	}

	if err := json.Unmarshal(res.Body, &timer); err != nil {
		return timer, errors.New("Error decoding response")
	}

	return timer, nil
}

func (h Hakuna) StopTimer() (TimeEntry, error) {
	now := time.Now()
	timeString := fmt.Sprintf("%d:%d", now.Hour(), now.Minute())

	reqData := StopTimerReq{EndTime: timeString}

	reqBody, err := json.Marshal(&reqData)
	if err != nil {
		log.Fatal("Error creating request body")
	}

	req := Request{Method: "PUT", Endpoint: "/timer", Body: reqBody}
	res := h.request(req)

	var timeEntry TimeEntry

	if err := getResponeError(res); err != nil {
		return timeEntry, err
	}

	if err := json.Unmarshal(res.Body, &timeEntry); err != nil {
		return timeEntry, errors.New("Error decoding response")
	}

	return timeEntry, nil
}

func (h Hakuna) request(req Request) Response {
	url := "https://" + h.SubDomain + ".hakuna.ch/api/v1" + req.Endpoint

	rq, err := http.NewRequest(req.Method, url, bytes.NewBuffer(req.Body))
	if err != nil {
		log.Fatal("Error doing Request", err)
	}

	rq.Header.Set("Content-Type", "application/json; charset=utf-8")
	rq.Header.Set("X-Auth-Token", h.Token)

	resp, err := h.Client.Do(rq)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading body. ", err)
	}

	return Response{Body: body, StatusCode: resp.StatusCode}
}

func getResponeError(res Response) error {
	var apiError ResponeError

	if res.StatusCode >= 300 {
		if err := json.Unmarshal(res.Body, &apiError); err != nil {
			return errors.New("Error decoding response")
		}
		return errors.New(apiError.Message)
	}
	return nil
}
