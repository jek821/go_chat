package main

import ()

func main() {
	client := NewClient()
	// imagine rune time happens

	// then this gets called
	go client.RunListener()

}
