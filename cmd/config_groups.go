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

	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"

	configserverproto "github.com/alibaba/ilogtail/config_server/service/proto"

	"github.com/liangry/cscli/httpclient"
	"github.com/spf13/cobra"
)

var configGroupsCmd = &cobra.Command{
	Use:   "groups",
	Short: "List agent groups associated with the specific config",
	Long: `A command line tool for Alibaba iLogtail Config Server

config groups: List agent groups associated with the specific config
	`,
	Aliases: []string{"gr", "gro", "grou", "group"},
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		reqBody := configserverproto.GetAppliedAgentGroupsRequest{}
		reqBody.RequestId = uuid.New().String()
		reqBody.ConfigName = configName
		reqBodyByte, _ := proto.Marshal(&reqBody)

		statusCode, resBodyByte, err := httpclient.SendRequest("GetAppliedAgentGroups", reqBodyByte)
		if err != nil {
			return err
		}

		resBody := new(configserverproto.GetAppliedAgentGroupsResponse)
		proto.Unmarshal(resBodyByte, resBody)
		if statusCode != http.StatusOK {
			code := resBody.Code.String()
			if len(code) > 0 && code != "ACCEPT" {
				return errors.New(fmt.Sprintf("%s - %s", resBody.Code, resBody.Message))
			}

			return errors.New(string(resBodyByte))
		}

		w := tabwriter.NewWriter(os.Stdout, 3, 0, 3, ' ', tabwriter.TabIndent)
		fmt.Fprintln(w, "Seq\tGroup Name")
		fmt.Fprintln(w, "---\t----------")
		for i, groupName := range resBody.AgentGroupNames {
			seq := fmt.Sprintf("%3d", i + 1)
			content := fmt.Sprintf("%s\t%s", seq, groupName)
			fmt.Fprintln(w, content)
		}
		w.Flush()

		return nil
	},
}

func init() {
	configCmd.AddCommand(configGroupsCmd)

	configGroupsCmd.Flags().StringVarP(&configName, "name", "n", "", "config name")
	configGroupsCmd.MarkFlagRequired("name")
}
