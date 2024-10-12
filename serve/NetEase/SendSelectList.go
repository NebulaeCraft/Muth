package NetEase

import (
	"MusicBot/config"
	"MusicBot/serve/player"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/lonelyevil/kook"
)

func SendSelectList(ctx *kook.KmarkdownMessageContext, musicsList *[]player.Music) {
	logger := config.Logger
	cardMsg := kook.CardMessageCard{
		Theme: kook.CardThemeSuccess,
		Size:  kook.CardSizeLg,
	}
	for _, music := range *musicsList {
		ID, _ := strconv.Atoi(music.ID)
		musicInfo, err := QueryMusicInfo(ID)
		if err != nil {
			logger.Error().Err(err).Msg("Query music info failed")
			return
		}
		section := kook.CardMessageSection{
			Mode: kook.CardMessageSectionModeRight,
			Text: kook.CardMessageElementKMarkdown{
				Content: "**歌曲：** " + musicInfo.Name + "\n**歌手：** " + strings.Join(musicInfo.Artists, ", ") + "\n**时长：** " + time.Duration(musicInfo.LastTime*1000000).String(),
			},
		}
		cardMsg.AddModule(section.SetAccessory(&kook.CardMessageElementButton{
			Theme: kook.CardThemePrimary,
			Text:  "点歌",
			Click: string(kook.CardMessageElementButtonClickReturnVal),
			Value: "NS" + music.ID,
		}))
	}
	cardMsgCtx, err := cardMsg.MarshalJSON()
	if err != nil {
		logger.Error().Err(err).Msg("Marshal card message failed")
		return
	}
	cardMsgCtxStr := fmt.Sprintf("[%s]", cardMsgCtx)
	resp, _ := ctx.Session.MessageCreate(&kook.MessageCreate{
		MessageCreateBase: kook.MessageCreateBase{
			TargetID: ctx.Common.TargetID,
			Content:  cardMsgCtxStr,
			Type:     kook.MessageTypeCard,
		},
	})

	go func() {
		time.Sleep(5 * time.Second)
		_ = ctx.Session.MessageDelete(resp.MsgID)
	}()
}
