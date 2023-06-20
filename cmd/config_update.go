/*
Copyright © 2023 liangry

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var configUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update configuration with the new file",
	Long: `A command line tool for Alibaba iLogtail Config Server

config update: Update configuration with the new file
	`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("update config", configName)
	},
}

func init() {
	configCmd.AddCommand(configUpdateCmd)

	configUpdateCmd.Flags().StringVarP(&configName, "name", "n", "", "config name")
	configUpdateCmd.Flags().StringVarP(&fileName, "file", "f", "", "config file in yaml format")
	configUpdateCmd.MarkFlagRequired("name")
	configUpdateCmd.MarkFlagRequired("file")
}
