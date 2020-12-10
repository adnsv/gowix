package wix

import (
	"io/ioutil"
	"os"
)

// CheckFileHasContent returns true if the specified
// file exists and has content that matches buf.
func CheckFileHasContent(fn string, buf []byte) bool {
	if _, err := os.Stat(fn); os.IsNotExist(err) {
		return false
	}
	oldBytes, err := ioutil.ReadFile(fn)
	if err != nil || len(oldBytes) != len(buf) {
		return false
	}
	for i, b := range oldBytes {
		if buf[i] != b {
			return false
		}
	}
	return true
}

// WriteFileIfChanged writes buf into a file.
// Does not overwrite if the file already has the specified content.
// Uses 0666 permission if overwriting is neccessary.
func WriteFileIfChanged(fn string, buf []byte) error {
	if CheckFileHasContent(fn, buf) {
		return nil
	}
	return ioutil.WriteFile(fn, buf, 0666)
}
