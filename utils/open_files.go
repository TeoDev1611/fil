package utils

import "os"

func OpenFiles(file string) string {
	dat, err := os.ReadFile(file)
	CheckErrors(err)
	return string(dat)
}
