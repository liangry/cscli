/*
Copyright © 2023 liangry

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var groupListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all agent groups in Config Server",
	Long: `A command line tool for Alibaba iLogtail Config Server

group list: List all agent groups in Config Server
	`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list agent group")
	},
}

func init() {
	groupCmd.AddCommand(groupListCmd)
}
