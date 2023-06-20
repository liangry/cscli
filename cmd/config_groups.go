/*
Copyright © 2023 liangry

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var configGroupsCmd = &cobra.Command{
	Use:   "groups",
	Short: "List agent groups associated with the specific config",
	Long: `A command line tool for Alibaba iLogtail Config Server

config groups: List agent groups associated with the specific config
	`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list applied agent groups of config", configName)
	},
}

func init() {
	configCmd.AddCommand(configGroupsCmd)

	configGroupsCmd.Flags().StringVarP(&configName, "name", "n", "", "config name")
	configGroupsCmd.MarkFlagRequired("name")
}
