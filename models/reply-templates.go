//
//  reply-templates.go
//  models
//
//  Created by d-exclaimation on 1:07 PM.
//  Copyright © 2020 d-exclaimation. All rights reserved.
//

package models

// Template AttachmentPayload
type TemplateType string

const (
	GenericTemplate TemplateType = "generic"
	ButtonTemplate  TemplateType = "button"
	MediaTemplate   TemplateType = "media"
)


// Interface to all types of Template Payload
type TemplatePayload interface {
	ConformToTemplatePayload()
}


// Generic Template
type GenericPayload struct {
	TemplateType TemplateType `json:"template_type"`
	ImageAspectRatio AspectRatio `json:"image_aspect_ratio"`
	Elements     []*GenericElement `json:"elements"`
}

type AspectRatio string

const (
	HorizontalAspectRatio AspectRatio = "horizontal"
	SquareAspectRatio 	  AspectRatio = "square"
)

func (*GenericPayload) ConformToTemplatePayload() {}

func NewGenericHPayload(elements ...*GenericElement) *GenericPayload {
	return &GenericPayload{
		TemplateType:     GenericTemplate,
		ImageAspectRatio: HorizontalAspectRatio,
		Elements:         elements,
	}
}

func NewGenericSPayload(elements ...*GenericElement) *GenericPayload {
	return &GenericPayload{
		TemplateType:     GenericTemplate,
		ImageAspectRatio: SquareAspectRatio,
		Elements:         elements,
	}
}


// Button Template
type ButtonPayload struct {
	TemplateType TemplateType `json:"template_type"`
	Text string `json:"text"`
	Buttons []*Buttons `json:"buttons"`
}

func (*ButtonPayload) ConformToTemplatePayload() {}

func NewButtonPayload(text string, buttons ...*Buttons) *ButtonPayload {
	return &ButtonPayload{
		TemplateType: ButtonTemplate,
		Text:         text,
		Buttons:      buttons,
	}
}


// Media Template
type MediaPayload struct {
	TemplateType TemplateType `json:"template_type"`
	Elements []*MediaElement `json:"elements"`
	Sharable bool `json:"sharable"`
}

func (*MediaPayload) ConformToTemplatePayload() {}

func NewMediaPayload(sharable bool, elements ...*MediaElement) *MediaPayload {
	return &MediaPayload{
		TemplateType: MediaTemplate,
		Elements:     elements,
		Sharable:     sharable,
	}
}
