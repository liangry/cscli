/*
Copyright © 2023 liangry

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var configGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a configuration in yaml format by name",
	Long: `A command line tool for Alibaba iLogtail Config Server

config get: Get a configuration in yaml format by name
	`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get config", configName)
	},
}

func init() {
	configCmd.AddCommand(configGetCmd)

	configGetCmd.Flags().StringVarP(&configName, "name", "n", "", "config name")
	configGetCmd.MarkFlagRequired("name")
}
