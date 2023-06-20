/*
Copyright © 2023 liangry

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var linkDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Remove configuration from agent group",
	Long: `A command line tool for Alibaba iLogtail Config Server

link delete: Remove configuration from agent group
	`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("remove config", configName, "from agent group", groupName)
	},
}

func init() {
	linkCmd.AddCommand(linkDeleteCmd)
}
