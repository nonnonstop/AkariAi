package talk

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type RequestData struct {
	Key      string `json:"key"`
	Message  string `json:"message"`
	BotName  string `json:"bot_name"`
	UserId   string `json:"user_id"`
	UserName string `json:"user_name"`
}

type ResponseData struct {
	Status string `json:"status"`
	Result string `json:"result"`
}

func (t *Client) Talk(message, userId, userName string) string {
	reqData := &RequestData{
		Key:      t.config.Token,
		Message:  message,
		BotName:  "アカリ",
		UserId:   userId,
		UserName: userName,
	}
	reqBody, err := json.Marshal(reqData)
	if err != nil {
		t.logger.Errorln("Failed to marshal: ", err)
		return ""
	}
	resBody, err := requestAPI(reqBody)
	if err != nil {
		t.logger.Errorln("Failed to request API: ", err)
		return ""
	}
	var resData ResponseData
	err = json.Unmarshal(resBody, &resData)
	if err != nil {
		t.logger.Errorln("Failed to unmarshal: ", err)
		return ""
	}
	return resData.Result
}

func requestAPI(reqBody []byte) ([]byte, error) {
	res, err := http.Post(
		"https://chatbot-api.userlocal.jp/api/chat",
		"application/json",
		bytes.NewBuffer(reqBody),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return io.ReadAll(res.Body)
}
