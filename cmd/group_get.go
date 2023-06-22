/*
Copyright © 2023 liangry

*/
package cmd

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"

	configserverproto "github.com/alibaba/ilogtail/config_server/service/proto"

	"github.com/liangry/cscli/httpclient"
	"github.com/spf13/cobra"
)

var groupGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get an agent group by name",
	Long: `A command line tool for Alibaba iLogtail Config Server

group get: Get an agent group by name
	`,
	Aliases: []string{"g", "ge"},
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		reqBody := configserverproto.GetAgentGroupRequest{}
		reqBody.RequestId = uuid.New().String()
		reqBody.GroupName = groupName
		reqBodyByte, _ := proto.Marshal(&reqBody)

		statusCode, resBodyByte, err := httpclient.SendRequest("GetAgentGroup", reqBodyByte)
		if err != nil {
			return err
		}

		resBody := new(configserverproto.GetAgentGroupResponse)
		proto.Unmarshal(resBodyByte, resBody)
		if statusCode != http.StatusOK {
			code := resBody.Code.String()
			if len(code) > 0 && code != "ACCEPT" {
				return errors.New(fmt.Sprintf("%s - %s", resBody.Code, resBody.Message))
			}

			return errors.New(string(resBodyByte))
		}

		w := tabwriter.NewWriter(os.Stdout, 3, 0, 3, ' ', tabwriter.TabIndent)
		fmt.Fprintln(w, "Group Name\tDescription\tTags")
		fmt.Fprintln(w, "----------\t-----------\t----")
		agentGroup := resBody.AgentGroup
		var tags []string
		for _, tag := range agentGroup.Tags {
			tags = append(tags, fmt.Sprintf("%s=%s", tag.Name, tag.Value))
		}
		tagString := strings.Join(tags, ", ")
		content := fmt.Sprintf("%s\t%s\t%s", agentGroup.GroupName, agentGroup.Description, tagString)
		fmt.Fprintln(w, content)
		w.Flush()

		return nil
	},
}

func init() {
	groupCmd.AddCommand(groupGetCmd)

	groupGetCmd.Flags().StringVarP(&groupName, "name", "n", "", "agent group name")
	groupGetCmd.MarkFlagRequired("name")
}
