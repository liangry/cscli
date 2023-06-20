/*
Copyright © 2023 liangry

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var groupAgentsCmd = &cobra.Command{
	Use:   "agents",
	Short: "List all agents joined the specific agent group",
	Long: `A command line tool for Alibaba iLogtail Config Server

group agents: List all agents joined the specific agent group
	`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list agents of group", groupName)
	},
}

func init() {
	groupCmd.AddCommand(groupAgentsCmd)

	groupAgentsCmd.Flags().StringVarP(&groupName, "name", "n", "", "agent group name")
	groupAgentsCmd.MarkFlagRequired("name")
}
