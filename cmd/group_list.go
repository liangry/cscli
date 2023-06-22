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

var groupListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all agent groups in Config Server",
	Long: `A command line tool for Alibaba iLogtail Config Server

group list: List all agent groups in Config Server
	`,
	Aliases: []string{"l", "li", "lis"},
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		reqBody := configserverproto.ListAgentGroupsRequest{}
		reqBody.RequestId = uuid.New().String()
		reqBodyByte, _ := proto.Marshal(&reqBody)

		statusCode, resBodyByte, err := httpclient.SendRequest("ListAgentGroups", reqBodyByte)
		if err != nil {
			return err
		}

		resBody := new(configserverproto.ListAgentGroupsResponse)
		proto.Unmarshal(resBodyByte, resBody)
		if statusCode != http.StatusOK {
			code := resBody.Code.String()
			if len(code) > 0 && code != "ACCEPT" {
				return errors.New(fmt.Sprintf("%s - %s", resBody.Code, resBody.Message))
			}

			return errors.New(string(resBodyByte))
		}

		w := tabwriter.NewWriter(os.Stdout, 3, 0, 3, ' ', tabwriter.TabIndent)
		fmt.Fprintln(w, "Seq\tGroup Name\tDescription\tTags")
		fmt.Fprintln(w, "---\t----------\t-----------\t----")
		for i, agentGroup := range resBody.AgentGroups {
			seq := fmt.Sprintf("%3d", i + 1)
			var tags []string
			for _, tag := range agentGroup.Tags {
				tags = append(tags, fmt.Sprintf("%s=%s", tag.Name, tag.Value))
			}
			tagString := strings.Join(tags, ", ")
			content := fmt.Sprintf("%s\t%s\t%s\t%s", seq, agentGroup.GroupName, agentGroup.Description, tagString)
			fmt.Fprintln(w, content)
		}
		w.Flush()

		return nil
	},
}

func init() {
	groupCmd.AddCommand(groupListCmd)
}
