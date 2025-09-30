package main

import (
	"MessagePlugin/core"
	"log"

	"github.com/MikaBot-Project/MikaPluginLib/pluginIO"
)

func main() {
	pluginIO.MessageRegister(core.Message)
	pluginIO.NoticeRegister("group_increase", core.WelcomeMessage)
	var data pluginIO.Message
	for {
		data = <-pluginIO.MessageChan
		log.Println("type: ", data.PostType)
		log.Println("echo: ", data.RawMessage)
	}
}
