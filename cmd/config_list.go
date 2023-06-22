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

var configListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all configurations in Config Server",
	Long: `A command line tool for Alibaba iLogtail Config Server

config list: List all configurations in Config Server
	`,
	Aliases: []string{"l", "li", "lis"},
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		reqBody := configserverproto.ListConfigsRequest{}
		reqBody.RequestId = uuid.New().String()
		reqBodyByte, _ := proto.Marshal(&reqBody)

		statusCode, resBodyByte, err := httpclient.SendRequest("ListConfigs", reqBodyByte)
		if err != nil {
			return err
		}

		resBody := new(configserverproto.ListConfigsResponse)
		proto.Unmarshal(resBodyByte, resBody)
		if statusCode != http.StatusOK {
			code := resBody.Code.String()
			if len(code) > 0 && code != "ACCEPT" {
				return errors.New(fmt.Sprintf("%s - %s", resBody.Code, resBody.Message))
			}

			return errors.New(string(resBodyByte))
		}

		w := tabwriter.NewWriter(os.Stdout, 3, 0, 3, ' ', tabwriter.TabIndent)
		fmt.Fprintln(w, "Seq\tConfig Name\tConfig Type")
		fmt.Fprintln(w, "---\t-----------\t-----------")
		for i, configDetail := range resBody.ConfigDetails {
			seq := fmt.Sprintf("%3d", i + 1)
			content := fmt.Sprintf("%s\t%s\t%s", seq, configDetail.Name, configDetail.Type)
			fmt.Fprintln(w, content)
		}
		w.Flush()

		return nil
	},
}

func init() {
	configCmd.AddCommand(configListCmd)
}
