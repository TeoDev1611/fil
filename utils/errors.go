package utils

import "log"

// CheckErrors function    Here are the util for check the errors on the app
func CheckErrors(e error) {
	if e != nil {
		log.Fatalf("Error happened: %s", e)
	}
}
