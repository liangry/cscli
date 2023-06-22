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

var groupUpdateCmd = &cobra.Command{
	Use:   "update [flags] tag1_name tag1_value tag2_name tag2_value ...",
	Short: "Update agent group",
	Long: `A command line tool for Alibaba iLogtail Config Server

group update: Update agent group
	`,
	Aliases: []string{"u", "up", "upd", "upda", "updat"},
	Args: tagValidator,
	RunE: func(cmd *cobra.Command, args []string) error {
		reqBody := configserverproto.UpdateAgentGroupRequest{}
		reqBody.RequestId = uuid.New().String()
		agentGroup := &configserverproto.AgentGroup{}
		agentGroup.GroupName = groupName
		agentGroup.Description = description
		tags := []*configserverproto.AgentGroupTag{}
		for i := 0; i < len(args); i += 2 {
			tag := &configserverproto.AgentGroupTag{}
			tag.Name = args[i]
			tag.Value = args[i + 1]
			tags = append(tags, tag)
		}
		agentGroup.Tags = tags
		reqBody.AgentGroup = agentGroup
		reqBodyByte, _ := proto.Marshal(&reqBody)

		statusCode, resBodyByte, err := httpclient.SendRequest("UpdateAgentGroup", reqBodyByte)
		if err != nil {
			return err
		}

		resBody := new(configserverproto.UpdateAgentGroupResponse)
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
	groupCmd.AddCommand(groupUpdateCmd)

	groupUpdateCmd.Flags().StringVarP(&groupName, "name", "n", "", "agent group name")
	groupUpdateCmd.Flags().StringVarP(&description, "desc", "d", "", "description of the agent group")
	groupUpdateCmd.MarkFlagRequired("name")
}
