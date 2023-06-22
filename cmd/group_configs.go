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

var groupConfigsCmd = &cobra.Command{
	Use:   "configs",
	Short: "List configurations associated with the specific agent group",
	Long: `A command line tool for Alibaba iLogtail Config Server

group configs: List configurations associated with the specific agent group
	`,
	Aliases: []string{"co", "con", "conf", "confi", "config"},
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		reqBody := configserverproto.GetAppliedConfigsForAgentGroupRequest{}
		reqBody.RequestId = uuid.New().String()
		reqBody.GroupName = groupName
		reqBodyByte, _ := proto.Marshal(&reqBody)

		statusCode, resBodyByte, err := httpclient.SendRequest("GetAppliedConfigsForAgentGroup", reqBodyByte)
		if err != nil {
			return err
		}

		resBody := new(configserverproto.GetAppliedConfigsForAgentGroupResponse)
		proto.Unmarshal(resBodyByte, resBody)
		if statusCode != http.StatusOK {
			code := resBody.Code.String()
			if len(code) > 0 && code != "ACCEPT" {
				return errors.New(fmt.Sprintf("%s - %s", resBody.Code, resBody.Message))
			}

			return errors.New(string(resBodyByte))
		}

		w := tabwriter.NewWriter(os.Stdout, 3, 0, 3, ' ', tabwriter.TabIndent)
		fmt.Fprintln(w, "Seq\tConfig Name")
		fmt.Fprintln(w, "---\t-----------")
		for i, configName := range resBody.ConfigNames {
			seq := fmt.Sprintf("%3d", i + 1)
			content := fmt.Sprintf("%s\t%s", seq, configName)
			fmt.Fprintln(w, content)
		}
		w.Flush()

		return nil
	},
}

func init() {
	groupCmd.AddCommand(groupConfigsCmd)

	groupConfigsCmd.Flags().StringVarP(&groupName, "name", "n", "", "agent group name")
	groupConfigsCmd.MarkFlagRequired("name")
}
