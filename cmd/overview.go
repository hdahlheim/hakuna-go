/*
Copyright © 2021 Henning Dahlheim <hactar@cyberkraft.ch>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// overviewCmd represents the overview command
var overviewCmd = &cobra.Command{
	Use:   "overview",
	Short: "A brief description of your command",
	RunE:  getOverview,
}

func init() {
	rootCmd.AddCommand(overviewCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// overviewCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// overviewCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getOverview(cmd *cobra.Command, args []string) error {
	h := getHakunaClient()

	overview, err := h.GetOverview()
	if err != nil {
		return err
	}

	fmt.Printf("----------Overview----------\nOvertime:\t\t%v\nVacation:\n- Days redeemed:\t%v\n- Days remaining:\t%v\n",
		overview.Overtime,
		overview.Vacation.RedeemedDays,
		overview.Vacation.RemainingDays,
	)
	return nil
}
