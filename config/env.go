//
//  env.go
//  config
//
//  Created by d-exclaimation on 10:21 PM.
//  Copyright © 2020 d-exclaimation. All rights reserved.
//

package config

import "os"

func GetPort() string {
	var port = os.Getenv("PORT")
	if len(port) < 4 {
		return ":1337"
	}
	return ":" + port
}
