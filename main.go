package main

import (
	"Muth/config"
	"Muth/handlers"
	"Muth/handlers/message"
	"Muth/serve/player"
	"Muth/serve/tts"
	"github.com/lonelyevil/kook"
	"github.com/lonelyevil/kook/log_adapter/plog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Load config
	err := config.LoadConfig("config/config.yaml")
	if err != nil {
		panic(err)
		return
	}
	// Setup logger
	logger := config.Logger
	// Setup KOOK
	s := kook.New(config.Config.BotToken, plog.NewLogger(logger))

	player.MusicPlayer = player.NewPlayer()
	tts.TTS = tts.NewTTS()

	// Register KOOK handlers
	handlers.RegistryHandlers(s, message.MessageHan)

	// Start KOOK
	err = s.Open()
	if err != nil {
		panic(err)
		return
	}
	logger.Info().Msg("Bot is running")

	// Waiting for exit signal
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt, os.Kill, syscall.SIGTERM)
	<-sc
	logger.Info().Msg("Bot is shutting down")
	err = s.Close()
}
