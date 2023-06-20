/*
Copyright © 2023 liangry

*/
package cmd

import (
	"github.com/spf13/cobra"
)

var linkCmd = &cobra.Command{
	Use:   "link",
	Short: "Collection of association between agent group and configuration",
	Long: `A command line tool for Alibaba iLogtail Config Server

link: Collection of association between agent group and configuration
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Usage()
	},
}

func init() {
	rootCmd.AddCommand(linkCmd)

	linkCmd.PersistentFlags().StringVarP(&groupName, "group", "g", "", "agent group name")
	linkCmd.PersistentFlags().StringVarP(&configName, "config", "c", "", "config name")
	linkCmd.MarkFlagRequired("group")
	linkCmd.MarkFlagRequired("config")
}
