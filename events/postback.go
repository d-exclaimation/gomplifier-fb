//
//  postback.go
//  events
//
//  Created by d-exclaimation on 10:15 PM.
//  Copyright © 2020 d-exclaimation. All rights reserved.
//

package events

import (
	. "github.com/d-exclaimation/fb-messenger-api/models"
)

func HandlePostBackEvent(senderPSID string, postback *Postback) {
	var (
		requestBody ReplyItem
		payload = postback.Payload
	)

	if payload == string(PayloadYes) {
		requestBody = &ReplyMessage{
			Text: "Thanks!!!",
		}
	} else if payload == string(PayloadNo) {
		requestBody = &ReplyMessage{
			Text: "Oops, try sending another image",
		}
	}
	sendMessageAPI(createReplyEvent(senderPSID, requestBody))
}
