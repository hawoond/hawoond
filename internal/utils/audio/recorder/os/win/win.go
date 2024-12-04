//go:build windows
// +build windows

package win

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/hawoond/hawoond/internal/utils/audio/recorder"
)

var (
	winmm                   = syscall.NewLazyDLL("winmm.dll")
	procWaveInOpen          = winmm.NewProc("waveInOpen")
	procWaveInPrepareHeader = winmm.NewProc("waveInPrepareHeader")
	procWaveInAddBuffer     = winmm.NewProc("waveInAddBuffer")
	procWaveInStart         = winmm.NewProc("waveInStart")
	procWaveInStop          = winmm.NewProc("waveInStop")
	procWaveInClose         = winmm.NewProc("waveInClose")
)

const (
	WAVE_MAPPER     = -1
	WAVE_FORMAT_PCM = 1
)

type WAVEFORMATEX struct {
	FormatTag      uint16
	Channels       uint16
	SamplesPerSec  uint32
	AvgBytesPerSec uint32
	BlockAlign     uint16
	BitsPerSample  uint16
	Size           uint16
}

type WAVEHDR struct {
	lpData          uintptr
	dwBufferLength  uint32
	dwBytesRecorded uint32
	dwUser          uintptr
	dwFlags         uint32
	dwLoops         uint32
	lpNext          uintptr
	reserved        uint32
}
type WindowsRecorder struct {
	hWaveIn syscall.Handle
	buffer  []byte
}

func NewRecorder() recorder.Recorder {
	return &WindowsRecorder{}
}

func (r *WindowsRecorder) StartRecording(filename string) error {
	fmt.Println("Starting recording")

	var format WAVEFORMATEX
	format.FormatTag = WAVE_FORMAT_PCM
	format.Channels = 1
	format.SamplesPerSec = 44100
	format.BitsPerSample = 16
	format.BlockAlign = format.Channels * format.BitsPerSample / 8
	format.AvgBytesPerSec = format.SamplesPerSec * uint32(format.BlockAlign)
	waveMapper := int32(WAVE_MAPPER)
	result, _, _ := procWaveInOpen.Call(uintptr(unsafe.Pointer(&r.hWaveIn)), uintptr(waveMapper), uintptr(unsafe.Pointer(&format)), 0, 0, 0)
	if result != 0 {
		return fmt.Errorf("failed to open waveIn device: %d", result)
	}

	bufferLength := 44100 * 2
	r.buffer = make([]byte, bufferLength)
	var header WAVEHDR
	header.lpData = uintptr(unsafe.Pointer(&r.buffer[0]))
	header.dwBufferLength = uint32(len(r.buffer))

	result, _, _ = procWaveInPrepareHeader.Call(uintptr(r.hWaveIn), uintptr(unsafe.Pointer(&header)), unsafe.Sizeof(header))
	if result != 0 {
		return fmt.Errorf("failed to prepare header: %d", result)
	}

	result, _, _ = procWaveInAddBuffer.Call(uintptr(r.hWaveIn), uintptr(unsafe.Pointer(&header)), unsafe.Sizeof(header))
	if result != 0 {
		return fmt.Errorf("failed to add buffer: %d", result)
	}

	result, _, _ = procWaveInStart.Call(uintptr(r.hWaveIn))
	if result != 0 {
		return fmt.Errorf("failed to start recording: %d", result)
	}

	return nil
}

func (r *WindowsRecorder) StopRecording() error {
	fmt.Println("Stopping recording")

	result, _, _ := procWaveInStop.Call(uintptr(r.hWaveIn))
	if result != 0 {
		return fmt.Errorf("failed to stop recording: %d", result)
	}

	result, _, _ = procWaveInClose.Call(uintptr(r.hWaveIn))
	if result != 0 {
		return fmt.Errorf("failed to close waveIn device: %d", result)
	}

	return nil
}
