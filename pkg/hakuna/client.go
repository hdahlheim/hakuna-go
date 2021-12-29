/*
Copyright © 2021 Henning Dahlheim <hactar@cyberkraft.ch>

*/
package hakuna

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
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

func New(subDomain string, token string, client http.Client) (*Hakuna, error) {
	switch {
	case subDomain == "":
		return nil, errors.New("subdomain can not be empty string")
	case token == "":
		return nil, errors.New("api token can not be empty string")
	default:
		return &Hakuna{SubDomain: subDomain, Token: token, Client: client}, nil
	}
}

func NewStartTimerReq(taskId int, startTime time.Time, note string, projectId int) (*StartTimerReq, error) {
	req := &StartTimerReq{
		TaskId:    taskId,
		StartTime: startTime.Format("15:04"),
	}

	switch {
	case req.TaskId == 0:
		return nil, errors.New("taskId must be greater than 0")
	case projectId != 0:
		req.ProjectId = projectId
	case note != "":
		req.Note = note
	}

	return req, nil
}

func NewStopTimerReq(time time.Time) (*StopTimerReq, error) {
	req := &StopTimerReq{
		EndTime: time.Format("15:04"),
	}

	return req, nil
}

func (h Hakuna) Ping() (Pong, error) {
	req := request{
		Method:   "GET",
		Endpoint: "/ping",
	}
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

func (h Hakuna) GetTimer() (Timer, error) {
	req := request{
		Method:   "GET",
		Endpoint: "/timer",
	}
	res, err := h.request(req)
	if err != nil {
		return Timer{}, err
	}

	if err := getResponeError(res); err != nil {
		return Timer{}, err
	}

	var timer Timer
	if err := json.Unmarshal(res.Body, &timer); err != nil {
		return Timer{}, err
	}

	return timer, nil
}

func (h Hakuna) StartTimer(data *StartTimerReq) (Timer, error) {

	reqBody, err := json.Marshal(&data)
	if err != nil {
		return Timer{}, err
	}

	req := request{
		Method:   "POST",
		Endpoint: "/timer",
		Body:     reqBody,
	}

	res, err := h.request(req)
	if err != nil {
		return Timer{}, err
	}

	if err := getResponeError(res); err != nil {
		return Timer{}, err
	}

	var timer Timer
	if err := json.Unmarshal(res.Body, &timer); err != nil {
		return timer, err
	}

	return timer, nil
}

func (h Hakuna) StopTimer(data *StopTimerReq) (TimeEntry, error) {
	reqBody, err := json.Marshal(&data)
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

func (h Hakuna) GetTimeEntries(start time.Time, end time.Time) ([]TimeEntry, error) {
	req := request{
		Method:   "GET",
		Endpoint: "/time_entries?start_date=" + start.Format("2006-01-02") + "&" + "end_date=" + end.Format("2006-01-02"),
	}
	res, err := h.request(req)
	if err != nil {
		return nil, err
	}

	if err := getResponeError(res); err != nil {
		return nil, err
	}

	var timeEntries []TimeEntry
	if err := json.Unmarshal(res.Body, &timeEntries); err != nil {
		return timeEntries, errors.New("error decoding response")
	}

	return timeEntries, nil
}

func (h Hakuna) GetOverview() (Overview, error) {
	req := request{
		Method:   "GET",
		Endpoint: "/overview",
	}
	res, err := h.request(req)
	if err != nil {
		return Overview{}, err
	}

	if err := getResponeError(res); err != nil {
		return Overview{}, err
	}

	var overview Overview
	if err := json.Unmarshal(res.Body, &overview); err != nil {
		return overview, errors.New("error decoding response")
	}

	return overview, nil
}

func (h Hakuna) request(req request) (response, error) {
	url := "https://" + h.SubDomain + ".hakuna.ch/api/v1" + req.Endpoint

	rq, err := http.NewRequest(req.Method, url, bytes.NewBuffer(req.Body))
	if err != nil {
		return response{}, err
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
