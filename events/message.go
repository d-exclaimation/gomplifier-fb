//
//  message.go
//  events
//
//  Created by d-exclaimation on 7:54 PM.
//  Copyright © 2020 d-exclaimation. All rights reserved.
//

package events

import (
	"github.com/d-exclaimation/fb-messenger-api/commands"
	"github.com/d-exclaimation/fb-messenger-api/games"
	. "github.com/d-exclaimation/fb-messenger-api/models"
	"strings"
)

func HandleMessageEvent(senderPSID string, message *Message, gameDocument map[string]*games.Sokoban) {
	if message.Text != "" {
		_, ok := gameDocument[senderPSID] // Check if game existed here

		if isWASD(message.Text) && ok {
			// Continue game
			sendMessageAPI(createReplyEvent(senderPSID, games.MoveGame(senderPSID, message.Text, gameDocument)))
			return
		}

		// if command does not start with a prefix quit early
		if !strings.HasPrefix(message.Text, "!") {
			return
		}

		if message.Text == "!start" {
			// Start a new one
			sendMessageAPI(createReplyEvent(senderPSID, games.NewGame(senderPSID, gameDocument)))
			return
		}

		// Check for any other commands
		var (
		 	commandName = strings.Split(message.Text, " ")[0]
			commandlist = commands.All()
			exec, exist = commandlist[commandName]
		)
		if !exist {
			return
		}

		sendMessageAPI(createReplyEvent(senderPSID, exec(senderPSID, message.Text)))


	} else if message.Attachments != nil {
		var attachmentUrl = message.Attachments[0].Payload.Url

		// Send a reply with attachments
		var requestBody = NewHGenericImage(
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
