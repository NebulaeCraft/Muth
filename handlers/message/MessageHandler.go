package message

import (
	"Muth/config"
	"Muth/serve/player"
	"Muth/serve/tts"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/lonelyevil/kook"
)

func IsInTTSChannel(channelID string) (bool, error) {
	id, err := strconv.ParseInt(channelID, 10, 64)
	if err != nil {
		return false, err
	}
	for _, v := range config.Config.TextChannel {
		if v.ID == int(id) {
			return true, nil
		}
	}
	return false, nil
}

func MessageHan(ctx *kook.KmarkdownMessageContext) {
	logger := config.Logger
	if ctx.Common.Type != kook.MessageTypeKMarkdown || ctx.Extra.Author.Bot {
		return
	}
	logger.Info().Msg("Message received: " + ctx.Common.Content)
	player.MusicPlayer.SetCtx(ctx)

	isInTTSChannel, err := IsInTTSChannel(ctx.Common.TargetID)
	if err != nil {
		logger.Error().Err(err).Msg("IsInTTSChannel")
	}

	if !isInTTSChannel {
		return
	}

	if strings.HasPrefix(ctx.Common.Content, "ping") {
		// Ping
		ctx.Common.Content = strings.TrimPrefix(ctx.Common.Content, "ping")
		_, _ = ctx.Session.MessageCreate(&kook.MessageCreate{
			MessageCreateBase: kook.MessageCreateBase{
				TargetID: ctx.Common.TargetID,
				Content:  "pong",
				Quote:    ctx.Common.MsgID,
			},
		})
	} else if strings.HasPrefix(ctx.Common.Content, "/tts v ") {
		// Change volume
		ctx.Common.Content = strings.TrimPrefix(ctx.Common.Content, "/tts v ")
		ChangeVolumeMessageHandler(ctx)
	} else if strings.HasPrefix(ctx.Common.Content, "/tts c ") {
		// Change channel
		ctx.Common.Content = strings.TrimPrefix(ctx.Common.Content, "/tts c ")
		ChangeChannelMessageHandler(ctx)
	} else {
		TTSMessageHandler(ctx)
	}
}

func ChangeVolumeMessageHandler(ctx *kook.KmarkdownMessageContext) {
	logger := config.Logger
	if ctx.Common.Content == "now" {
		player.MusicPlayer.SendMsg(fmt.Sprintf("当前音量: %d", player.MusicPlayer.Volume))
		return
	}
	volume, err := strconv.ParseInt(ctx.Common.Content, 10, 64)
	if err != nil {
		logger.Error().Err(err).Msg("Parse volume failed")
		player.MusicPlayer.SendMsg("解析音量失败：输入不合法")
		return
	}
	player.MusicPlayer.SetVolume(int(volume))
	player.MusicPlayer.SendMsg(fmt.Sprintf("音量已调整为 %ddB", volume))
}

func ChangeChannelMessageHandler(ctx *kook.KmarkdownMessageContext) {
	logger := config.Logger
	if ctx.Common.Content == "list" {
		player.MusicPlayer.SendMsg(fmt.Sprintf("可用频道列表: %s", config.ListChannel()))
		return
	}
	channel, err := config.FindChannelID(ctx.Common.Content)
	if err != nil {
		logger.Error().Err(err).Msg("Find channel failed")
		player.MusicPlayer.SendMsg("解析频道失败：频道不存在")
		return
	}
	player.MusicPlayer.SetChannel(channel)
	player.MusicPlayer.SendMsg(fmt.Sprintf("频道已切换为 %s", ctx.Common.Content))
}

func generateIdentifier(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)
	identifier := make([]byte, length)
	for i := range identifier {
		identifier[i] = charset[r.Intn(len(charset))]
	}
	return string(identifier)
}

func TTSMessageHandler(ctx *kook.KmarkdownMessageContext) {
	logger := config.Logger

	if len(ctx.Common.Content) > 300 {
		player.MusicPlayer.SendMsg("消息内容过长！")
		return
	}

	// Get user nickname
	nickname, err := ctx.Session.UserView(ctx.Common.AuthorID)
	if err != nil {
		logger.Error().Err(err).Str("AuthorID", ctx.Common.AuthorID).Msg("GetUserView failed")
		player.MusicPlayer.SendMsg("无法获取用户信息，AuthorID: " + ctx.Common.AuthorID)
	}
	logger.Info().Msgf("TTS Message: %s, Author: %s", ctx.Common.Content, nickname.Nickname)

	// Synthesize audio
	text := nickname.Nickname + "说：" + ctx.Common.Content
	path := "./assets/voice/" + generateIdentifier(6) + ".mp3"
	path, _ = tts.TTS.Speak(text, path, "zh-CN-YunxiNeural")

	music := &player.Music{
		File: path,
	}

	// Add to player
	player.MusicPlayer.AddMusic(music)
}
