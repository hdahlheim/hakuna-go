/*
Copyright © 2021 Henning Dahlheim <hactar@cyberkraft.ch>

*/
package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/hdahlheim/hakuna-go/internal/lib"
	"github.com/spf13/cobra"
)

// timeEntryCmd represents the timeEntry command
var timeEntryCmd = &cobra.Command{
	Use:   "time-entry",
	Short: "Functions related to time entries.",
}

var listTimeEntryCmd = &cobra.Command{
	Use:   "list",
	Short: "List your time entries.",
	RunE:  listTimeEntries,
}

func init() {
	rootCmd.AddCommand(timeEntryCmd)
	timeEntryCmd.AddCommand(listTimeEntryCmd)
	listTimeEntryCmd.Flags().StringP("since", "S", "today", "--since=\"today\" | --since=\"2021-12-01\"")
	listTimeEntryCmd.Flags().StringP("until", "U", "today", "--until=\"yesterday\" | --until=\"2021-12-31\"")
}

func listTimeEntries(cmd *cobra.Command, args []string) error {
	fmt.Fprintf(os.Stderr, "Loading time entries...\n")
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
		return errors.New("end date must be after or equal to start date")
	}

	timeEntries, err := h.GetTimeEntries(startDate.Format("2006-01-02"), endDate.Format("2006-01-02"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Error] "+err.Error()+"\n")
	}

	for i, entry := range timeEntries {
		fmt.Printf("------Time Entry-----------\nId:\t\t%v\nDate:\t\t%v\nStart:\t\t%v\nEnd:\t\t%v\nDuration:\t%v\n", entry.ID, entry.Date, entry.StartTime, entry.EndTime, entry.Duration)
		if i == len(timeEntries)-1 {
			fmt.Printf("----------------------------\n")
		}
	}
	return nil
}
