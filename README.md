# GtranscodO

## What is GtranscodO?

GtranscodO is a simple tool to transcode video files. It is based on the [FFmpeg](https://ffmpeg.org/) library.

I've created this tool to transcode my video files to a format that is supported by my TV,

to reduce the size of the files and to be able to stream them over the network, and to also avoind transcoding them on the fly.

## How to use GtranscodO?

GtranscodO is a command line tool. You can use it like this

```bash
    gtranscodo -help

      -cache-dir string
        	Directory to store cache files (default "/tmp")
      -hardware-device string
        	Hardware device if using hardware acceleration (default "/dev/dri/renderD128")
      -path string
        	Path to directory to scan
      -scan-rate int
        	Rate to scan directory in seconds (default 60)
      -use-hardware
        	Use hardware acceleration (default false)
```

## Notes

- GtranscodO is a work in progress. It is not yet ready for production use.
- GtranscodO is a personal project.
- GtranscodO uses the VA-API for hardware acceleration. (Tested on AMD rx 7600 on Arch Linux)
- Make sure you have FFmpeg and VA-API installed on your system.
