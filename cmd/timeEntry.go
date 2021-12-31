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
	timeEntryCmd.Flags().StringP("format", "f", "table", "output format defaults to table (table, json, csv)")
}

func listTimeEntries(cmd *cobra.Command, args []string) error {
	h := initHakunaClient()

	since, err := cmd.LocalFlags().GetString("since")
	if err != nil {
		return err
	}

	until, err := cmd.LocalFlags().GetString("until")
	if err != nil {
		return err
	}

	format, err := cmd.LocalFlags().GetString("format")
	if err != nil {
		return err
	}

	startDate, err := lib.ParseDate(since)
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

	data := make([][]string, len(timeEntries))
	for i, row := range data {
		entry := timeEntries[i]
		data[i] = append(row,
			fmt.Sprint(entry.ID),
			fmt.Sprint(entry.Date),
			fmt.Sprint(entry.StartTime),
			fmt.Sprint(entry.EndTime),
			fmt.Sprint(entry.Duration),
			fmt.Sprint(entry.Note),
			fmt.Sprint(entry.Task.Name),
		)
	}

	return lib.RenderData(
		format,
		[]string{
			"ID",
			"Date",
			"Start time",
			"End Time",
			"Duration",
			"Note",
			"Task name",
		},
		data,
		timeEntries,
	)
}
