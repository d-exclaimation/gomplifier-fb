//
//  event.go
//  events
//
//  Created by d-exclaimation on 5:55 PM.
//  Copyright © 2020 d-exclaimation. All rights reserved.
//

package events

import (
	. "github.com/d-exclaimation/fb-messenger-api/models"
)

func createReplyEvent(senderPSID string, response ReplyItem) *ReplyEvent {
	return &ReplyEvent{
		MessagingTypes: ResponseReply,
		Recipient: FbUser{
			ID: senderPSID,
		},
		Message:   response,
	}
}

