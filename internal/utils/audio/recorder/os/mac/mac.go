//go:build darwin
// +build darwin

package mac

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework CoreAudio -framework AudioToolbox
#include <CoreAudio/CoreAudio.h>
#include <AudioToolbox/AudioQueue.h>
#include <stdlib.h>

typedef struct {
    AudioQueueRef queue;
    AudioQueueBufferRef buffer;
} AudioQueueContext;

void HandleInputBuffer(
    void *custom_data,
    AudioQueueRef inAQ,
    AudioQueueBufferRef inBuffer,
    const AudioTimeStamp *inStartTime,
    UInt32 inNumberPacketDescriptions,
    const AudioStreamPacketDescription *inPacketDescs
) {
    AudioQueueContext *context = (AudioQueueContext*)custom_data;
    fwrite(inBuffer->mAudioData, 1, inBuffer->mAudioDataByteSize, stdout);
    AudioQueueEnqueueBuffer(context->queue, context->buffer, 0, NULL);
}

OSStatus StartRecording(AudioQueueContext *context) {
    AudioStreamBasicDescription format;
    format.mSampleRate = 44100.0;
    format.mFormatID = kAudioFormatLinearPCM;
    format.mFormatFlags = kAudioFormatFlagIsSignedInteger | kAudioFormatFlagIsPacked;
    format.mBytesPerPacket = 2;
    format.mFramesPerPacket = 1;
    format.mBytesPerFrame = 2;
    format.mChannelsPerFrame = 1;
    format.mBitsPerChannel = 16;
    format.mReserved = 0;

    OSStatus status = AudioQueueNewInput(&format, HandleInputBuffer, context, NULL, kCFRunLoopCommonModes, 0, &context->queue);
    if (status != 0) {
        return status;
    }

    AudioQueueAllocateBuffer(context->queue, 4096, &context->buffer);
    AudioQueueEnqueueBuffer(context->queue, context->buffer, 0, NULL);

    status = AudioQueueStart(context->queue, NULL);
    return status;
}

OSStatus StopRecording(AudioQueueContext *context) {
    return AudioQueueStop(context->queue, true);
}
*/
import "C"
import (
	"fmt"

	"github.com/hawoond/hawoond/internal/utils/audio/recorder"
)

type MacRecorder struct {
	context C.AudioQueueContext
}

func NewRecorder() recorder.Recorder {
	return &MacRecorder{}
}

func (r *MacRecorder) StartRecording(filename string) error {
	fmt.Println("Starting recording")

	status := C.StartRecording(&r.context)
	if status != 0 {
		return fmt.Errorf("failed to start recording: %d", status)
	}

	return nil
}

func (r *MacRecorder) StopRecording() error {
	fmt.Println("Stopping recording")

	status := C.StopRecording(&r.context)
	if status != 0 {
		return fmt.Errorf("failed to stop recording: %d", status)
	}

	return nil
}
