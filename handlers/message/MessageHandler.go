package message

import (
	"MusicBot/config"
	"MusicBot/serve/NetEase"
	"MusicBot/serve/player"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/lonelyevil/kook"
)

func MessageHan(ctx *kook.KmarkdownMessageContext) {
	logger := config.Logger
	if ctx.Common.Type != kook.MessageTypeKMarkdown || ctx.Extra.Author.Bot {
		return
	}
	logger.Info().Msg("Message received: " + ctx.Common.Content)
	player.MusicPlayer.SetCtx(ctx)
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
	} else if strings.HasPrefix(ctx.Common.Content, "/version") {
		// Version
		_, _ = ctx.Session.MessageCreate(&kook.MessageCreate{
			MessageCreateBase: kook.MessageCreateBase{
				TargetID: ctx.Common.TargetID,
				Content:  "golang 1.0",
				Quote:    ctx.Common.MsgID,
			},
		})
	} else if strings.HasPrefix(ctx.Common.Content, "/reload") {

	} else if strings.HasPrefix(ctx.Common.Content, "/v ") {
		// Change volume
		ctx.Common.Content = strings.TrimPrefix(ctx.Common.Content, "/v ")
		ChangeVolumeMessageHandler(ctx)
	} else if strings.HasPrefix(ctx.Common.Content, "/c ") {
		// Change channel
		ctx.Common.Content = strings.TrimPrefix(ctx.Common.Content, "/c ")
		ChangeChannelMessageHandler(ctx)
	} else if strings.HasPrefix(ctx.Common.Content, "/n ") || strings.HasPrefix(ctx.Common.Content, "/s ") || strings.HasPrefix(ctx.Common.Content, "/b ") || strings.HasPrefix(ctx.Common.Content, "/q ") {
		if strings.HasPrefix(ctx.Common.Content, "/n ") {
			// Netease Music
			ctx.Common.Content = strings.TrimPrefix(ctx.Common.Content, "/n ")
			NetEaseMusicMessageHandler(ctx)
		} else if strings.HasPrefix(ctx.Common.Content, "/s ") {
			// Netease Music Search
			ctx.Common.Content = strings.TrimPrefix(ctx.Common.Content, "/s ")
			NetEaseMusicSearchMessageHandler(ctx)
		}
	} else if ctx.Common.Content == "/skip" {
		// Skip
		SkipMusicMessageHandler(ctx)
	} else if ctx.Common.Content == "/list" {
		// List
		player.MusicPlayer.SendMusicList()
	}
}

func NetEaseMusicSearchMessageHandler(ctx *kook.KmarkdownMessageContext) {
	logger := config.Logger
	searchList, err := NetEase.SearchMusic(ctx.Common.Content)
	if err != nil {
		logger.Error().Err(err).Msg("Search music failed")
		player.MusicPlayer.SendMsg("搜索音乐失败")
		return
	}
	NetEase.SendSelectList(ctx, searchList)
}

func NetEaseMusicMessageHandler(ctx *kook.KmarkdownMessageContext) {
	logger := config.Logger

	re := regexp.MustCompile(`song\?id=(\d+)`)
	match := re.FindStringSubmatch(ctx.Common.Content)
	if len(match) > 1 {
		ctx.Common.Content = match[1]
	}

	id, err := strconv.ParseInt(ctx.Common.Content, 10, 64)
	if err != nil {
		logger.Error().Err(err).Msg("Parse music id failed")
		player.MusicPlayer.SendMsg("解析音乐ID失败：输入不合法")
		return
	}

	logger.Info().Msgf("Query music for id: %d", id)
	musicResult, err := NetEase.QueryMusic(int(id))
	if err != nil {
		logger.Error().Err(err).Msg("Query music failed")
		player.MusicPlayer.SendMsg("查询音乐失败")
		return
	}
	player.MusicPlayer.AddMusic(musicResult)
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

func SkipMusicMessageHandler(ctx *kook.KmarkdownMessageContext) {
	player.MusicPlayer.SkipMusic()
}
