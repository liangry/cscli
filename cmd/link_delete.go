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

var linkDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Remove configuration from agent group",
	Long: `A command line tool for Alibaba iLogtail Config Server

link delete: Remove configuration from agent group
	`,
	Aliases: []string{"d", "de", "del", "dele", "delet"},
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		reqBody := configserverproto.RemoveConfigFromAgentGroupRequest{}
		reqBody.RequestId = uuid.New().String()
		reqBody.GroupName = groupName
		reqBody.ConfigName = configName
		reqBodyByte, _ := proto.Marshal(&reqBody)

		statusCode, resBodyByte, err := httpclient.SendRequest("RemoveConfigFromAgentGroup", reqBodyByte)
		if err != nil {
			return err
		}

		resBody := new(configserverproto.RemoveConfigFromAgentGroupResponse)
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
	linkCmd.AddCommand(linkDeleteCmd)
}
