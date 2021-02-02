//
//  ImageView.go
//  components
//
//  Created by d-exclaimation on 11:06 PM.
//  Copyright © 2020 d-exclaimation. All rights reserved.
//

package components

import (
	. "github.com/d-exclaimation/fb-messenger-api/models"
)

func ImageView(imageURL string) ReplyItem {
	return &ReplyAttachments{
		Attachment: ReplyAttachment{
			Type:    "template",
			Payload: TemplatePayload{
				TemplateType: "generic",
				Elements: []Elements{
					{
						Title:    "",
						Subtitle: "",
						ImageURL: imageURL,
						Buttons:  nil,
					},
				},
			},
		},
	}
}
