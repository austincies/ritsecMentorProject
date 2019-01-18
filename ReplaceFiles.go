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
