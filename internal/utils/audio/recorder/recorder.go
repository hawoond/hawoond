package recorder

import (
	"errors"
	"runtime"

	"github.com/hawoond/hawoond/internal/utils/audio/recorder/os/linux"
	"github.com/hawoond/hawoond/internal/utils/audio/recorder/os/mac"
	"github.com/hawoond/hawoond/internal/utils/audio/recorder/os/win"
)

type Recorder struct {
	RecorderInf
}

type RecorderInf interface {
	StartRecording(filename string) error
	StopRecording() error
}

func NewRecorder() (RecorderInf, error) {
	switch runtime.GOOS {
	case "windows":
		return &win.WindowsRecorder{}, nil
	case "darwin":
		return &mac.MacRecorder{}, nil
	case "linux":
		return &linux.LinuxRecorder{}, nil
	default:
		return nil, errors.New("unsupported platform")
	}
}

// func (re Recorder) InitRecorder() {
// 	recorder, err := NewRecorder()
// 	if err != nil {
// 		fmt.Printf("Error: %v\n", err)
// 		return
// 	}
// 	re.RecorderInf = recorder
// }
