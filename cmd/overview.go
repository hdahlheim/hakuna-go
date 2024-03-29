/*
Copyright © 2021 Henning Dahlheim <hactar@cyberkraft.ch>

This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at https://mozilla.org/MPL/2.0/.

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// overviewCmd represents the overview command
var overviewCmd = &cobra.Command{
	Use:   "overview",
	Short: "Show your overtime and vactation days",
	RunE:  getOverview,
}

func init() {
	rootCmd.AddCommand(overviewCmd)
}

func getOverview(cmd *cobra.Command, args []string) error {
	h := initHakunaClient()

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
