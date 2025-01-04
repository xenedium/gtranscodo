package main

import (
	"flag"
	"log"
	"path/filepath"
)

func main() {

	var useHardware bool
	var hardwareDevice string
	var path string
	var cacheDir string
	var scanRate int
	flag.BoolVar(&useHardware, "use-hardware", false, "Use hardware acceleration (default false)")
	flag.StringVar(&hardwareDevice, "hardware-device", "/dev/dri/renderD128", "Hardware device if using hardware acceleration")
	flag.StringVar(&path, "path", "", "Path to directory to scan")
	flag.StringVar(&cacheDir, "cache-dir", "/tmp", "Directory to store cache files")
	flag.IntVar(&scanRate, "scan-rate", 60, "Rate to scan directory in seconds")
	flag.Parse()

	if path == "" {
		flag.PrintDefaults()
		log.Fatal("Path is required")
	}

	var files = make(chan string)

	go ScanFiles(path, files, scanRate)

	go func() {
		for {
			select {
			case file := <-files:
				err := TranscodeToHevcGeneric(file, filepath.Join(cacheDir, filepath.Base(file)), useHardware, hardwareDevice)
				if err == nil {
					err = MoveFile(filepath.Join(cacheDir, filepath.Base(file)), file)
					if err != nil {
						log.Fatalf("[utils] - Error moving file %s: %s", file, err)
					} else {
						log.Printf("[utils] - Moved file %s", file)
					}
				}
			}
		}
	}()

	select {}
}
