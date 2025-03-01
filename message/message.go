package message

import (
	"context"
	"strings"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types/events"
	"go.vnia.dev/helper"
	"go.vnia.dev/lib"
)

// Config
var (
	prefix = "."
	self   = false
	owner  = "60102810046"
)

func Msg(client *whatsmeow.Client, msg *events.Message) {
	// simple
	simp := lib.NewSimpleImpl(client, msg)
	// dll
	from := msg.Info.Chat
	sender := msg.Info.Sender.String()
	args := strings.Split(simp.GetCMD(), " ")
	command := strings.ToLower(args[0])
	//query := strings.Join(args[1:], ` `)
	pushName := msg.Info.PushName
	isOwner := strings.Contains(sender, owner)
	//isAdmin := simp.GetGroupAdmin(from, sender)
	//isGroup := msg.Info.IsGroup
	extended := msg.Message.GetExtendedTextMessage()
	quotedMsg := extended.GetContextInfo().GetQuotedMessage()
	quotedImage := quotedMsg.GetImageMessage()
	//quotedVideo := quotedMsg.GetVideoMessage()
	//quotedSticker := quotedMsg.GetStickerMessage()
	// Self
	if self && !isOwner {
		return
	}
	// Switch Cmd
	switch command {
	case prefix + "menu":
		simp.Reply(helper.Menu(pushName, prefix))
	case prefix + "owner":
		simp.SendContact(from, owner, "aiman")
	case prefix + "source":
		simp.Reply("Source Code : https://github.com/ai-man-123/go-whatsapp-bot")
	case prefix + "sticker":
		if quotedImage != nil {
			data, _ := client.Download(quotedImage)
			stc := simp.CreateStickerIMG(data)
			client.SendMessage(context.Background(), from, "", stc)
		} else if msg.Message.GetImageMessage() != nil {
			data, _ := client.Download(msg.Message.GetImageMessage())
			stc := simp.CreateStickerIMG(data)
			client.SendMessage(context.Background(), from, "", stc)
		}
	}
}
