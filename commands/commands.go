//
//  commands.go
//  commands
//
//  Created by d-exclaimation on 1:45 PM.
//  Copyright © 2020 d-exclaimation. All rights reserved.
//

package commands

import "github.com/d-exclaimation/fb-messenger-api/models"

type Executable func(string, string) models.ReplyItem

func All() map[string]Executable {
	return map[string]Executable{
		"!8ball": EightBall,
		"!meme": Memes,
	}
}
