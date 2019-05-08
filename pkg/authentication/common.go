/**
 * Copyright (c) 2019 Adrian P.K. <apk@kuguar.io>
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */

package authentication

import (
	"log"
	"os"

	c "gitlab.com/mikrowezel/backend/granica/internal/config"
)

func checkError(err error, msg ...string) {
	if err != nil {
		if len(msg) > 0 && msg[0] != "" {
			log.Println("level", c.LogLevel.Error, "message", msg[0])
		}
		log.Println("level", c.LogLevel.Error, "message", err.Error())
		os.Exit(1)
	}
}
