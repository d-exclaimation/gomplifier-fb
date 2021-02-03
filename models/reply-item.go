//
//  reply-item.go
//  models
//
//  Created by d-exclaimation on 12:30 PM.
//  Copyright © 2020 d-exclaimation. All rights reserved.
//

package models

//
// ReplyItem is all the types of Message that can be used in the ReplyEvent
//

// Reply item Interface
type ReplyItem interface {
	ConformToReplyItem()
}

// Reply Message, Send a reply with just text
type ReplyMessage struct {
	Text string `json:"text"`
}

func (*ReplyMessage) ConformToReplyItem() {}

/*
Example:
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


// Reply Attachments, Send a reply with a bunch of attachments
type ReplyAttachments struct {
	Attachment *ReplyAttachment `json:"attachment"`
}

func (*ReplyAttachments) ConformToReplyItem() {}

/*
Example:
{
  "messaging_type": "<MESSAGING_TYPE>",
  "recipient": {
    "id": "<PSID>"
  },
  "message": {
    "attachment":{
      "type":"template",
      "payload":{
        "template_type":"generic",
        "elements":[
           {
            "title":"Welcome!",
            "image_url":"https://petersfancybrownhats.com/company_image.png",
            "subtitle":"We have the right hat for everyone.",
            "default_action": {
              "type": "web_url",
              "url": "https://petersfancybrownhats.com/view?item=103",
              "webview_height_ratio": "tall",
            },
            "buttons":[
              {
                "type":"web_url",
                "url":"https://petersfancybrownhats.com",
                "title":"View Website"
              },{
                "type":"postback",
                "title":"Start Chatting",
                "payload":"DEVELOPER_DEFINED_PAYLOAD"
              }
            ]
          }
        ]
      }
    }
  }
}
*/

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
