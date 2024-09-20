package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func DownloadCache() error {
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

func DownloadBepinex() error {

}
