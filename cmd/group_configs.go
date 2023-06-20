/*
Copyright © 2023 liangry

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var groupConfigsCmd = &cobra.Command{
	Use:   "configs",
	Short: "List configurations associated with the specific agent group",
	Long: `A command line tool for Alibaba iLogtail Config Server

group configs: List configurations associated with the specific agent group
	`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list applied configs of group", groupName)
	},
}

func init() {
	groupCmd.AddCommand(groupConfigsCmd)

	groupConfigsCmd.Flags().StringVarP(&groupName, "name", "n", "", "agent group name")
	groupConfigsCmd.MarkFlagRequired("name")
}
