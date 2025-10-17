package core

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/MikaBot-Project/MikaPluginLib/pluginIO"
)

var random *rand.Rand

func init() {
	source := rand.NewSource(time.Now().UnixNano())
	random = rand.New(source)
}

func sendMessage(message any, msg *pluginIO.Message) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("sendMessage err:", err)
		}
	}()
	switch message.(type) {
	case string:
		var send string
		if msg.AtMe {
			send = fmt.Sprintf("[CQ:at,qq=%d] %s", msg.UserId, message.(string))
		} else {
			send = message.(string)
		}
		pluginIO.SendMessage(send, msg.UserId, msg.GroupId)
	case []any:
		var send []pluginIO.MessageItem
		if msg.AtMe {
			send = append(send, pluginIO.MessageItem{
				Type: "at",
				Data: map[string]any{
					"qq": msg.UserId,
				},
			})
		}
		for _, v := range message.([]any) {
			val := v.(map[string]interface{})
			send = append(send, pluginIO.MessageItem{
				Type: val["type"].(string),
				Data: val["data"].(map[string]any),
			})
		}
		pluginIO.SendMessage(send, msg.UserId, msg.GroupId)
	}
}
