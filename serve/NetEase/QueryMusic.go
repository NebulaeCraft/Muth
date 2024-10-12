package NetEase

import (
	"MusicBot/config"
	"MusicBot/serve/player"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

type MusicResp struct {
	Data []struct {
		ID                 int         `json:"id"`
		URL                string      `json:"url"`
		Br                 int         `json:"br"`
		Size               int         `json:"size"`
		Md5                string      `json:"md5"`
		Code               int         `json:"code"`
		Expi               int         `json:"expi"`
		Type               string      `json:"type"`
		Gain               float64     `json:"gain"`
		Peak               float64     `json:"peak"`
		Fee                int         `json:"fee"`
		Uf                 interface{} `json:"uf"`
		Payed              int         `json:"payed"`
		Flag               int         `json:"flag"`
		CanExtend          bool        `json:"canExtend"`
		FreeTrialInfo      interface{} `json:"freeTrialInfo"`
		Level              string      `json:"level"`
		EncodeType         string      `json:"encodeType"`
		FreeTrialPrivilege struct {
			ResConsumable      bool        `json:"resConsumable"`
			UserConsumable     bool        `json:"userConsumable"`
			ListenType         interface{} `json:"listenType"`
			CannotListenReason interface{} `json:"cannotListenReason"`
		} `json:"freeTrialPrivilege"`
		FreeTimeTrialPrivilege struct {
			ResConsumable  bool `json:"resConsumable"`
			UserConsumable bool `json:"userConsumable"`
			Type           int  `json:"type"`
			RemainTime     int  `json:"remainTime"`
		} `json:"freeTimeTrialPrivilege"`
		URLSource   int         `json:"urlSource"`
		RightSource int         `json:"rightSource"`
		PodcastCtrp interface{} `json:"podcastCtrp"`
		EffectTypes interface{} `json:"effectTypes"`
		Time        int         `json:"time"`
	} `json:"data"`
	Code int `json:"code"`
}

type MusicReq struct {
	ID      int
	Quality string
}

type MusicInfoResp struct {
	Songs []struct {
		Name string `json:"name"`
		ID   int    `json:"id"`
		Pst  int    `json:"pst"`
		T    int    `json:"t"`
		Ar   []struct {
			ID    int           `json:"id"`
			Name  string        `json:"name"`
			Tns   []interface{} `json:"tns"`
			Alias []interface{} `json:"alias"`
		} `json:"ar"`
		Alia []interface{} `json:"alia"`
		Pop  int           `json:"pop"`
		St   int           `json:"st"`
		Rt   string        `json:"rt"`
		Fee  int           `json:"fee"`
		V    int           `json:"v"`
		Crbt interface{}   `json:"crbt"`
		Cf   string        `json:"cf"`
		Al   struct {
			ID     int           `json:"id"`
			Name   string        `json:"name"`
			PicURL string        `json:"picUrl"`
			Tns    []interface{} `json:"tns"`
			PicStr string        `json:"pic_str"`
			Pic    int64         `json:"pic"`
		} `json:"al"`
		Dt int `json:"dt"`
		H  struct {
			Br   int     `json:"br"`
			Fid  int     `json:"fid"`
			Size int     `json:"size"`
			Vd   float64 `json:"vd"`
			Sr   int     `json:"sr"`
		} `json:"h"`
		M struct {
			Br   int     `json:"br"`
			Fid  int     `json:"fid"`
			Size int     `json:"size"`
			Vd   float64 `json:"vd"`
			Sr   int     `json:"sr"`
		} `json:"m"`
		L struct {
			Br   int     `json:"br"`
			Fid  int     `json:"fid"`
			Size int     `json:"size"`
			Vd   float64 `json:"vd"`
			Sr   int     `json:"sr"`
		} `json:"l"`
		Sq struct {
			Br   int     `json:"br"`
			Fid  int     `json:"fid"`
			Size int     `json:"size"`
			Vd   float64 `json:"vd"`
			Sr   int     `json:"sr"`
		} `json:"sq"`
		Hr                   interface{}   `json:"hr"`
		A                    interface{}   `json:"a"`
		Cd                   string        `json:"cd"`
		No                   int           `json:"no"`
		RtURL                interface{}   `json:"rtUrl"`
		Ftype                int           `json:"ftype"`
		RtUrls               []interface{} `json:"rtUrls"`
		DjID                 int           `json:"djId"`
		Copyright            int           `json:"copyright"`
		SID                  int           `json:"s_id"`
		Mark                 int           `json:"mark"`
		OriginCoverType      int           `json:"originCoverType"`
		OriginSongSimpleData interface{}   `json:"originSongSimpleData"`
		TagPicList           interface{}   `json:"tagPicList"`
		ResourceState        bool          `json:"resourceState"`
		Version              int           `json:"version"`
		SongJumpInfo         interface{}   `json:"songJumpInfo"`
		EntertainmentTags    interface{}   `json:"entertainmentTags"`
		AwardTags            interface{}   `json:"awardTags"`
		Single               int           `json:"single"`
		NoCopyrightRcmd      interface{}   `json:"noCopyrightRcmd"`
		Rtype                int           `json:"rtype"`
		Rurl                 interface{}   `json:"rurl"`
		Mst                  int           `json:"mst"`
		Cp                   int           `json:"cp"`
		Mv                   int           `json:"mv"`
		PublishTime          int64         `json:"publishTime"`
		Tns                  []string      `json:"tns"`
	} `json:"songs"`
	Privileges []struct {
		ID                 int         `json:"id"`
		Fee                int         `json:"fee"`
		Payed              int         `json:"payed"`
		St                 int         `json:"st"`
		Pl                 int         `json:"pl"`
		Dl                 int         `json:"dl"`
		Sp                 int         `json:"sp"`
		Cp                 int         `json:"cp"`
		Subp               int         `json:"subp"`
		Cs                 bool        `json:"cs"`
		Maxbr              int         `json:"maxbr"`
		Fl                 int         `json:"fl"`
		Toast              bool        `json:"toast"`
		Flag               int         `json:"flag"`
		PreSell            bool        `json:"preSell"`
		PlayMaxbr          int         `json:"playMaxbr"`
		DownloadMaxbr      int         `json:"downloadMaxbr"`
		MaxBrLevel         string      `json:"maxBrLevel"`
		PlayMaxBrLevel     string      `json:"playMaxBrLevel"`
		DownloadMaxBrLevel string      `json:"downloadMaxBrLevel"`
		PlLevel            string      `json:"plLevel"`
		DlLevel            string      `json:"dlLevel"`
		FlLevel            string      `json:"flLevel"`
		Rscl               interface{} `json:"rscl"`
		FreeTrialPrivilege struct {
			ResConsumable  bool        `json:"resConsumable"`
			UserConsumable bool        `json:"userConsumable"`
			ListenType     interface{} `json:"listenType"`
		} `json:"freeTrialPrivilege"`
		ChargeInfoList []struct {
			Rate          int         `json:"rate"`
			ChargeURL     interface{} `json:"chargeUrl"`
			ChargeMessage interface{} `json:"chargeMessage"`
			ChargeType    int         `json:"chargeType"`
		} `json:"chargeInfoList"`
	} `json:"privileges"`
	Code int `json:"code"`
}

func QueryMusic(id int) (*player.Music, error) {
	logger := config.Logger
	if player.MusicPlayer.GetMusicByID(strconv.Itoa(id)) != nil {
		logger.Info().Msg(fmt.Sprintf("Music %s added from cache", player.MusicPlayer.GetMusicByID(strconv.Itoa(id)).Name))
		return player.MusicPlayer.GetMusicByID(strconv.Itoa(id)), nil
	}
	url, err := QueryMusicURL(id)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to query music url")
		return nil, err
	}
	path, err := DownloadMusic(id, url)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to download music")
		return nil, err
	}
	music, err := QueryMusicInfo(id)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to query music info")
		return nil, err
	}
	music.File = path
	logger.Info().Msgf("Music %s downloaded", music.Name)
	return music, nil
}

