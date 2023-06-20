/*
Copyright © 2023 liangry

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var groupGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get an agent group by name",
	Long: `A command line tool for Alibaba iLogtail Config Server

group get: Get an agent group by name
	`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get agent group", groupName)
	},
}

func init() {
	groupCmd.AddCommand(groupGetCmd)

	groupGetCmd.Flags().StringVarP(&groupName, "name", "n", "", "agent group name")
	groupGetCmd.MarkFlagRequired("name")
}
