//
//  memes.go
//  commands
//
//  Created by d-exclaimation on 2:14 PM.
//  Copyright © 2020 d-exclaimation. All rights reserved.
//

package commands

import (
	"github.com/d-exclaimation/fb-messenger-api/models"
	"math/rand"
	"strings"
	"time"
)

func Memes(_ string, text string) models.ReplyItem {
	var params = strings.Split(text, " ") // Get all the commands as slices


	// Image dictionaries
	var images = map[string]string{
		"https://i.imgur.com/lfUAvZw.png": "https://i.imgur.com/KgfNw0E.png",
		"https://i.imgur.com/euvFB6k.jpg": "https://i.imgur.com/euvFB6k.jpg",
		"https://i.imgur.com/f6pPlXJ.jpg": "https://i.imgur.com/f6pPlXJ.jpg",
		"https://i.imgur.com/7wQWk3N.jpeg": "https://i.imgur.com/7wQWk3N.jpeg",
		"https://i.imgur.com/tUlWWxn.jpeg": "https://i.imgur.com/tUlWWxn.jpeg",
		"https://i.imgur.com/S48ghFN.png": "https://i.imgur.com/S48ghFN.png",
	}

	// Get the link to an image
	var res string

	rand.Seed(time.Now().UnixNano())

	// If image is given by parameter, just use that. Otherwise, random from map
	if len(params) > 1 {
		res = params[1]
	} else {
		var full = mapToSlice(images)
		var i = rand.Intn(len(full))
		res = full[i]
	}

	return models.NewImage(res)
}


func mapToSlice(dict map[string]string) []string {
	var res = make([]string, 0)
	for key := range dict {
		res = append(res, key)
	}
	return res
}