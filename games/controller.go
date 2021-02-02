//
//  controller.rust
//  games
//
//  Created by d-exclaimation on 1:14 PM.
//  Copyright © 2020 d-exclaimation. All rights reserved.
//

package games

import (
	"github.com/d-exclaimation/fb-messenger-api/models"
	"strings"
)

func NewGame(senderPSID string, gameDocument map[string]*Sokoban) models.ReplyItem {
	// Create a new instance of the game
	gameDocument[senderPSID] = NewSokoban(10, 6)
	return CreateGameItem(gameDocument[senderPSID].ShowParts())
}

func MoveGame(senderPSID string, text string, gameDocument map[string]*Sokoban) models.ReplyItem {
	gameDocument[senderPSID].Move(strings.ToLower(text))
	return CreateGameItem(gameDocument[senderPSID].ShowParts())
}

func CreateGameItem(state [][]string, _ string, footer string, isWin bool) models.ReplyItem {
	var (
		header = "Sokoban"
	)
	if isWin {
		header = "Game Over"
	}
	return &models.ReplyMessage{
		Text: header + "\n" + rows(state) + "\n" + footer,
	}
}

func rows(content [][]string) string {
	var res []string
	for i := 0; i < len(content); i++ {
		var row = ""
		for j := 0; j < len(content[i]); j++ {
			row += content[i][j]
		}
		res = append(res, row)
	}
	return strings.Join(res, "\n")
}
