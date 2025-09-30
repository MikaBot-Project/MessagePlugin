package core

import (
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
		pluginIO.SendMessage(message.(string), msg.UserId, msg.GroupId)
	case []any:
		log.Println("send msg", message)
		var send []pluginIO.MessageItem
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
