package ffmpeg

import (
	"io"
)

// This could be an interface

type FFMPEG struct {
	// config
}

func NewFFMPEG(r io.Reader, w io.Writer) {}

func (f *FFMPEG) Start() {}
