//
//  message.go
//  events
//
//  Created by d-exclaimation on 7:54 PM.
//  Copyright © 2020 d-exclaimation. All rights reserved.
//

package events

import (
	"github.com/d-exclaimation/fb-messenger-api/games"
	. "github.com/d-exclaimation/fb-messenger-api/models"
	"strings"
)

func HandleMessageEvent(senderPSID string, message *Message, gameDocument map[string]*games.Sokoban) {
	var requestBody ReplyItem
	if message.Text != "" {

		_, ok := gameDocument[senderPSID]
		if isWASD(message.Text) && ok {
			// Continue game
			requestBody = games.MoveGame(senderPSID, message.Text, gameDocument)
			sendMessageAPI(createReplyEvent(senderPSID, requestBody))
			return
		}

		if !strings.HasPrefix(message.Text, "!") {
			return
		}

		var commandName = strings.Split(message.Text, " ")[0]

		if commandName == "!start" {
			// Start a new one
			requestBody = games.NewGame(senderPSID, gameDocument)
			sendMessageAPI(createReplyEvent(senderPSID, requestBody))
			return
		}

	} else if message.Attachments != nil {
		var attachmentUrl = message.Attachments[0].Payload.Url

		// Send a reply with attachments
		requestBody = NewHGenericImage(
			"Is this the right picture", "Tap a button to confirm",
			attachmentUrl, nil,
			YesPostBack(),
			NoPostBack())
		sendMessageAPI(createReplyEvent(senderPSID, requestBody))
	}
}


func isWASD(word string) bool {
	var wasd = map[string]bool{
		"w": true,
		"a": true,
		"s": true,
		"d": true,
	}
	_, ok := wasd[strings.ToLower(word)]
	return ok
}
