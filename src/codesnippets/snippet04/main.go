package main

import (
	"encoding/json"
	"net/http"
	"bytes"
	"time"
	"errors"
	"log"
)

/*
群发消息(用一个库)：(通知模式)
Sending a Slack Message (without a library)
*/
type SlackRequesstBody struct {
	Text string `json:"text"`
}

func main() {
	//通知地址
	webhookurl := "https://hooks.slack.com/services/X1234"
	err := SendSlackNotification(webhookurl, "test message ")
	if err != nil {
		log.Fatal(err)
	}
}

//发送通知
func SendSlackNotification(webhookurl string, data string) error {
	//组装成json数据
	slackBody, _ := json.Marshal(SlackRequesstBody{Text: data})
	//发送请求
	//生成请求对象
	req, err := http.NewRequest(http.MethodPost, webhookurl, bytes.NewBuffer(slackBody))
	if err != nil {
		return err
	}
	//设置请求头
	req.Header.Set("Content-Type", "application/json")

	//生成链接客户端
	client := &http.Client{Timeout: 10 * time.Second}
	//建立链接
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	//解析响应文本
	//定义一个缓冲流
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	if buf.String() != "ok" {
		return errors.New("non-ok response return from slack ")
	}
	return nil
}
