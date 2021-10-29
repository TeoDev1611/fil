package utils

import "os"

// OpenFiles function     here is the util for open the files and not repeat code
func OpenFiles(file string) string {
	dat, err := os.ReadFile(file)
	CheckErrors(err)
	return string(dat)
}
