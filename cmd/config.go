/*
Copyright © 2023 liangry

*/
package cmd

import (
	"errors"
	"io/ioutil"
	"os"

	"github.com/go-yaml/yaml"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Collection of configuration",
	Long: `A command line tool for Alibaba iLogtail Config Server

config: Collection of configuration
	`,
	Aliases: []string{"c", "co", "con", "conf", "confi"},
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Usage()
	},
}

var configName string
var fileName string
var fileContent string

type Plugin struct {
	Type string `yaml:"Type"`
}

type Config struct {
	Enable      bool         `yaml:"enable"`
	Inputs      []Plugin     `yaml:"inputs"`
	Processors  []Plugin     `yaml:"processors,omitempty"`
	Aggregators []Plugin     `yaml:"aggregators,omitempty"`
	Flushers    []Plugin     `yaml:"flushers"`
}

func init() {
	rootCmd.AddCommand(configCmd)
}

func validateConfig(cmd *cobra.Command, args []string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	fileStat, _ := file.Stat()
	if fileStat.Size() > 1024 * 1024 {
		return errors.New("File size too large")
	}

	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	var config Config
	err = yaml.Unmarshal(content, &config)
	if err != nil {
		return err
	}

	if len(config.Inputs) == 0 {
		return errors.New("Config error: missing inputs")
	}
	if len(config.Flushers) == 0 {
		return errors.New("Config error: missing flushers")
	}
	for _, input := range config.Inputs {
		if input.Type == "" {
			return errors.New("Config error: input missing Type")
		}
	}
	for _, flusher := range config.Flushers {
		if flusher.Type == "" {
			return errors.New("Config error: flusher missing Type")
		}
	}
	if len(config.Processors) > 0 {
		for _, processor := range config.Processors {
			if processor.Type == "" {
				return errors.New("Config error: processor missing Type")
			}
		}
	}
	if len(config.Aggregators) > 0 {
		for _, aggregator := range config.Aggregators {
			if aggregator.Type == "" {
				return errors.New("Config error: aggregator missing Type")
			}
		}
	}
	fileContent = string(content)
	return nil
}
