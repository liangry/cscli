/*
Copyright © 2023 liangry

*/
package cmd

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"

	configserverproto "github.com/alibaba/ilogtail/config_server/service/proto"

	"github.com/liangry/cscli/httpclient"
	"github.com/spf13/cobra"
)

var configUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update configuration with the new file",
	Long: `A command line tool for Alibaba iLogtail Config Server

config update: Update configuration with the new file
	`,
	Aliases: []string{"u", "up", "upd", "upda", "updat"},
	Args: cobra.NoArgs,
	PreRunE: validateConfig,
	RunE: func(cmd *cobra.Command, args []string) error {
		reqBody := configserverproto.UpdateConfigRequest{}
		reqBody.RequestId = uuid.New().String()
		configDetail := &configserverproto.ConfigDetail{}
		configDetail.Name = configName
		configDetail.Type = configserverproto.ConfigType_PIPELINE_CONFIG
		configDetail.Detail = fileContent
		reqBody.ConfigDetail = configDetail
		reqBodyByte, _ := proto.Marshal(&reqBody)

		statusCode, resBodyByte, err := httpclient.SendRequest("UpdateConfig", reqBodyByte)
		if err != nil {
			return err
		}

		resBody := new(configserverproto.UpdateConfigResponse)
		proto.Unmarshal(resBodyByte, resBody)
		if statusCode != http.StatusOK {
			code := resBody.Code.String()
			if len(code) > 0 && code != "ACCEPT" {
				return errors.New(fmt.Sprintf("%s - %s", resBody.Code, resBody.Message))
			}

			return errors.New(string(resBodyByte))
		}

		fmt.Println(resBody.Message)
		return nil
	},
}

func init() {
	configCmd.AddCommand(configUpdateCmd)

	configUpdateCmd.Flags().StringVarP(&configName, "name", "n", "", "config name")
	configUpdateCmd.Flags().StringVarP(&fileName, "file", "f", "", "config file in yaml format")
	configUpdateCmd.MarkFlagRequired("name")
	configUpdateCmd.MarkFlagRequired("file")
}
