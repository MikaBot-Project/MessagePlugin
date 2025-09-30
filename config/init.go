package config

import (
	"github.com/MikaBot-Project/MikaPluginLib/pluginConfig"
)

var GroupMsgMap map[string]map[string][]any
var WelcomeMsgMap map[string]any

func init() {
	pluginConfig.ReadAllJson("text/", &GroupMsgMap)
	pluginConfig.ReadJson("welcome.json", &WelcomeMsgMap)
}
