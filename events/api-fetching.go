//
//  api-fetching.go
//  events
//
//  Created by d-exclaimation on 8:10 PM.
//  Copyright © 2020 d-exclaimation. All rights reserved.
//

package events

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func sendMessageAPI(bodyObject interface{}) {
	var body, err = json.Marshal(bodyObject)
	if err != nil {
		log.Println(err)
		return
	}
	var (
		responseBody = bytes.NewBuffer(body)
		uri = fmt.Sprintf("https://graph.facebook.com/v9.0/me/messages?access_token=%s", os.Getenv("PAGE_ACCESS_TOKEN"))
	)

	resp, err := http.Post(uri, "application/json", responseBody)
	if err != nil {
		log.Println(err)
	}

	_ = resp.Body.Close()
}

func postToAPI(bodyObject interface{}, url string) {
	var body, err = json.Marshal(bodyObject)
	if err != nil {
		log.Println(err)
		return
	}
	var (
		responseBody = bytes.NewBuffer(body)
		resp, postErr = http.Post(url, "application/json", responseBody)
	)
	if postErr != nil {
		log.Println(postErr)
	}
	_ = resp.Body.Close()
}
