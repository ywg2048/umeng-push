package push

import (
	"encoding/json"
	"time"
)

type CancelReq struct {
	Appkey    string `json:"appkey"`
	Timestamp int64  `json:"timestamp"`
	TaskId    string `json:"task_id"`
}

type CancelResp struct {
	Ret  string `json:"ret"`
	Data struct {
		TaskId string `json:"task_id"`
	} `json:"data"`
}

func (c *Client) Cancel(taskId string) (ret CancelResp, err error) {
	var result []byte
	data := CancelReq{c.AppKey, time.Now().Unix(), taskId}
	if result, err = c.Request(Host+CancelPath, data); err != nil {
		return
	}

	if err = json.Unmarshal(result, &ret); err != nil {
		return
	}
	return
}
