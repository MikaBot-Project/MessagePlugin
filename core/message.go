package core

import (
	"MessagePlugin/config"
	"regexp"
	"strconv"

	"github.com/MikaBot-Project/MikaPluginLib/pluginIO"
)

func Message(msg pluginIO.Message) {
	for _, messageItem := range msg.MessageArray {
		if messageItem.Type == "text" {
			text := messageItem.GetString("text")
			if msg.AtMe || msg.GroupId == 0 {
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
						return
					}
				}
			}
			message, ok := config.GroupMsgMap[strconv.FormatInt(msg.GroupId, 10)][text]
			if ok {
				sendMessage(message[random.Intn(len(message))], &msg)
				return
			} else {
				message, ok = config.GroupMsgMap["0"][text]
				if ok {
					sendMessage(message[random.Intn(len(message))], &msg)
					return
				}
			}
			var messageMap map[string][]any
			messageMap, ok = config.GroupRegexpMap[strconv.FormatInt(msg.GroupId, 10)]
			if ok {
				for key, messageMapItem := range messageMap {
					matchString, err := regexp.MatchString(key, text)
					if err != nil {
						return
					}
					if matchString {
						sendMessage(messageMapItem[random.Intn(len(messageMapItem))], &msg)
						return
					}
				}
			}
			messageMap, ok = config.GroupRegexpMap["0"]
			if ok {
				for key, messageMapItem := range messageMap {
					matchString, err := regexp.MatchString(key, text)
					if err != nil {
						return
					}
					if matchString {
						sendMessage(messageMapItem[random.Intn(len(messageMapItem))], &msg)
						return
					}
				}
			}
		}
	}
}
