package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"os/exec"
)

var inputPath = "big_buck_bunny.mp4"
var outputPathBase = "sample-%d.jpg"
var outputRoot = "frames/"

// var frames = []string{"01:45", "01:50", "01:55", "2:00"}
// ffmpeg -ss 01:45 -i sample.mp4 -vframes 1 -q:v 2 sample-`date +%s`.jpg

var cmd = "ffmpeg -i sample.mp4 -r 25 -f image2pipe -c:v bmp -"
var args = []string{"-i", "sample.mp4", "-r", "25", "-f", "image2pipe", "-c:v", "bmp", "-"}

func main() {
	var cmd = exec.Command("ffmpeg", args...)
	cmd.Stderr = os.Stderr
	stdin, err := cmd.StdinPipe()
	if nil != err {
		log.Fatalf("Error obtaining stdin: %s", err.Error())
	}
	stdout, err := cmd.StdoutPipe()
	if nil != err {
		log.Fatalf("Error obtaining stdout: %s", err.Error())
	}
	reader := bufio.NewReader(stdout)
	go func(reader io.Reader) {
		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			log.Printf("Reading from subprocess: %s", scanner.Text())
			stdin.Write([]byte("some sample text\n"))
		}
	}(reader)

	if err := cmd.Start(); nil != err {
		log.Fatalf("Error starting program: %s, %s", cmd.Path, err.Error())
	}
	cmd.Wait()
}
