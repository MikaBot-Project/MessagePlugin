package core

import (
	"MessagePlugin/config"
	"strconv"

	"github.com/MikaBot-Project/MikaPluginLib/pluginIO"
)

func Message(msg pluginIO.Message) {
	for _, messageItem := range msg.MessageArray {
		if messageItem.Type == "text" {
			message, ok := config.GroupMsgMap[strconv.FormatInt(msg.GroupId, 10)][messageItem.GetString("text")]
			if ok {
				sendMessage(message[random.Intn(len(message))], &msg)
			} else {
				message, ok = config.GroupMsgMap["0"][messageItem.GetString("text")]
				if ok {
					sendMessage(message[random.Intn(len(message))], &msg)
				}
			}
		}
	}
}
