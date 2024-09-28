package main

import (
	"log"
	"os"
	"time"
)

const (
	TMPDIR  = "/tmp/spire"
	PAYLOAD = TMPDIR + "/payload.json"
)

var GAME_PATH = os.Getenv("GAME_PATH")

func main() {
	getCache()
	if err := SanitizeInput(&GAME_PATH); err != nil {
		log.Fatal("GAME_PATH is not set")
	}
	DownloadBepinex()
}

func getCache() error {
	info, err := os.Stat(PAYLOAD)

	if err != nil {
		// payload isn't accessible, requesting
		DownloadCache()
	} else if !(info.ModTime().After(time.Now().Add(-2 * time.Hour))) {
		// download again if payload is older than 2 hours
		DownloadCache()
	}
	// else payload was updated recently, won't update

	return nil
}
