//
//  reply-item.go
//  models
//
//  Created by d-exclaimation on 12:30 PM.
//  Copyright © 2020 d-exclaimation. All rights reserved.
//

package models

// Reply item Interface
type ReplyItem interface {
	ConformToReplyItem()
}

// Reply Message, Send a reply with just text
type ReplyMessage struct {
	Text string `json:"text"`
}

func (*ReplyMessage) ConformToReplyItem() {}


// Reply Attachments, Send a reply with a bunch of attachments
type ReplyAttachments struct {
	Attachment *ReplyAttachment `json:"attachment"`
}

func (*ReplyAttachments) ConformToReplyItem() {}


// Reply Attachment part of the ReplyAttachments
type ReplyAttachment struct {
	Type    string          `json:"type"`
	Payload TemplatePayload `json:"payload"`
}

func NewTemplateAttachment(payload TemplatePayload) *ReplyAttachment {
	return &ReplyAttachment{
		Type:    "template",
		Payload: payload,
	}
}
