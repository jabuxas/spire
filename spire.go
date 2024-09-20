package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	TMPDIR  = "/tmp/spire"
	PAYLOAD = TMPDIR + "/payload.json"
)

func main() {
	get_cache()
}

func get_cache() error {
	info, err := os.Stat(PAYLOAD)

	if err != nil {
		fmt.Println("payload isn't accessible, requesting")
		download_cache()
	} else if !(info.ModTime().After(time.Now().Add(-2 * time.Hour))) {
		download_cache()
	} else {
		fmt.Println("payload was updated recently, won't update")
	}

	return nil
}

func download_cache() error {
	resp, err := http.Get("https://thunderstore.io/api/v1/package/")
	if err != nil {
		log.Fatal("couldn't download thunderstore payload")
		return err
	}
	defer resp.Body.Close()

	if err := os.MkdirAll(TMPDIR, os.ModePerm); err != nil {
		log.Fatal("couldn't create tmpdir")
		return err
	}
	out, err := os.Create(PAYLOAD)
	if err != nil {
		log.Fatal("couldn't create payload file")
		return err
	}
	defer out.Close()

	if _, err := io.Copy(out, resp.Body); err != nil {
		log.Fatal("couldn't write to payload file")
		return err
	}

	return nil
}
