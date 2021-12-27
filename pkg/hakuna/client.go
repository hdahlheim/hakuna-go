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

type request struct {
	Method   string
	Endpoint string
	Body     []byte
}

type response struct {
	StatusCode int
	Body       []byte
}

func New(subDomain string, token string, client http.Client) Hakuna {
	return Hakuna{SubDomain: subDomain, Token: token, Client: client}
}

func (h Hakuna) Ping() (Pong, error) {
	req := request{Method: "GET", Endpoint: "/ping"}
	res, err := h.request(req)
	if err != nil {
		return Pong{}, err
	}

	var pong Pong
	if err := json.Unmarshal(res.Body, &pong); err != nil {
		return Pong{}, err
	}

	return pong, nil
}

func (h Hakuna) StartTimer(data StartTimerReq) (Timer, error) {
	var timer Timer

	reqBody, err := json.Marshal(&data)
	if err != nil {
		return Timer{}, err
	}

	req := request{Method: "POST", Endpoint: "/timer", Body: reqBody}
	res, err := h.request(req)
	if err != nil {
		return timer, err
	}

	if err := getResponeError(res); err != nil {
		return timer, err
	}

	if err := json.Unmarshal(res.Body, &timer); err != nil {
		return timer, err
	}

	return timer, nil
}

func (h Hakuna) StopTimer() (TimeEntry, error) {
	now := time.Now()
	timeString := fmt.Sprintf("%d:%d", now.Hour(), now.Minute())

	reqData := StopTimerReq{EndTime: timeString}

	reqBody, err := json.Marshal(&reqData)
	if err != nil {
		return TimeEntry{}, err
	}

	req := request{Method: "PUT", Endpoint: "/timer", Body: reqBody}
	res, err := h.request(req)
	if err != nil {
		return TimeEntry{}, err
	}

	if err := getResponeError(res); err != nil {
		return TimeEntry{}, err
	}

	var timeEntry TimeEntry
	if err := json.Unmarshal(res.Body, &timeEntry); err != nil {
		return timeEntry, errors.New("error decoding response")
	}

	return timeEntry, nil
}

func (h Hakuna) request(req request) (response, error) {
	url := "https://" + h.SubDomain + ".hakuna.ch/api/v1" + req.Endpoint

	rq, err := http.NewRequest(req.Method, url, bytes.NewBuffer(req.Body))
	if err != nil {
		log.Fatal("Error doing Request", err)
	}

	rq.Header.Set("Content-Type", "application/json; charset=utf-8")
	rq.Header.Set("X-Auth-Token", h.Token)

	resp, err := h.Client.Do(rq)
	if err != nil {
		return response{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return response{}, err
	}

	return response{Body: body, StatusCode: resp.StatusCode}, nil
}

func getResponeError(res response) error {
	var apiError ResponeError

	if res.StatusCode >= 300 {
		if err := json.Unmarshal(res.Body, &apiError); err != nil {
			return errors.New("error decoding response")
		}
		return errors.New(apiError.Message)
	}
	return nil
}
