/*
Copyright © 2023 liangry

*/
package cmd

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"text/tabwriter"
	"time"

	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"

	configserverproto "github.com/alibaba/ilogtail/config_server/service/proto"

	"github.com/liangry/cscli/httpclient"
	"github.com/spf13/cobra"
)

var groupAgentsCmd = &cobra.Command{
	Use:   "agents",
	Short: "List all agents joined the specific agent group",
	Long: `A command line tool for Alibaba iLogtail Config Server

group agents: List all agents joined the specific agent group
	`,
	Aliases: []string{"a", "ag", "age", "agen", "agent"},
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error{
		reqBody := configserverproto.ListAgentsRequest{}
		reqBody.RequestId = uuid.New().String()
		reqBody.GroupName = groupName
		reqBodyByte, _ := proto.Marshal(&reqBody)

		statusCode, resBodyByte, err := httpclient.SendRequest("ListAgents", reqBodyByte)
		if err != nil {
			return err
		}

		resBody := new(configserverproto.ListAgentsResponse)
		proto.Unmarshal(resBodyByte, resBody)
		if statusCode != http.StatusOK {
			code := resBody.Code.String()
			if len(code) > 0 && code != "ACCEPT" {
				return errors.New(fmt.Sprintf("%s - %s", resBody.Code, resBody.Message))
			}

			return errors.New(string(resBodyByte))
		}

		w := tabwriter.NewWriter(os.Stdout, 3, 0, 3, ' ', tabwriter.TabIndent)
		fmt.Fprintln(w, "Seq\tID\tType\tVersion\tIP\tStartup Time\tInterval")
		fmt.Fprintln(w, "---\t--\t----\t-------\t--\t------------\t--------")
		for i, agent := range resBody.Agents {
			seq := fmt.Sprintf("%3d", i + 1)
			interval := fmt.Sprintf("%8d", agent.Interval)
			t := time.Unix(agent.StartupTime, 0)
			startupTime := t.Format("2006-01-02 15:04:05")
			content := fmt.Sprintf("%s\t%s\t%s\t%s\t%s\t%s\t%s", seq, agent.AgentId, agent.AgentType, agent.Attributes.Version, agent.Attributes.Ip, startupTime, interval)
			fmt.Fprintln(w, content)
		}
		w.Flush()

		return nil
	},
}

func init() {
	groupCmd.AddCommand(groupAgentsCmd)

	groupAgentsCmd.Flags().StringVarP(&groupName, "name", "n", "", "agent group name")
	groupAgentsCmd.MarkFlagRequired("name")
}
