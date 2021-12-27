package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	hakuna "github.com/hdahlheim/hakuna-go"
)

func main() {

	pingAPI := flag.Bool("ping", false, "Ping API")
	startTimer := flag.Bool("startTimer", false, "Start the hakuna timer")
	stopTimer := flag.Bool("stopTimer", false, "Stop the hakuna timer")
	toggleTimer := flag.Bool("toggleTimer", false, "Toggle the hakuna timer")
	tokenFlag := flag.String("token", "", "Toggle the hakuna timer")

	flag.Parse()

	subdomain := os.Getenv("HAKUNA_CLI_SUBDOMAIN")
	tokenEnv := os.Getenv("HAKUNA_CLI_API_TOKEN")

	if subdomain == "" {
		fmt.Fprintf(os.Stderr, "Subdomain needed\nSet the Subdomain using the env var HAKUNA_CLI_SUBDOMAIN\n")
		os.Exit(1)
	}

	var token string

	if *tokenFlag != "" {
		token = *tokenFlag
	} else if tokenEnv != "" {
		token = tokenEnv
	} else {
		fmt.Fprintf(os.Stderr, "API token needed\nSet the API token using the -token flag or the env var HAKUNA_CLI_API_TOKEN\n")
		os.Exit(1)
	}

	client := http.Client{Timeout: time.Second * 2}
	h := hakuna.New(subdomain, token, client)

	if *pingAPI {
		fmt.Fprintf(os.Stderr, "Ping API\n")
		pong := h.Ping()
		fmt.Printf("%v\n", pong.Pong)
	} else if *startTimer {
		fmt.Fprintf(os.Stderr, "Timer starting timer\n")
		data := hakuna.StartTimerReq{TaskId: 2}

		t, err := h.StartTimer(data)

		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error()+"\n")
			os.Exit(1)
		}

		fmt.Printf("Timer started at %v\n", t.StartTime)
	} else if *stopTimer {
		fmt.Fprintf(os.Stderr, "Stopping timer\n")

		entry, err := h.StopTimer()
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error()+"\n")
			os.Exit(1)
		}

		fmt.Printf("Time stopped\nDuration:\t %v\nStarted:\t%v\nStopped:\t%v\n", entry.Duration, entry.StartTime, entry.EndTime)
	} else if *toggleTimer {
		fmt.Printf("!Todo\n")
	}
}
