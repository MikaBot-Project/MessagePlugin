package core

import (
	"MessagePlugin/config"
	"strconv"

	"github.com/MikaBot-Project/MikaPluginLib/pluginIO"
)

func Message(msg pluginIO.Message) {
	for _, messageItem := range msg.MessageArray {
		if messageItem.Type == "text" {
			text := messageItem.GetString("text")
			if msg.AtMe {
				if text[0] == ' ' {
					text = text[1:]
				}
				message, ok := config.GroupAtMeMsgMap[strconv.FormatInt(msg.GroupId, 10)][text]
				if ok {
					sendMessage(message[random.Intn(len(message))], &msg)
					return
				} else {
					message, ok = config.GroupAtMeMsgMap["0"][text]
					if ok {
						sendMessage(message[random.Intn(len(message))], &msg)
					}
				}
			}
			message, ok := config.GroupMsgMap[strconv.FormatInt(msg.GroupId, 10)][text]
			if ok {
				sendMessage(message[random.Intn(len(message))], &msg)
			} else {
				message, ok = config.GroupMsgMap["0"][text]
				if ok {
					sendMessage(message[random.Intn(len(message))], &msg)
				}
			}
		}
	}
}
