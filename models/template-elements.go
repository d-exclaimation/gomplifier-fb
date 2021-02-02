//
//  template-elements.go
//  models
//
//  Created by d-exclaimation on 2:49 PM.
//  Copyright © 2020 d-exclaimation. All rights reserved.
//

package models

type TemplateElements interface {
	ConformToTemplateElement()
}

// Generic Template Element
type GenericElement struct {
	Title    string    `json:"title"`
	Subtitle string    `json:"subtitle"`
	ImageURL string    `json:"image_url"`
	DefaultAction *URLButton `json:"default_action"`
	Buttons  []*Buttons `json:"buttons"`
}

func (*GenericElement) ConformToTemplateElement() {}

func NewGenericElement(title string, subtitle string, buttons ...*Buttons) *GenericElement {
	return &GenericElement{
		Title:         title,
		Subtitle:      subtitle,
		Buttons:       buttons,
	}
}

func NewGenericImageElement(title string, subtitle string, imageURL string, defaultAction *URLButton, buttons ...*Buttons) *GenericElement {
	return &GenericElement{
		Title:         title,
		Subtitle:      subtitle,
		ImageURL:      imageURL,
		DefaultAction: defaultAction,
		Buttons:       buttons,
	}
}


// Media Template Element
type MediaElement struct {
	MediaType MediaType `json:"media_type"`
	URL string `json:"url"`
	Buttons []*Buttons `json:"buttons"`
}

type MediaType string

const (
	VideoMedia MediaType = "video"
	ImageMedia MediaType = "image"
)

func (*MediaElement) ConformToTemplateElement() {}

func NewVideoElement(url string, buttons ...*Buttons) *MediaElement {
	return &MediaElement{
		MediaType: VideoMedia,
		URL:       url,
		Buttons:   buttons,
	}
}

func NewImageElement(url string, buttons ...*Buttons) *MediaElement {
	return &MediaElement{
		MediaType: ImageMedia,
		URL:       url,
		Buttons:   buttons,
	}
}
