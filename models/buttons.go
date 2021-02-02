//
//  buttons.go
//  models
//
//  Created by d-exclaimation on 2:59 PM.
//  Copyright © 2020 d-exclaimation. All rights reserved.
//

package models

type Buttons struct {
	Type    string `json:"type"`
	Title   string `json:"title"`
	Payload MyPayload `json:"payload"`
}

type WebUrl string

const WebURLType WebUrl = "web_url"

type URLButton struct {
	Type WebUrl `json:"type"`
	URL string `json:"url"`
	Title string `json:"title"`
}
