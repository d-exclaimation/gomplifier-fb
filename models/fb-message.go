//
//  fb-message.go
//  models
//
//  Created by d-exclaimation on 10:40 PM.
//  Copyright © 2020 d-exclaimation. All rights reserved.
//

package models

type Message struct {
	Mid  string               `json:"mid"`
	Text string               `json:"text,omitempty"`
	QuickReply *QuickReply    `json:"quick_reply,omitempty"`
	ReplyTo *ReplyTo 	   	  `json:"reply_to,omitempty"`
	Attachments []*Attachment `json:"attachments,omitempty"`
}

type QuickReply struct {
	Payload string `json:"payload"`
}

type ReplyTo struct {
	Mid string `json:"mid"`
}

type Attachment struct {
	Type    string            `json:"type"`
	Payload AttachmentPayload `json:"payload"`
}

type AttachmentPayload struct {
	Url string 				 `json:"url,omitempty"`
	Title string 			 `json:"title,omitempty"`
	StickerId string 		 `json:"sticker_id,omitempty"`
	Coordinates *Coordinates `json:"coordinates,omitempty"`
	Product *Product 		 `json:"product,omitempty"`
}

type Product struct {
	ID string 		  `json:"id,omitempty"`
	RetailerID string `json:"retailer_id,omitempty"`
	ImageURL string   `json:"image_url,omitempty"`
	Title string      `json:"title,omitempty"`
	Subtitle string   `json:"subtitle,omitempty"`
}

type Coordinates struct {
	Lat int64  `json:"lat"`
	Long int64 `json:"long"`
}
