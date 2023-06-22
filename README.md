# Alibaba iLogtail Config Server 命令行工具

这是基于阿里巴巴 iLogtail 项目 Config Server [通信协议](https://ilogtail.gitbook.io/ilogtail-docs/config-server/communication-protocol)的一个命令行实现，使用 [Cobra](https://cobra.dev/) 开发。

## 快速开始

```shell
git clone https://github.com/liangry/cscli
cd cscli
go build
# 执行命令前，先启动 iLogtai Config Server
export CONFIG_SERVER_ADDRESS=http://127.0.0.1:8899
./cscli group list
```

## 命令速查

| 命令 | 接口 | 示例 | 缩写 |
| - | - | - | - |
| group list | ListAgentGroups | cscli group list | cscli g l |
| group get | GetAgentGroup | cscli group get --name default | cscli g g -n default |
| group create | CreateAgentGroup | cscli group create --name test --desc "unit test" cluster unit | cscli g cr -n test -d "unit test" cluster unit |
| group update | UpdateAgentGroup | cscli group update --name test --desc "uat test" cluster uat | cscli g u -n test -d "uat test" cluster uat |
| group delete | DeleteAgentGroup | cscli group delete --name test | cscli g d -n test |
| config create | CreateConfig | cscli config create --name test --file /path/to/config | cscli c c -n test -f /path/to/config |
| config list | ListConfigs | cscli config list | cscli c l |
| config get | GetConfig | cscli config get --name test | cscli c ge -n test |
| config update | UpdateConfig | cscli config update --name test --file /path/to/config | cscli c u -n test -f /path/to/config |
| config delete | DeleteConfig | cscli config delete --name test | cscli c d -n test |
| link create | ApplyConfigToAgentGroup | cscli link create --group default --config test | cscli l c -g default -c test |
| link delete | RemoveConfigFromAgentGroup | cscli link delete --group default --config test | cscli l d -g default -c test |
| config groups | GetAppliedAgentGroups | cscli config groups --name test | cscli c gr -n test |
| group configs | GetAppliedConfigsForAgentGroup | cscli group configs --name default | cscli g co -n default |
| group agents | ListAgents | cscli group agents --name default | cscli g a -n default |

