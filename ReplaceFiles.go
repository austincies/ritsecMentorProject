package ritsecMentorProject

import "os"

// Structure to replace file with
type newFile struct {
	name     string  // New name for file
	icon     os.File // New icon for file if possible
	contents os.File // New file contents to store at file location if possible
}

// Change the name of a fully replaced file
func renameFile() {

}

// Make a web request to get the image
func getImage(url string) os.File {
	image := os.File{}
	return image
}

// Scorch the original file
func scorch(file os.File) {
	for i := 0; i < 7; i++ { //scorch the file 7 times

	}
}

// Replace the old file with the new image
func replace() {

}
