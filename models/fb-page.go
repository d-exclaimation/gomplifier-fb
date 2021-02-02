//
//  fb-page.go
//  models
//
//  Created by d-exclaimation on 5:48 PM.
//  Copyright © 2020 d-exclaimation. All rights reserved.
//

package models

// JSON Object Facebook Post Hook
type FbPostHook struct {
	Object string      `json:"object"`
	Entry  []EntryHook `json:"entry"`
}

type MessagingEvent struct {
	Sender    FbUser    `json:"sender"`
	Recipient FbUser    `json:"recipient"`
	Timestamp int64     `json:"timestamp"`
	Message   *Message  `json:"message,omitempty"`
	Postback  *Postback `json:"postback,omitempty"`
}

type EntryHook struct {
	ID        string           `json:"id"`
	Time      int64            `json:"time"`
	Messaging []MessagingEvent `json:"messaging"`
}

// Facebook and Messenger User
type FbUser struct {
	ID string `json:"id"`
}
