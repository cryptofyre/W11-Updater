package main

import (
	"bufio"
	"log"
	"os"

	"github.com/cavaliercoder/grab"
)

func main() {
	// App Start
	log.Printf("[Main] Application started.")

	// Inital Checking
	log.Printf("[FS] Checking for $WINDOWS.~BT directory.")

	folderInfo, err := os.Stat("C:/$WINDOWS.~BT")
	if os.IsNotExist(err) {
		log.Fatal("[FS] $WINDOWS.~BT does not exist. Try attempting an update then reopening the application.")
	}
	log.Println("[FS] $WINDOWS.~BT folder found with the following data ", folderInfo)

	// Permission Setting
	log.Printf("[FS] Setting permissions to the $WINDOWS.~BT directory for editing.")
	err = os.Chmod("C:/$WINDOWS.~BT/Sources", 0700)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("[FS] Finished setting permissions.")

	// Patching update package
	resp, err := grab.Get("C:/$WINDOWS.~BT/Sources/appraiserres.dll", "https://cryptofyre.org/cdn/appraiserres.dll")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("[Main] Successfully patched update file", resp.Filename)

	// Finished.
	log.Printf("You should now be able to install the update if you attempt to do so again.")
	log.Printf("(Press ENTER to close the application)")
	pause := bufio.NewScanner(os.Stdin)
	pause.Scan()
}

// told ya its super simple.
