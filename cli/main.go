package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/hdahlheim/gohakuna"
)

func main() {

	pingAPI := flag.Bool("ping", false, "Ping API")
	startTimer := flag.Bool("startTimer", false, "Start the hakuna timer")
	stopTimer := flag.Bool("stopTimer", false, "Stop the hakuna timer")
	toggleTimer := flag.Bool("toggleTimer", false, "Toggle the hakuna timer")
	token := flag.String("token", "", "Toggle the hakuna timer")

	flag.Parse()

	if *token == "" {
		fmt.Fprintf(os.Stderr, "API Token needed\n")
		os.Exit(1)
	}

	client := http.Client{Timeout: time.Second * 2}
	h := gohakuna.New("cyon", *token, client)

	if *pingAPI {
		fmt.Printf("Ping API \n")
		pong := h.Ping()
		fmt.Printf("%v \n", pong.Pong)
	} else if *startTimer {
		fmt.Printf("Timer started \n")
		h.StartTimer()
	} else if *stopTimer {
		fmt.Printf("Timer stopped \n")
		h.StopTimer()
	} else if *toggleTimer {
		fmt.Printf("Timer toggeld \n")
	}
}
