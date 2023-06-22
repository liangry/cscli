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

var configGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a configuration in yaml format by name",
	Long: `A command line tool for Alibaba iLogtail Config Server

config get: Get a configuration in yaml format by name
	`,
	Aliases: []string{"ge"},
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		reqBody := configserverproto.GetConfigRequest{}
		reqBody.RequestId = uuid.New().String()
		reqBody.ConfigName = configName
		reqBodyByte, _ := proto.Marshal(&reqBody)

		statusCode, resBodyByte, err := httpclient.SendRequest("GetConfig", reqBodyByte)
		if err != nil {
			return err
		}

		resBody := new(configserverproto.GetConfigResponse)
		proto.Unmarshal(resBodyByte, resBody)
		if statusCode != http.StatusOK {
			code := resBody.Code.String()
			if len(code) > 0 && code != "ACCEPT" {
				return errors.New(fmt.Sprintf("%s - %s", resBody.Code, resBody.Message))
			}

			return errors.New(string(resBodyByte))
		}

		fmt.Println(resBody.ConfigDetail.Detail)
		return nil
	},
}

func init() {
	configCmd.AddCommand(configGetCmd)

	configGetCmd.Flags().StringVarP(&configName, "name", "n", "", "config name")
	configGetCmd.MarkFlagRequired("name")
}
