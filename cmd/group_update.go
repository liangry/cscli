/*
Copyright © 2023 liangry

*/
package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

var groupUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update agent group",
	Long: `A command line tool for Alibaba iLogtail Config Server

group update: Update agent group
	`,
	Args: tagValidator,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("User defined agent group must have at least one tag")
		}
		fmt.Println("update agent group", groupName)
		for i := 0; i < len(args); i += 2 {
			fmt.Println("tag", args[i], "=", args[i + 1])
		}
		return nil
	},
}

func init() {
	groupCmd.AddCommand(groupUpdateCmd)

	groupUpdateCmd.Flags().StringVarP(&groupName, "name", "n", "", "agent group name")
	groupUpdateCmd.Flags().StringVarP(&description, "desc", "d", "", "description of the agent group")
	groupUpdateCmd.MarkFlagRequired("name")
}
