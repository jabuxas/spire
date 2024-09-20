package main

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func DownloadCache() error {
	resp, err := http.Get("https://thunderstore.io/api/v1/package/")
	if err != nil {
		return fmt.Errorf("couldn't download payload: %w", err)
	}
	defer resp.Body.Close()

	if err = os.MkdirAll(TMPDIR, os.ModePerm); err != nil {
		return fmt.Errorf("couldn't create tmpdir: %w", err)
	}
	out, err := os.Create(PAYLOAD)
	if err != nil {
		return fmt.Errorf("couldn't create file: %w", err)
	}
	defer out.Close()

	if _, err = io.Copy(out, resp.Body); err != nil {
		return fmt.Errorf("couldn't write to file: %w", err)
	}

	return nil
}

func DownloadBepinex() error {
	tmpPath := TMPDIR + "/bepinex.zip"

	// download only if it doesnt exist
	if _, err := os.Stat(tmpPath); err != nil {
		resp, err := http.Get("https://github.com/BepInEx/BepInEx/releases/download/v5.4.23.2/BepInEx_win_x64_5.4.23.2.zip")
		if err != nil {
			return fmt.Errorf("could not download bepinex: %w", err)
		}
		defer resp.Body.Close()

		out, err := os.Create(tmpPath)
		if err != nil {
			return fmt.Errorf("could not create file: %w", err)
		}
		defer out.Close()

		_, err = io.Copy(out, resp.Body)
		if err != nil {
			return fmt.Errorf("could not save file: %w", err)
		}
	}

	err := unzip(tmpPath, GAME_PATH)
	if err != nil {
		return fmt.Errorf("could not extract BepInEx: %w", err)
	}

	return nil
}

func unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		fpath := filepath.Join(dest, f.Name)

		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}

		rc, err := f.Open()
		if err != nil {
			outFile.Close()
			return err
		}

		_, err = io.Copy(outFile, rc)

		outFile.Close()
		rc.Close()

		if err != nil {
			return err
		}
	}
	return nil
}

func SanitizeInput(input *string) error {
	if *input == "" {
		return fmt.Errorf("%s is not set", *input)
	}
	if strings.HasSuffix(*input, "/") {
		*input = strings.TrimSuffix(*input, "/")
	}
	return nil
}
