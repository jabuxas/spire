package main

import (
	"fmt"
	"os"
	"time"
)

const (
	TMPDIR  = "/tmp/spire"
	PAYLOAD = TMPDIR + "/payload.json"
)

var GAME_LOCATION = os.Getenv("GAME_PATH")

func main() {
	getCache()
}

func getCache() error {
	info, err := os.Stat(PAYLOAD)

	if err != nil {
		fmt.Println("payload isn't accessible, requesting")
		DownloadCache()
	} else if !(info.ModTime().After(time.Now().Add(-2 * time.Hour))) {
		// download again if payload is older than 2 hours
		DownloadCache()
	} else {
		fmt.Println("payload was updated recently, won't update")
	}

	return nil
}

func installBepinex() error {
	_, err := os.Stat(GAME_LOCATION + "/BepInEx")

	if err != nil {
		// install bepinex

	} else {
		// bepinex already installed
	}

	return nil
}
