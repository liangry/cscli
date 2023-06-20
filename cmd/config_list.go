/*
Copyright © 2023 liangry

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var configListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all configurations in Config Server",
	Long: `A command line tool for Alibaba iLogtail Config Server

config list: List all configurations in Config Server
	`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list config")
	},
}

func init() {
	configCmd.AddCommand(configListCmd)
}
