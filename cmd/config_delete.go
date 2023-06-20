/*
Copyright © 2023 liangry

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var configDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a configuration by name",
	Long: `A command line tool for Alibaba iLogtail Config Server

config delete: Delete a configuration by name
	`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete config", configName)
	},
}

func init() {
	configCmd.AddCommand(configDeleteCmd)

	configDeleteCmd.Flags().StringVarP(&configName, "name", "n", "", "config name")
	configDeleteCmd.MarkFlagRequired("name")
}
