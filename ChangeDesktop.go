package ritsecMentorProject

import (
	"golang.org/x/sys/windows/registry"
	"log"
)

// Change the user's desktop background and shortcuts

//TODO: turn off slideshow mode because it doesnt work if the desktop background is a slideshow
func changeDesktop() {
	key, e := registry.OpenKey(registry.CURRENT_USER, `Control Panel\Desktop`, registry.QUERY_VALUE)
	if e != nil {
		log.Fatal(e)
	}
	defer key.Close()
	//TODO: change to relative file path
	e = key.SetStringValue("Wallpaper", "C:/Users/austi/Pictures/f52w9odxown01.jpg")
	if e != nil {
		log.Fatal(e)
	}
}
