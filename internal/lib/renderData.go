/*
Copyright © 2021 Henning Dahlheim <hactar@cyberkraft.ch>

This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at https://mozilla.org/MPL/2.0/.

*/
package lib

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

func RenderData(format string, headers []string, data [][]string, raw interface{}) error {
	switch format {
	case "table":
		table := tablewriter.NewWriter(os.Stdout)
		table.SetAutoWrapText(false)
		table.SetAlignment(tablewriter.ALIGN_LEFT)
		table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
		table.SetHeader(headers)
		table.SetBorder(false)
		table.SetRowLine(true)
		table.AppendBulk(data)
		table.Render()
	case "json":
		enc := json.NewEncoder(os.Stdout)

		err := enc.Encode(raw)
		if err != nil {
			return err
		}
	case "csv":
		enc := csv.NewWriter(os.Stdout)

		err := enc.WriteAll(data)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("unknown format")
	}

	return nil
}
