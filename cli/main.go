package main

import (
	"fmt"
)

func main() {
	fmt.Println("deprecated")
	// pingAPI := flag.Bool("ping", false, "Ping API")
	// startTimer := flag.Bool("startTimer", false, "Start the hakuna timer")
	// stopTimer := flag.Bool("stopTimer", false, "Stop the hakuna timer")
	// toggleTimer := flag.Bool("toggleTimer", false, "Toggle the hakuna timer")
	// tokenFlag := flag.String("token", "", "Toggle the hakuna timer")

	// flag.Parse()

	// subdomain := os.Getenv("HAKUNA_CLI_SUBDOMAIN")
	// tokenEnv := os.Getenv("HAKUNA_CLI_API_TOKEN")

	// if subdomain == "" {
	// 	fmt.Fprintf(os.Stderr, "Subdomain needed\nSet the Subdomain using the env var HAKUNA_CLI_SUBDOMAIN\n")
	// 	os.Exit(1)
	// }

	// var token string
	// switch {
	// case *tokenFlag != "":
	// 	token = *tokenFlag
	// case tokenEnv != "":
	// 	token = tokenEnv
	// default:
	// 	fmt.Fprintf(os.Stderr, "API token needed\nSet the API token using the -token flag or the env var HAKUNA_CLI_API_TOKEN\n")
	// 	os.Exit(1)
	// }

	// client := http.Client{Timeout: time.Second * 2}
	// h, err := hakuna.New(subdomain, token, client)
	// if err != nil {
	// 	fmt.Fprint(os.Stderr, err.Error())
	// }

	// switch {
	// case *pingAPI:
	// 	fmt.Fprintf(os.Stderr, "Ping API\n")

	// 	pong, err := h.Ping()
	// 	if err != nil {
	// 		fmt.Fprint(os.Stderr, err.Error())
	// 		os.Exit(1)
	// 	}

	// 	fmt.Printf("%v\n", pong.Pong)
	// case *startTimer:
	// 	fmt.Fprintf(os.Stderr, "Timer starting timer\n")
	// 	data := hakuna.StartTimerReq{TaskId: 2}

	// 	t, err := h.StartTimer(data)
	// 	if err != nil {
	// 		fmt.Fprintf(os.Stderr, err.Error()+"\n")
	// 		os.Exit(1)
	// 	}

	// 	fmt.Printf("Timer started at %v\n", t.StartTime)
	// case *stopTimer:
	// 	fmt.Fprintf(os.Stderr, "Stopping timer\n")

	// 	entry, err := h.StopTimer()
	// 	if err != nil {
	// 		fmt.Fprintf(os.Stderr, err.Error()+"\n")
	// 		os.Exit(1)
	// 	}

	// 	fmt.Printf("Time stopped\nDuration:\t %v\nStarted:\t%v\nStopped:\t%v\n", entry.Duration, entry.StartTime, entry.EndTime)
	// case *toggleTimer:
	// 	fmt.Printf("!Todo\n")
	// }
}
