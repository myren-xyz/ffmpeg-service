package main

import (
	"fmt"
	"os/exec"
)

func convertFile(jobID string, fileExt string) {
	audioPath := fmt.Sprintf("./%s/inp.%s", jobID, fileExt)
	m3u8 := fmt.Sprintf("./%s/outputlist.m3u8", jobID)
	ts := "./" + jobID + "/output%03d.ts"
	cmd := exec.Command("./ffmpeg", "-i", audioPath, "-c:a", "libmp3lame", "-b:a", "320k", "-map", "0:0", "-f", "segment", "-segment_time", "10", "-segment_list", m3u8, "-segment_format", "mpegts", ts)
	_, err := cmd.Output()

	if err != nil {
		jobs.RLock()
		j := jobs.store[jobID]
		jobs.RUnlock()
		passToChannel(&j, "failed converting")
		killSig(&j)
		return
	}

	jobs.RLock()
	j := jobs.store[jobID]
	jobs.RUnlock()
	passToChannel(&j, "converted")
}
