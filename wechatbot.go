package gw2w

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const wechatURL = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=%s"

// WeChatBOTNewsBody .
type WeChatBOTNewsBody struct {
	ChatID        *string       `json:"chatid"`
	MSGType       string        `json:"msgtype"`
	VisibleToUser *string       `json:"visible_to_user"`
	News          WeChatBOTNews `json:"news"`
}

// WeChatBOTNews .
type WeChatBOTNews struct {
	Articles []WeChatBOTNewsArticle `json:"articles"`
}

// WeChatBOTNewsArticle .
type WeChatBOTNewsArticle struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	PicURL      string `json:"picurl"`
}

// WeChatBOTRsp .
type WeChatBOTRsp struct {
	ErrCode int    `json:"errcode"`
	ErrMSG  string `json:"errmsg"`
}

// Send .
func (b WeChatBOTNewsBody) Send(ctx context.Context, key string) error {
	buf, err := json.Marshal(b)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", fmt.Sprintf(wechatURL, key), bytes.NewBuffer(buf))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	rsp, err := client.Do(req)
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return err
	}
	weChatBOTRsp := &WeChatBOTRsp{}
	if err = json.Unmarshal(body, weChatBOTRsp); err != nil {
		return err
	}
	if weChatBOTRsp.ErrCode != 0 {
		return fmt.Errorf("error(%d) %s", weChatBOTRsp.ErrCode, weChatBOTRsp.ErrMSG)
	}
	defer rsp.Body.Close()

	return nil
}
