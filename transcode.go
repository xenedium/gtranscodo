package main

import (
	"log"
	"time"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func TranscodeToHevcGeneric(mediaPath string, destinationPath string, hardwareAccel bool, hardwareDevice string) error {
	if hardwareAccel {
		return TranscodeToHevcHaccel(mediaPath, destinationPath, hardwareDevice)
	}
	return TranscodeToHevcCpu(mediaPath, destinationPath)
}

func TranscodeToHevcHaccel(mediaPath string, destinationPath string, hardwareDevice string) error {
	log.Printf("[trans-haccel] - Transcoding %s to %s", mediaPath, destinationPath)
	err := ffmpeg.
		Input(mediaPath).
		Output(destinationPath, ffmpeg.KwArgs{
			"c:v:0":        "hevc_vaapi",
			"vaapi_device": hardwareDevice,
			"vf":           "format=nv12,hwupload",
			"movflags":     "use_metadata_tags",
		}).OverWriteOutput().Silent(true).
		Run()
	if err == nil {
		log.Printf("[trans-haccel] - Transcoded %s to %s", mediaPath, destinationPath)
		return nil
	}
	log.Printf("[trans-haccel] - Error transcoding %s to %s: %s", mediaPath, destinationPath, err)

	// Re add the file to the queue if transcoding fails after 5 minutes
	// Actively moving files may cause issues with the transcoding process
	// This is a temporary solution
	var retryCount int = 0
	if retryCount <= 10 {
		go func() {
			time.Sleep(300 * time.Second)
			retryCount++
			TranscodeToHevcHaccel(mediaPath, destinationPath, hardwareDevice)
		}()
	}
	return err
}

func TranscodeToHevcCpu(mediaPath string, destinationPath string) error {
	log.Printf("[trans-cpu] - Transcoding %s to %s", mediaPath, destinationPath)
	err := ffmpeg.
		Input(mediaPath).
		Output(destinationPath, ffmpeg.KwArgs{
			"c:v":      "libx265",
			"movflags": "use_metadata_tags",
		}).OverWriteOutput().Silent(true).
		Run()
	if err == nil {
		log.Printf("[trans-cpu] - Transcoded %s to %s", mediaPath, destinationPath)
		return nil
	}
	log.Printf("[trans-cpu] - Error transcoding %s to %s: %s", mediaPath, destinationPath, err)

	var retryCount int = 0
	if retryCount <= 10 {
		go func() {
			time.Sleep(300 * time.Second)
			retryCount++
			TranscodeToHevcCpu(mediaPath, destinationPath)
		}()
	}

	return err
}
