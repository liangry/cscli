/*
Copyright © 2023 liangry

*/
package cmd

import (
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Collection of configuration",
	Long: `A command line tool for Alibaba iLogtail Config Server

config: Collection of configuration
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Usage()
	},
}

var configName string
var fileName string

func init() {
	rootCmd.AddCommand(configCmd)
}
