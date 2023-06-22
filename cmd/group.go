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
	Aliases: []string{"g", "gr", "gro", "grou"},
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
	if len(args) == 0 {
		return errors.New("User defined agent group must have at least one tag")
	}
	m := make(map[string]bool)
	for i := 0; i < len(args); i += 2 {
		tag := args[i]
		if m[tag] {
			return errors.New("Duplicate tag " + tag)
		} else {
			m[tag] = true
		}
	}
	return nil
}

var defaultGroupProtector = func(cmd *cobra.Command, args []string) error {
	if groupName == "default" {
		return errors.New("Default group cannot be overwritten")
	}
	return nil
}

func init() {
	rootCmd.AddCommand(groupCmd)
}
