package main

import (
	"os/exec"
)

func convert() error {
	cmd := exec.Command("./ffmpeg", "-i", "inp.mp3", "-c:a", "libmp3lame", "-b:a", "320k", "-map", "0:0", "-f", "segment", "-segment_time", "10", "-segment_list", "outputlist.m3u8", "-segment_format", "mpegts", "output%03d.ts")
	_, err := cmd.Output()

	if err != nil {
		return err
	}

	return nil
}
