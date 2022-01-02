/*
Copyright © 2021 Henning Dahlheim hactar@cyberkraft.ch

This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at https://mozilla.org/MPL/2.0/.

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

// timerCmd represents the timer start command
var startTimerCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a new timer",
	RunE:  startTimer,
}

// timerCmd represents the timer stop command
var stopTimerCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop the current timer",
	RunE:  stopTimer,
}

func init() {
	rootCmd.AddCommand(timerCmd)
	timerCmd.AddCommand(startTimerCmd, stopTimerCmd)

	// start command flags
	startTimerCmd.Flags().IntP("taskId", "t", 0, "--taskId=1")
	startTimerCmd.Flags().StringP("startTime", "s", "now", "--startTime=\"12:30\"")
	startTimerCmd.Flags().StringP("note", "n", "", "--note=\"Honest work!\"")

	// end command flags
	stopTimerCmd.Flags().StringP("endTime", "e", "now", "--endTime=\"12:30\"")
}

func getTimer(cmd *cobra.Command, args []string) error {
	h := initHakunaClient()

	timer, err := h.GetTimer()
	if err != nil {
		return err
	}

	if timer.StartTime == "" && timer.DurationInSeconds == 0.0 {
		fmt.Printf("No timer running\n")
		return nil
	}

	fmt.Printf("Timer running since %v %v\nDuration:\t%v\nTask:\t\t%v\n",
		timer.Date,
		timer.StartTime,
		timer.Duration,
		timer.Task.Name,
	)
	return nil
}

func startTimer(cmd *cobra.Command, args []string) error {
	taskIdFlag, err := cmd.LocalFlags().GetInt("taskId")
	if err != nil {
		return err
	}

	taskId := lib.Ternary(taskIdFlag != 0, taskIdFlag, cliConfig.Default.TaskId).(int)

	if taskId <= 0 {
		return fmt.Errorf("taskId is required")
	}

	var startTime time.Time
	startTimeFlag, err := cmd.LocalFlags().GetString("startTime")
	if err != nil {
		return err
	} else {
		pTime, err := lib.ParseTime(startTimeFlag)
		if err != nil {
			return err
		}
		startTime = pTime
	}

	h := initHakunaClient()

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
	h := initHakunaClient()

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

	fmt.Printf("Timer stopped at %v. The timer was running for %v\n",
		timer.EndTime,
		timer.Duration,
	)

	return nil
}
