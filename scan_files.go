package main

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/tidwall/gjson"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func ScanFiles(dir string, files chan string, rate int) {
	cachedFiles := make([]string, 0)
	for {
		log.Printf("[scan] - Scanning %s", dir)
		filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if !info.IsDir() && CheckCodec(path) && !contains(cachedFiles, path) {
				cachedFiles = append(cachedFiles, path)
				files <- path
				log.Printf("[scan] - Found file %s", path)
			}
			return nil
		})
		log.Printf("[scan] - Sleeping for %d seconds", rate)
		time.Sleep(time.Duration(rate) * time.Second)
	}
}

func CheckCodec(path string) bool {
	a, err := ffmpeg.Probe(path)
	if err != nil {
		return false
	}

	if gjson.Get(a, "streams.0.codec_name").String() != "hevc" {
		return true
	}

	return false
}
