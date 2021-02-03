//
//  reply-event.go
//  models
//
//  Created by d-exclaimation on 10:50 PM.
//  Copyright © 2020 d-exclaimation. All rights reserved.
//

package models

//
// Reply Event is the JSON that Facebook API Accept as Reply Messages
//

type ReplyTypes string

const (
	ResponseReply ReplyTypes = "RESPONSE"
	UpdateReply   ReplyTypes = "UPDATE"
	MessageTagReply ReplyTypes = "MESSAGE_TAG"
)

// ReplyEvent Struct
type ReplyEvent struct {
	MessagingTypes ReplyTypes   `json:"messaging_types"`
	Recipient FbUser 		`json:"recipient"`
	Message ReplyItem   	`json:"message"`
}

/*
Example JSON:
{
  "messaging_type": "<MESSAGING_TYPE>",
  "recipient": {
    "id": "<PSID>"
  },
  "message": {
    "text": "hello, world!"
  }
}
 */