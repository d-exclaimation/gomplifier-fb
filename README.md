# gomplifier-fb
 A bot for Facebook Messanger Graph API

## Overview
This is another piece of a project I am building, the project is related to bots and servers as services.

This is the code to create a bot using [`Gin`](https://github.com/gin-gonic/gin) for Web Server and my implementation of Facebook API Library/Wrapper/Thing idk for Golang,

## Facebook Graph API for Messnger
The Facebook API for Messenger is a lot more involved since their API focuses on being language-agnostic and thus, not provide a single library / way to communicate back and forward

I'm using Gin to setup a simple web server to take the POST and GET Request Required, to be fair any http libraries would work. The difficult part is to parse the json given.

I'm taking the approach of creating full struct with all the optionals for each JSON and have `omitempty` with `pointer` create `nil` for value not given to deal with the dynamic nature of the JSON given back.

```go
type MessagingEvent struct {
	Sender    FbUser    `json:"sender"`
	Recipient FbUser    `json:"recipient"`
	Timestamp int64     `json:"timestamp"`
	Message   *Message  `json:"message,omitempty"`
	Postback  *Postback `json:"postback,omitempty"`
}
```
> This example of mine only ask for `messages` and `messages_postback` events. Facebook allows you to customize what event you want to listen to

For sending one back, I'm using empty implementation interfaces to differs one type from another with no custom Marshalling or Unmarshalling.

Note that a dynamic language like Javascript would probably be more easier to use. However, I worry that when things got complex, I might struggle with the dynamic nature

```go
// Setup HTTP Server for receiving requests from LINE platform
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
```

## Bot functionalities

### Commands:
> 1. Play Sokoban `!start => w, a, s, d`

I haven't added all the commands I wanted but since I already made a Go chatbot, the logic can be reimplemented relatively easily here as well. 

This is one of the reason I am using Go for this project instead of Javascript, Typescript or Rust

### Connecting to other bots

`Again, No info here yet, I am keeping this for myself atm sorry :)`


### Facebook API Library Wrapper for Go
Since I implemented my own structs to parse and send back JSON to the API, I am thinking of making that part of the code as its own libraries to be re-used by me or anyone else

Hopefully, this part of code will live as its own libraries which I will link here:
#### Facebook Messenger API Library / Wrapper / stuff
[`not yet`](https://github.com/d-exclaimation/gomplifier-fb)
