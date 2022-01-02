/*
Copyright © 2021 Henning Dahlheim <hactar@cyberkraft.ch>

This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at https://mozilla.org/MPL/2.0/.

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
