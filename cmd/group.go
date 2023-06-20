/*
Copyright © 2023 liangry

*/
package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

var groupCmd = &cobra.Command{
	Use:   "group",
	Short: "Collection of agent group",
	Long: `A command line tool for Alibaba iLogtail Config Server

group: Collection of agent group
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Usage()
	},
}

var groupName string
var description string

var tagValidator = func(cmd *cobra.Command, args []string) error {
	if len(args) % 2 != 0 {
		return errors.New("Tags must appear in pairs")
	}
	return nil
}

func init() {
	rootCmd.AddCommand(groupCmd)
}
