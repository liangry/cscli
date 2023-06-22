package httpclient

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

var methods = map[string]string{
	"CreateAgentGroup": "POST",
	"UpdateAgentGroup": "PUT",
	"DeleteAgentGroup": "DELETE",
	"GetAgentGroup": "POST",
	"ListAgentGroups": "POST",

	"CreateConfig": "POST",
	"UpdateConfig": "PUT",
	"DeleteConfig": "DELETE",
	"GetConfig": "POST",
	"ListConfigs": "POST",

	"ApplyConfigToAgentGroup": "PUT",
	"RemoveConfigFromAgentGroup": "DELETE",
	"GetAppliedConfigsForAgentGroup": "POST",
	"GetAppliedAgentGroups": "POST",
	"ListAgents": "POST",
}

func SendRequest(action string, reqBodyByte []byte) (int, []byte, error) {
	configServerAddress := os.Getenv("CONFIG_SERVER_ADDRESS")
	if configServerAddress == "" {
		configServerAddress = "http://127.0.0.1:8899"
	}
	url := configServerAddress + "/User/" + action
	req, _ := http.NewRequest(methods[action], url, bytes.NewReader(reqBodyByte))
	req.Header.Set("Content-Type", "application/protobuf")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return http.StatusServiceUnavailable, nil, err
	}
	defer res.Body.Close()
	resBodyByte, _ := io.ReadAll(res.Body)
	return res.StatusCode, resBodyByte, nil
}
