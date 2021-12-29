/*
Copyright © 2021 Henning Dahlheim hactar@cyberkraft.ch

*/
package cmd

import (
	"fmt"
	"time"

	"github.com/hdahlheim/hakuna-go/internal/lib"
	"github.com/hdahlheim/hakuna-go/pkg/hakuna"
	"github.com/spf13/cobra"
)

// timerCmd represents the timer command
var timerCmd = &cobra.Command{
	Use:   "timer",
	Short: "Interact with the hakuna timer resource",
	RunE:  getTimer,
}

var startTimerCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a new timer",
	RunE:  startTimer,
}

var stopTimerCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop the current timer",
	RunE:  stopTimer,
}

func init() {
	rootCmd.AddCommand(timerCmd)
	timerCmd.AddCommand(startTimerCmd, stopTimerCmd)

	startTimerCmd.Flags().IntP("taskId", "T", 0, "--taskId=1")
	startTimerCmd.Flags().StringP("startTime", "S", "now", "--startTime=\"12:30\"")
	startTimerCmd.Flags().StringP("note", "N", "", "--note=\"Honest work!\"")

	stopTimerCmd.Flags().StringP("endTime", "E", "now", "--endTime=\"12:30\"")
}

func getTimer(cmd *cobra.Command, args []string) error {
	h := getHakunaClient()

	timer, err := h.GetTimer()
	if err != nil {
		return err
	}

	if timer.StartTime == "" && timer.DurationInSeconds == 0.0 {
		fmt.Printf("No timer running\n")
		return nil
	}

	fmt.Printf("Timer running since %v %v\nDuration:\t%v\nTask:\t\t%v\n", timer.Date, timer.StartTime, timer.Duration, timer.Task.Name)
	return nil
}

func startTimer(cmd *cobra.Command, args []string) error {
	taskId, err := cmd.LocalFlags().GetInt("taskId")
	if err != nil {
		return err
	}

	if taskId <= 0 {
		return fmt.Errorf("taskId is required")
	}

	var startTime time.Time
	startTimeStr, err := cmd.LocalFlags().GetString("startTime")
	if err != nil {
		return err
	} else {
		pTime, err := lib.ParseTime(startTimeStr)
		if err != nil {
			return err
		}
		startTime = pTime
	}

	h := getHakunaClient()

	req, err := hakuna.NewStartTimerReq(taskId, startTime, "", 0)
	if err != nil {
		return err
	}

	timer, err := h.StartTimer(req)
	if err != nil {
		return err
	}

	fmt.Printf("Timer started at %v\n", timer.StartTime)
	return nil
}

func stopTimer(cmd *cobra.Command, args []string) error {
	h := getHakunaClient()

	endTimeFlag, err := cmd.LocalFlags().GetString("endTime")
	if err != nil {
		return err
	}

	var endTime time.Time
	if endTimeFlag != "" {
		pTime, err := lib.ParseTime(endTimeFlag)
		if err != nil {
			return err
		}
		endTime = pTime
	} else {
		endTime = time.Now()
	}

	data, err := hakuna.NewStopTimerReq(endTime)
	if err != nil {
		return err
	}

	timer, err := h.StopTimer(data)
	if err != nil {
		return err
	}

	fmt.Printf("Timer stopped at %v. The timer was running for %v\n", timer.EndTime, timer.Duration)
	return nil
}
