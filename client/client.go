package main

func main() {
	StartClient()
	// Start message sending loop
	go Listener()
	initializeMenu()
}
