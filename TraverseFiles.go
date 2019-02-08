package ritsecMentorProject

import "net/http"

//Recursively traverse filepath
//	Replace files when found
//	Replace folders when contents replaced?
func traverseFiles() {
	// TODO: traverse files
	//  open each directory encountered, replace each file encountered with ReplaceFiles.replaceFile()
}

// get all of the images from a remote repository (drive folder)
func getPics(urls []string) []*http.Response {
	var pics []*http.Response
	//TODO: iterate through list of urls getting each image with ReplaceFiles.getImage()
	// If nil returned, then skip that entry
	return pics
}
