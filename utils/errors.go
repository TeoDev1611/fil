package utils

import "log"

func CheckErrors(e error) {
	if e != nil {
		log.Fatalf("Error happened: %s", e)
	}
}
