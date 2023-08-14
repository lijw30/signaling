package main

import (
	"github.com/lijw30/signaling/src/action"
	"github.com/lijw30/signaling/src/framework"
)

func init() {
	framework.GActionRouter["/xrtcclient/push"] = action.NewXrtcClientAction()
}
