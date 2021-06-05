package main

import (
	"fmt"
	"log"
	"os/exec"
)

func convert() {
	cmd := exec.Command("./ffmpeg", "-i", "inp.mp3", "-c:a", "libmp3lame", "-b:a", "320k", "-map", "0:0", "-f", "segment", "-segment_time", "10", "-segment_list", "outputlist.m3u8", "-segment_format", "mpegts", "output%03d.ts")
	stdout, err := cmd.Output()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(stdout))
}
