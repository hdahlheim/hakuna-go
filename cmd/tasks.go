/*
Copyright © 2021 Henning Dahlheim <hactar@cyberkraft.ch>

*/
package cmd

import (
	"fmt"

	"github.com/hdahlheim/hakuna-go/internal/lib"
	"github.com/spf13/cobra"
)

// const taskTpl = `--------Task--------------------------------
// Id:            %v
// Name:          %v
// Archived:      %v
// `

// tasksCmd represents the tasks command
var tasksCmd = &cobra.Command{
	Use:   "tasks",
	Short: "List all task",
	RunE: func(cmd *cobra.Command, args []string) error {
		format, err := cmd.LocalFlags().GetString("format")
		if err != nil {
			return err
		}

		h := initHakunaClient()

		tasks, err := h.GetTasks()
		if err != nil {
			return err
		}

		data := make([][]string, len(tasks))
		for i, row := range data {
			task := tasks[i]
			data[i] = append(row,
				fmt.Sprint(task.ID),
				fmt.Sprint(task.Name),
				fmt.Sprint(task.Archived),
			)
		}

		return lib.RenderData(
			format,
			[]string{"ID", "Name", "Archived"},
			data,
			tasks,
		)
	},
}

func init() {
	rootCmd.AddCommand(tasksCmd)
	tasksCmd.Flags().StringP("format", "f", "table", "output format defaults to table (table, json, csv)")
}
