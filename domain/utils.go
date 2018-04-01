package domain

import "log"

func LogIfPresent(error error) {
	if error != nil {
		log.Println(error)
	}
}
