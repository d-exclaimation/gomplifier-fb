//
//  fb-postback.go
//  models
//
//  Created by d-exclaimation on 11:13 PM.
//  Copyright © 2020 d-exclaimation. All rights reserved.
//

package models

type PostBackRef struct {
	Ref    string `json:"ref"`
	Source string `json:"source"`
	Type   string `json:"type"`
}

type Postback struct {
	Title    string      `json:"title"`
	Payload  string      `json:"payload"`
	Referral PostBackRef `json:"referral"`
}

