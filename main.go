//
//  main.go
//  fb-messenger-api
//
//  Created by d-exclaimation on 8:50 PM.
//  Copyright © 2020 d-exclaimation. All rights reserved.
//

package main

import (
	"github.com/d-exclaimation/fb-messenger-api/config"
	"github.com/d-exclaimation/fb-messenger-api/events"
	"github.com/d-exclaimation/fb-messenger-api/games"
	"github.com/d-exclaimation/fb-messenger-api/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)


func main() {
	gin.SetMode(gin.ReleaseMode)
	var (
		app = gin.Default()
		gameDocument = make(map[string]*games.Sokoban)
	)

	app.POST("/webhook", func(context *gin.Context) {
		var body models.FbPostHook

		if err := context.BindJSON(&body); err != nil {
			log.Println(err)
			context.String(http.StatusOK, "Cannot parse JSON so lol")
			return
		}

		if body.Object == "page" {
			for i := 0; i < len(body.Entry); i++ {
				var webhookEvent = body.Entry[i].Messaging[0]
				if webhookEvent.Message != nil {

					// Handle message events
					events.HandleMessageEvent(webhookEvent.Sender.ID, webhookEvent.Message, gameDocument)

				} else if webhookEvent.Postback != nil {

					// Handle Postback events
					events.HandlePostBackEvent(webhookEvent.Sender.ID, webhookEvent.Postback)
				}
			}
			context.String(http.StatusOK, "")
			return
		}

		context.String(http.StatusNotFound, "Wrong item")
	})

	app.GET("/webhook", getHook)

	if err := app.Run(config.GetPort()); err != nil {
		log.Fatalln(err)
	}
}

func getHook(context *gin.Context) {
	var (
		verifyToken = os.Getenv("VERIFYTOKEN")
		mode = context.DefaultQuery("hub.mode", "nil")
		token = context.DefaultQuery("hub.verify_token", "nil")
		challenge = context.DefaultQuery("hub.challenge", "nil")
	)

	if mode != "nil" && token != "nil" {
		if mode == "subscribe" && token == verifyToken {
			log.Println("Verified")
			context.String(http.StatusOK, challenge)
		} else {
			context.String(http.StatusForbidden, "Wrong verify token")
		}
		return
	}
	context.String(http.StatusBadRequest, "Bad Request")
}