/*
Copyright © 2021 Henning Dahlheim <hactar@cyberkraft.ch>

This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at https://mozilla.org/MPL/2.0/.

*/
package cmd

import (
	"fmt"
	"time"

	"github.com/hdahlheim/hakuna-go/internal/lib"
	"github.com/spf13/cobra"
)

// absencesCmd represents the absences command
var absenceCmd = &cobra.Command{
	Use:     "absence",
	Aliases: []string{"absences", "absences list", "absences list"},
	Short:   "List absences for the year",
	RunE: func(cmd *cobra.Command, args []string) error {
		h := initHakunaClient()

		yearFlag, err := cmd.LocalFlags().GetString("year")
		if err != nil {
			return err
		}

		format, err := cmd.LocalFlags().GetString("format")
		if err != nil {
			return err
		}

		year := lib.Ternary(yearFlag != "", yearFlag, time.Now().Format("2006")).(string)

		absences, err := h.GetAbsences(year)
		if err != nil {
			return err
		}

		data := make([][]string, len(absences))
		for i, row := range data {
			absence := absences[i]
			data[i] = append(row,
				fmt.Sprint(absence.ID),
				fmt.Sprint(absence.StartDate),
				fmt.Sprint(absence.EndDate),
				fmt.Sprint(absence.AbsenceType.Name),
				fmt.Sprint(absence.FirstHalfDay),
				fmt.Sprint(absence.SecondHalfDay),
				fmt.Sprint(absence.IsRecurring),
				fmt.Sprint(absence.WeeklyRepeatInterval),
				fmt.Sprint(absence.AbsenceType.IsVacation),
				fmt.Sprint(absence.AbsenceType.GrantsWorkTime),
			)
		}

		return lib.RenderData(format,
			[]string{
				"Id",
				"Start Date",
				"End Date",
				"Absence type",
				"First half",
				"Second half",
				"Is recurring",
				"Weekly repeat interval",
				"Is Vacation",
				"Grants work time",
			},
			data,
			absences,
		)
	},
}

func init() {
	rootCmd.AddCommand(absenceCmd)
	absenceCmd.Flags().StringP("year", "y", "", "--year=2021")
	absenceCmd.Flags().StringP("format", "f", "table", "--format=json")
}
