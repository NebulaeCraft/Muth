package button

import (
	"MusicBot/config"
	"MusicBot/serve/NetEase"
	"MusicBot/serve/player"
	"fmt"
	"strconv"
	"strings"

	"github.com/lonelyevil/kook"
)

func ButtonHan(ctx *kook.MessageButtonClickContext) {
	config.Logger.Info().Msg("Button received: " + ctx.Extra.Value)
	if strings.HasPrefix(ctx.Extra.Value, "NS") {
		ctx.Extra.Value = strings.TrimPrefix(ctx.Extra.Value, "NS")
		NetEaseSearchButtonHan(ctx)
	} else if strings.HasPrefix(ctx.Extra.Value, "DEL") {
		ctx.Extra.Value = strings.TrimPrefix(ctx.Extra.Value, "DEL")
		DeleteMusicButtonHan(ctx)
	} else if ctx.Extra.Value == "CONFIRM" {
		player.MusicPlayer.SendMsg("你知道个🔨")
	}
}

func NetEaseSearchButtonHan(ctx *kook.MessageButtonClickContext) {
	logger := config.Logger
	id, _ := strconv.ParseInt(ctx.Extra.Value, 10, 64)
	musicResult, err := NetEase.QueryMusic(int(id))
	if err != nil {
		logger.Error().Err(err).Msg("Query music failed")
		player.MusicPlayer.SendMsg("查询音乐失败")
		return
	}
	player.MusicPlayer.AddMusic(musicResult)
}

func DeleteMusicButtonHan(ctx *kook.MessageButtonClickContext) {
	logger := config.Logger
	name := player.MusicPlayer.RemoveMusic(ctx.Extra.Value)
	if name == "" {
		logger.Error().Msg("Delete music failed")
		player.MusicPlayer.SendMsg("删除音乐失败")
		return
	}
	player.MusicPlayer.SendMsg(fmt.Sprintf("已删除音乐 %s", name))
}
