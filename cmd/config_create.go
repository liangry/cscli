/*
Copyright © 2023 liangry

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var configCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a named configuration in yaml format",
	Long: `A command line tool for Alibaba iLogtail Config Server

config create: Create a named configuration in yaml format
	`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create config", configName)
	},
}

func init() {
	configCmd.AddCommand(configCreateCmd)

	configCreateCmd.Flags().StringVarP(&configName, "name", "n", "", "config name")
	configCreateCmd.Flags().StringVarP(&fileName, "file", "f", "", "config file in yaml format")
	configCreateCmd.MarkFlagRequired("name")
	configCreateCmd.MarkFlagRequired("file")
}
