/*
Copyright © 2021 Henning Dahlheim <hactar@cyberkraft.ch>

*/
package cmd

import (
	"errors"
	"fmt"

	"github.com/hdahlheim/hakuna-go/internal/lib"
	"github.com/spf13/cobra"
)

// timeEntryCmd represents the timeEntry command
var timeEntryCmd = &cobra.Command{
	Use:     "entry",
	Short:   "Functions related to time entries.",
	Aliases: []string{"entry list"},
	RunE:    listTimeEntries,
}

func init() {
	rootCmd.AddCommand(timeEntryCmd)
	timeEntryCmd.Flags().StringP("since", "S", "today", "--since=\"today\" | --since=\"2021-12-01\"")
	timeEntryCmd.Flags().StringP("until", "U", "today", "--until=\"yesterday\" | --until=\"2021-12-31\"")
}

const timeEntryTpl = `--------Time Entry--------
Id:          	    %v
Date:           %v
Start time:          %v
End time:            %v
Duration:             %v
Task:               %v
`

func listTimeEntries(cmd *cobra.Command, args []string) error {
	h := getHakunaClient()

	since, err := cmd.LocalFlags().GetString("since")
	if err != nil {
		return err
	}

	startDate, err := lib.ParseDate(since)
	if err != nil {
		return err
	}

	until, err := cmd.LocalFlags().GetString("until")
	if err != nil {
		return err
	}

	endDate, err := lib.ParseDate(until)
	if err != nil {
		return err
	}

	if startDate.Unix() > endDate.Unix() {
		return errors.New("end date must be after or equal to the start date")
	}

	timeEntries, err := h.GetTimeEntries(startDate, endDate)
	if err != nil {
		return err
	}

	for i, entry := range timeEntries {
		fmt.Printf(timeEntryTpl,
			entry.ID,
			entry.Date,
			entry.StartTime,
			entry.EndTime,
			entry.Duration,
			entry.Task.Name,
		)
		if i == len(timeEntries)-1 {
			fmt.Printf("--------------------------\n")
		}
	}
	return nil
}
