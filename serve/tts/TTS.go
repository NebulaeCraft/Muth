package tts

import (
	"Muth/config"
	"bytes"
	edgeTTS "github.com/xiaolibuzai-ovo/edge-tts/pkg/wrapper"
	"golang.org/x/net/context"
	"io"
	"os"
	"os/exec"
	"strings"
)

type TTSEngine struct {
	EdgeTTSEngine *edgeTTS.EdgeTTS
}

var TTS *TTSEngine

func (engine *TTSEngine) Speak(text, path, voice string) (string, error) {
	logger := config.Logger

	tts := edgeTTS.NewEdgeTTS()

	speech, err := tts.TextToSpeech(context.Background(), text, voice)
	if err != nil {
		return "", err
	}

	file, err := os.Create(path)
	if err != nil {
		return "", err
	}
	defer file.Close()
	io.Copy(file, speech)

	newPath := strings.Replace(path, ".mp3", "-s.mp3", -1)
	cmd := exec.Command("ffmpeg", "-f", "lavfi", "-t", config.Config.SilentDuration, "-i", "anullsrc=r=44100:cl=stereo", "-i", path, "-filter_complex", "[0][1]concat=n=2:v=0:a=1", "-y", newPath)
	var outputBuffer bytes.Buffer
	cmd.Stdout = &outputBuffer
	cmd.Stderr = &outputBuffer

	if err := cmd.Run(); err != nil {
		logger.Err(err).Msgf("Error when concating output: %s", outputBuffer.String())
	}

	return newPath, nil
}

func NewTTS() *TTSEngine {
	tts := new(TTSEngine)
	return tts
}
