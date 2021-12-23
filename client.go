package gohakuna

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Hakuna struct {
	SubDomain string
	Token     string
	Client    http.Client
}

type Request struct {
	Method   string
	Endpoint string
	Body     io.Reader
}

type Response struct {
	Status int
	Body   []byte
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

func (h Hakuna) StartTimer() Timer {
	req := Request{Method: "POST", Endpoint: "/timer"}
	res := h.request(req)

	var timer Timer
	if err := json.Unmarshal(res.Body, &timer); err != nil {
		log.Fatal("Error decoding pong response")
	}

	return timer
}

func (h Hakuna) StopTimer() TimeEntry {
	req := Request{Method: "GET", Endpoint: "/ping"}
	res := h.request(req)

	var timeEntry TimeEntry
	if err := json.Unmarshal(res.Body, &timeEntry); err != nil {
		log.Fatal("Error decoding pong response")
	}

	return timeEntry
}

func (h Hakuna) request(req Request) Response {
	url := "https://" + h.SubDomain + ".hakuna.ch/api/v1" + req.Endpoint

	rq, err := http.NewRequest(req.Method, url, req.Body)
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

	return Response{Body: body, Status: resp.StatusCode}
}

func decodeBody(body []byte, pointer interface{}) {
	err := json.Unmarshal(body, &pointer)
	if err != nil {
		log.Fatal("Error decoding body")
	}
}
