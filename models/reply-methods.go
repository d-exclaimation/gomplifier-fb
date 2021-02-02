//
//  reply-methods.go
//  models
//
//  Created by d-exclaimation on 4:30 PM.
//  Copyright © 2020 d-exclaimation. All rights reserved.
//

package models

// Template Attachments
func NewCustomTemplate(payload TemplatePayload) *ReplyAttachments {
	return &ReplyAttachments{
		Attachment: NewTemplateAttachment(payload),
	}
}

// Generics
func NewHGeneric(title string, subtitle string, buttons ...*Buttons) *ReplyAttachments {
	var genericPayload = NewGenericHPayload(NewGenericElement(title, subtitle, buttons...))
	return &ReplyAttachments{
		Attachment: NewTemplateAttachment(genericPayload),
	}
}

func NewSGeneric(title string, subtitle string, buttons ...*Buttons) *ReplyAttachments {
	var genericPayload = NewGenericSPayload(NewGenericElement(title, subtitle, buttons...))
	return &ReplyAttachments{
		Attachment: NewTemplateAttachment(genericPayload),
	}
}

func NewHGenericImage(title string, subtitle string, imageURL string, defaultAction *URLButton, buttons ...*Buttons) *ReplyAttachments {
	var genericPayload = NewGenericHPayload(NewGenericImageElement(title, subtitle, imageURL, defaultAction, buttons...))
	return &ReplyAttachments{
		Attachment: NewTemplateAttachment(genericPayload),
	}
}

func NewSGenericImage(title string, subtitle string, imageURL string, defaultAction *URLButton, buttons ...*Buttons) *ReplyAttachments {
	var genericPayload = NewGenericSPayload(NewGenericImageElement(title, subtitle, imageURL, defaultAction, buttons...))
	return &ReplyAttachments{
		Attachment: NewTemplateAttachment(genericPayload),
	}
}

// Buttons
func NewPrompt(text string, buttons ...*Buttons) *ReplyAttachments {
	var buttonPayload = NewButtonPayload(text, buttons...)
	return &ReplyAttachments{
		Attachment: NewTemplateAttachment(buttonPayload),
	}
}

func NewYesNoPrompt(text string, yes MyPayload, no MyPayload) *ReplyAttachments {
	var (
		yesButton = &Buttons{
			Type:    "postback",
			Title:   "Yes",
			Payload: yes,
		}
		noButton = &Buttons{
			Type:    "postback",
			Title:   "No",
			Payload: no,
		}
		buttonPayload = NewButtonPayload(text, yesButton, noButton)
	)
	return &ReplyAttachments{
		Attachment: NewTemplateAttachment(buttonPayload),
	}
}

// Medias
func NewVideo(url string, buttons ...*Buttons) *ReplyAttachments {
	var mediaPayload = NewMediaPayload(true, NewVideoElement(url, buttons...))
	return &ReplyAttachments{
		Attachment: NewTemplateAttachment(mediaPayload),
	}
}

func NewImage(url string, buttons ...*Buttons) *ReplyAttachments {
	var mediaPayload = NewMediaPayload(true, NewImageElement(url, buttons...))
	return &ReplyAttachments{
		Attachment: NewTemplateAttachment(mediaPayload),
	}
}
