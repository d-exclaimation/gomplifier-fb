//
//  8ball.go
//  commands
//
//  Created by d-exclaimation on 1:49 PM.
//  Copyright © 2020 d-exclaimation. All rights reserved.
//

package commands

import (
	"github.com/d-exclaimation/fb-messenger-api/models"
	"math/rand"
	"time"
)

func EightBall(_ string, _ string) models.ReplyItem {
	var yesResponds = []string{"the oracle said yes", "God himself told me yes", "yes", "it's probably yes", "yes"}
	var noResponds = []string{"the oracle said no", "God said no", "No", "Probably no", "nope"}

	rand.Seed(time.Now().UnixNano())

	var respond = rand.Intn(5)
	var yesNo = rand.Intn(2)

	var reply = noResponds[respond]

	if yesNo == 0 {
		reply = yesResponds[respond]
	}

	return &models.ReplyMessage{
		Text: reply,
	}
}