func QueryMusicURL(id int) (string, error) {
	return "", nil
	//logger := config.Logger
	//musicReq := &MusicReq{
	//	ID:      id,
	//	Quality: "standard",
	//}
	//client := &http.Client{}
	////apiUrl := fmt.Sprintf("%s/song/url/v1", config.Config.NetEaseAPI)
	//req, err := http.NewRequest("GET", apiUrl, nil)
	//if err != nil {
	//	logger.Error().Err(err).Msg("Failed to create request")
	//	return "", err
	//}
	//req.Header.Set("Cookie", config.Config.NetEaseCookie)
	//params := req.URL.Query()
	//params.Add("id", strconv.Itoa(musicReq.ID))
	//params.Add("level", musicReq.Quality)
	//req.URL.RawQuery = params.Encode()
	//
	//resp, err := client.Do(req)
	//if err != nil {
	//	logger.Error().Err(err).Msg("Failed to send request")
	//	return "", err
	//}
	//defer resp.Body.Close()
	//
	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	logger.Error().Err(err).Msg("Failed to read response")
	//	return "", err
	//}
	//if resp.StatusCode != http.StatusOK {
	//	logger.Error().Msg("Failed to get music url")
	//	return "", err
	//}
	//var musicResp MusicResp
	//logger.Debug().Msg(string(body))
	//if err := json.Unmarshal(body, &musicResp); err != nil {
	//	logger.Error().Err(err).Msg("Failed to unmarshal response")
	//	return "", err
	//}
	//return musicResp.Data[0].URL, nil
}

