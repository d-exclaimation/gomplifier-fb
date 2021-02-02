//
//  fb-reply.go
//  models
//
//  Created by d-exclaimation on 10:50 PM.
//  Copyright © 2020 d-exclaimation. All rights reserved.
//

package models

type ReplyTypes string

const (
	ResponseReply ReplyTypes = "RESPONSE"
	UpdateReply   ReplyTypes = "UPDATE"
	MessageTagReply ReplyTypes = "MESSAGE_TAG"
)

type ReplyEvent struct {
	MessagingTypes ReplyTypes   `json:"messaging_types"`
	Recipient FbUser 		`json:"recipient"`
	Message ReplyItem   	`json:"message"`
}
