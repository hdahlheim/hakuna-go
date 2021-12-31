/*
Copyright © 2021 Henning Dahlheim <hactar@cyberkraft.ch>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// pingCmd represents the ping command
var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Test the api connection",
	RunE:  pingAPI,
}

func init() {
	rootCmd.AddCommand(pingCmd)
}

func pingAPI(cmd *cobra.Command, args []string) error {
	h := initHakunaClient()
	fmt.Fprintf(os.Stderr, "Pinging api...\n")
	pong, err := h.Ping()
	if err != nil {
		return err
	}
	fmt.Printf("Pong at %v\n", pong.Pong)
	return nil
}