func DownloadMusic(id int, url string) (string, error) {
	logger := config.Logger
	resp, err := http.Get(url)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to download music")
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		logger.Error().Msg("Failed to download music, status code: " + strconv.Itoa(resp.StatusCode))
		return "", err
	}
	defer resp.Body.Close()
	f, err := os.Create("./assets/voice/N" + strconv.Itoa(id) + ".mp3")
	defer f.Close()
	if err != nil {
		logger.Error().Err(err).Msg("Failed to create file")
		return "", err
	}
	_, err = io.Copy(f, resp.Body)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to write file")
		return "", err
	}
	return "./assets/voice/N" + strconv.Itoa(id) + ".mp3", nil
}

func QueryMusicInfo(id int) (*player.Music, error) {
	return nil, nil
	//logger := config.Logger
	//client := &http.Client{}
	//req, err := http.NewRequest("GET", config.Config.NetEaseAPI+"/song/detail", nil)
	//if err != nil {
	//	logger.Error().Err(err).Msg("Failed to create request")
	//	return nil, err
	//}
	//req.Header.Set("Cookie", config.Config.NetEaseCookie)
	//params := req.URL.Query()
	//params.Add("ids", strconv.Itoa(id))
	//req.URL.RawQuery = params.Encode()
	//
	//resp, err := client.Do(req)
	//if err != nil {
	//	logger.Error().Err(err).Msg("Failed to send request")
	//	return nil, err
	//}
	//defer resp.Body.Close()
	//
	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	logger.Error().Err(err).Msg("Failed to read response")
	//	return nil, err
	//}
	//if resp.StatusCode != http.StatusOK {
	//	logger.Error().Msg("Failed to get music info")
	//	return nil, err
	//}
	//var musicInfoResp MusicInfoResp
	//if err := json.Unmarshal(body, &musicInfoResp); err != nil {
	//	logger.Error().Err(err).Msg("Failed to unmarshal response")
	//	return nil, err
	//}
	//
	//ar := make([]string, 0)
	//for _, v := range musicInfoResp.Songs[0].Ar {
	//	ar = append(ar, v.Name)
	//}
	//return &player.Music{
	//	ID:       strconv.Itoa(id),
	//	Name:     musicInfoResp.Songs[0].Name,
	//	Artists:  ar,
	//	Album:    musicInfoResp.Songs[0].Al.PicURL + "?param=130y130",
	//	LastTime: musicInfoResp.Songs[0].Dt,
	//}, nil
}
