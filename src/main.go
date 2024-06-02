package main

import (
  "fmt"
  "log"
  "os/exec"
)

const (
  VIDEOIN = "video-in/teaser0_06022024.mp4"
  OUTNAME = "video-out/teser0_comp.mp4"
  //DURATION = 34
  //TARGETMB = 20
)

func main() {
  //targetBitrate := (TARGETMB * 1024 * 1024 * 8) / DURATION
  ffmpegCmd := exec.Command(
    "ffmpeg",
    "-i", VIDEOIN,
    "-b:v", "4934k",
    "-c:a", "aac",
    "-b:a", "190k",
    "-ac", "2",
    "-t", "34",
    OUTNAME,
  )
  ffmpegOutput, err := ffmpegCmd.CombinedOutput()
  if err != nil {
    log.Fatalf("\nmain(): An error occured while running ffmpegCmd: %s\n%v", string(ffmpegOutput), err)
  }
  fmt.Printf("\nmain(): Successfully compressed video %s and saved as %s with ffmpeg\n\n%s", VIDEOIN, OUTNAME, string(ffmpegOutput))
}
