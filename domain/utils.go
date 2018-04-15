package domain

import "log"

func LogIfPresent(error error) {
	if error != nil {
		log.Println(error)
	}
}

func LogWithMessageIfPresent(message string, error error) {
	if error != nil {
		log.Printf(message+": %v", error)
	}
}
