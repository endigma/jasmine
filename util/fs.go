package util

import (
	"io/ioutil"
	"os"
	"strings"
)

// FileExists reports whether the named file or directory exists.
func FileExists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func ReadFileContent(path string) (string, error) {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return strings.Trim(string(f), "\n "), nil
}
