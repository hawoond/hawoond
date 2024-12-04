//go:build linux
// +build linux

package linux

/*
#cgo LDFLAGS: -lasound
#include <alsa/asoundlib.h>

typedef struct {
    snd_pcm_t *pcm_handle;
    snd_pcm_hw_params_t *params;
    unsigned int sample_rate;
    int dir;
    snd_pcm_uframes_t frames;
    char *buffer;
    int size;
} ALSAContext;

int StartRecording(ALSAContext *context, const char *filename) {
    // PCM 장치 열기 (기본 장치 사용)
    int rc = snd_pcm_open(&context->pcm_handle, "default", SND_PCM_STREAM_CAPTURE, 0);
    if (rc < 0) {
        fprintf(stderr, "unable to open pcm device: %s\n", snd_strerror(rc));
        return rc;
    }

    // 하드웨어 파라미터 설정
    snd_pcm_hw_params_alloca(&context->params);
    snd_pcm_hw_params_any(context->pcm_handle, context->params);
    snd_pcm_hw_params_set_access(context->pcm_handle, context->params, SND_PCM_ACCESS_RW_INTERLEAVED);
    snd_pcm_hw_params_set_format(context->pcm_handle, context->params, SND_PCM_FORMAT_S16_LE);
    snd_pcm_hw_params_set_channels(context->pcm_handle, context->params, 1); // Mono
    context->sample_rate = 44100;
    snd_pcm_hw_params_set_rate_near(context->pcm_handle, context->params, &context->sample_rate, &context->dir);
    snd_pcm_hw_params(context->pcm_handle, context->params);

    // 버퍼 크기 설정
    snd_pcm_hw_params_get_period_size(context->params, &context->frames, &context->dir);
    context->size = context->frames * 2; // 2바이트(16비트) * Mono
    context->buffer = (char *) malloc(context->size);

    // WAV 파일 열기
    FILE *file = fopen(filename, "wb");
    if (!file) {
        fprintf(stderr, "unable to open file %s\n", filename);
        return -1;
    }

    // 녹음 시작
    while (1) {
        rc = snd_pcm_readi(context->pcm_handle, context->buffer, context->frames);
        if (rc == -EPIPE) {
            snd_pcm_prepare(context->pcm_handle);
        } else if (rc < 0) {
            fprintf(stderr, "read from audio interface failed: %s\n", snd_strerror(rc));
        } else if (rc != (int)context->frames) {
            fprintf(stderr, "short read, read %d frames\n", rc);
        }

        // 녹음된 데이터를 파일에 저장
        fwrite(context->buffer, context->size, 1, file);
    }

    // 파일 닫기 및 리소스 해제
    fclose(file);
    snd_pcm_drain(context->pcm_handle);
    snd_pcm_close(context->pcm_handle);
    free(context->buffer);

    return 0;
}

int StopRecording(ALSAContext *context) {
    snd_pcm_drop(context->pcm_handle);
    snd_pcm_close(context->pcm_handle);
    free(context->buffer);
    return 0;
}
*/
import "C"
import (
	"fmt"
	"unsafe"

	"github.com/hawoond/hawoond/internal/utils/audio/recorder"
)

type LinuxRecorder struct {
	context C.ALSAContext
}

func NewRecorder() recorder.Recorder {
	return &LinuxRecorder{}
}

func (r *LinuxRecorder) StartRecording(filename string) error {
	fmt.Println("Starting recording on Linux...")

	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))

	status := C.StartRecording(&r.context, cFilename)
	if status != 0 {
		return fmt.Errorf("Failed to start recording: %d", status)
	}

	return nil
}

func (r *LinuxRecorder) StopRecording() error {
	fmt.Println("Stopping recording on Linux...")

	status := C.StopRecording(&r.context)
	if status != 0 {
		return fmt.Errorf("Failed to stop recording: %d", status)
	}

	return nil
}
