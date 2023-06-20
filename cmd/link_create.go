/*
Copyright © 2023 liangry

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var linkCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Apply configuration to agent group",
	Long: `A command line tool for Alibaba iLogtail Config Server

link create: Apply configuration to agent group
	`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("apply config", configName, "to agent group", groupName)
	},
}

func init() {
	linkCmd.AddCommand(linkCreateCmd)
}
