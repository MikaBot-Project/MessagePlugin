package core

import (
	"MessagePlugin/config"
	"strconv"

	"github.com/MikaBot-Project/MikaPluginLib/pluginIO"
)

func WelcomeMessage(msg pluginIO.Message) {
	message, ok := config.WelcomeMsgMap[strconv.FormatInt(msg.GroupId, 10)]
	msg.AtMe = true
	if ok {
		sendMessage(message, &msg)
	} else {
		message, ok = config.GroupMsgMap["0"]
		if ok {
			sendMessage(message, &msg)
		}
	}
}
