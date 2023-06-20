/*
Copyright © 2023 liangry

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var groupDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an agent group by name",
	Long: `A command line tool for Alibaba iLogtail Config Server

group delete: Delete an agent group by name
	`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete agent group", groupName)
	},
}

func init() {
	groupCmd.AddCommand(groupDeleteCmd)

	groupDeleteCmd.Flags().StringVarP(&groupName, "name", "n", "", "agent group name")
	groupDeleteCmd.MarkFlagRequired("name")
}
