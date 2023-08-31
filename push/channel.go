package push

import (
	"encoding/json"
	"time"
)

type ChannelReq struct {
	Appkey    string `json:"appkey"`
	Timestamp int64  `json:"timestamp"`
	TaskId    string `json:"task_id"`
}

type ChannelResp struct {
	Ret  string      `json:"ret"`
	Data ChannelData `json:"data"`
}

type ChannelErr struct {
	ErrorCode string `json:"error_code"`
	ErrorInfo string `json:"error_info"`
	Num       string `json:"num"`
}

type ChannelData struct {
	Stats []struct {
		Channel            string       `json:"channel"`
		ChannelArriveCount int          `json:"channel_arrive_count"`
		ChannelClick       int          `json:"channel_click"`
		ChannelSentCount   int          `json:"channel_sent_count"`
		Errors             []ChannelErr `json:"errors"`
	} `json:"stats"`
}

func (c *Client) Channel(taskId string) (ret ChannelData, err error) {
	var result []byte
	data := ChannelReq{c.AppKey, time.Now().Unix(), taskId}
	if result, err = c.Request(Host+ChannelPath, data); err != nil {
		return
	}
	var resp ChannelResp
	if err = json.Unmarshal(result, &resp); err != nil {
		return
	}
	ret = resp.Data
	return
}
