package NetEase

import (
	"MusicBot/serve/player"
)

type SearchResp struct {
	Result struct {
		Songs []struct {
			ID      int    `json:"id"`
			Name    string `json:"name"`
			Artists []struct {
				ID        int           `json:"id"`
				Name      string        `json:"name"`
				PicURL    interface{}   `json:"picUrl"`
				Alias     []interface{} `json:"alias"`
				AlbumSize int           `json:"albumSize"`
				PicID     int           `json:"picId"`
				FansGroup interface{}   `json:"fansGroup"`
				Img1V1URL string        `json:"img1v1Url"`
				Img1V1    int           `json:"img1v1"`
				Trans     interface{}   `json:"trans"`
			} `json:"artists"`
			Album struct {
				ID     int    `json:"id"`
				Name   string `json:"name"`
				Artist struct {
					ID        int           `json:"id"`
					Name      string        `json:"name"`
					PicURL    interface{}   `json:"picUrl"`
					Alias     []interface{} `json:"alias"`
					AlbumSize int           `json:"albumSize"`
					PicID     int           `json:"picId"`
					FansGroup interface{}   `json:"fansGroup"`
					Img1V1URL string        `json:"img1v1Url"`
					Img1V1    int           `json:"img1v1"`
					Trans     interface{}   `json:"trans"`
				} `json:"artist"`
				PublishTime int64 `json:"publishTime"`
				Size        int   `json:"size"`
				CopyrightID int   `json:"copyrightId"`
				Status      int   `json:"status"`
				PicID       int64 `json:"picId"`
				Mark        int   `json:"mark"`
			} `json:"album,omitempty"`
			Duration    int           `json:"duration"`
			CopyrightID int           `json:"copyrightId"`
			Status      int           `json:"status"`
			Alias       []interface{} `json:"alias"`
			Rtype       int           `json:"rtype"`
			Ftype       int           `json:"ftype"`
			TransNames  []string      `json:"transNames,omitempty"`
			Mvid        int           `json:"mvid"`
			Fee         int           `json:"fee"`
			RURL        interface{}   `json:"rUrl"`
			Mark        int           `json:"mark"`
		} `json:"songs"`
		HasMore   bool `json:"hasMore"`
		SongCount int  `json:"songCount"`
	} `json:"result"`
	Code int `json:"code"`
}

type SearchReq struct {
	Keywords string
	Type     int
}

func SearchMusic(keywords string) (*[]player.Music, error) {
	return nil, nil
	//logger := config.Logger
	//searchReq := &SearchReq{
	//	Keywords: keywords,
	//	Type:     1,
	//}
	//client := &http.Client{}
	//req, err := http.NewRequest("GET", config.Config.NetEaseAPI+"/search", nil)
	//if err != nil {
	//	logger.Error().Err(err).Msg("Failed to create request")
	//	return nil, err
	//}
	//req.Header.Set("Cookie", config.Config.NetEaseCookie)
	//params := req.URL.Query()
	//params.Add("keywords", searchReq.Keywords)
	//params.Add("type", strconv.Itoa(searchReq.Type))
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
	//	logger.Error().Msg("Failed to get music url")
	//	return nil, err
	//}
	//var searchResp SearchResp
	//logger.Debug().Msg(string(body))
	//if err := json.Unmarshal(body, &searchResp); err != nil {
	//	logger.Error().Err(err).Msg("Failed to unmarshal response")
	//	return nil, err
	//}
	//var musicsList []player.Music
	//for i, song := range searchResp.Result.Songs {
	//	if i > config.Config.SearchLimit {
	//		break
	//	}
	//	musicResp := &player.Music{
	//		ID: strconv.Itoa(song.ID),
	//	}
	//	musicsList = append(musicsList, *musicResp)
	//}
	//return &musicsList, nil
}
