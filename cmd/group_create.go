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

var groupCreateCmd = &cobra.Command{
	Use:   "create [flags] tag1_name tag1_value tag2_name tag2_value ...",
	Short: "Create a named agent group",
	Long: `A command line tool for Alibaba iLogtail Config Server

group create: Create a named agent group
	`,
	Aliases: []string{"cr", "cre", "crea", "creac", "creat"},
	Args: tagValidator,
	RunE: func(cmd *cobra.Command, args []string) error {
		reqBody := configserverproto.CreateAgentGroupRequest{}
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

		statusCode, resBodyByte, err := httpclient.SendRequest("CreateAgentGroup", reqBodyByte)
		if err != nil {
			return err
		}

		resBody := new(configserverproto.CreateAgentGroupResponse)
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
	groupCmd.AddCommand(groupCreateCmd)

	groupCreateCmd.Flags().StringVarP(&groupName, "name", "n", "", "agent group name")
	groupCreateCmd.Flags().StringVarP(&description, "desc", "d", "", "description of the agent group")
	groupCreateCmd.MarkFlagRequired("name")
	groupCreateCmd.MarkFlagRequired("desc")
}
