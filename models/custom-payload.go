//
//  custom-payload.go
//  models
//
//  Created by d-exclaimation on 12:30 PM.
//  Copyright © 2020 d-exclaimation. All rights reserved.
//

package models

type MyPayload string

const (
	PayloadYes MyPayload = "yes"
	PayloadNo  MyPayload = "no"
)

func YesPostBack() *Buttons {
	return &Buttons{
		Type:    "postback",
		Title:   "Yes",
		Payload: PayloadYes,
	}
}

func NoPostBack() *Buttons {
	return &Buttons{
		Type:    "postback",
		Title:   "No",
		Payload: PayloadNo,
	}
}