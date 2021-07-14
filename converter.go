package main

import (
	"fmt"
	"os/exec"
)

func convertFile(jobID string) {
	audioPath := fmt.Sprintf("./%s/inp.mp3", jobID)
	m3u8 := fmt.Sprintf("./%s/outputlist.m3u8", jobID)
	ts := "./" + jobID + "/output%03d.ts"
	cmd := exec.Command("./ffmpeg", "-i", audioPath, "-c:a", "libmp3lame", "-b:a", "320k", "-map", "0:0", "-f", "segment", "-segment_time", "10", "-segment_list", m3u8, "-segment_format", "mpegts", ts)
	_, err := cmd.Output()

	if err != nil {
		j := jobs[jobID]
		passToChannel(&j, "failed converting")
		killSig(&j)
		return
	}

	j := jobs[jobID]
	passToChannel(&j, "converted")
}
