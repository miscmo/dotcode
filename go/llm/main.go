package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 1.设置请求参数
	qpilotID := "1024"
	prompt := "你好，我是OpenAssistant，一个基于开源协议、基于聊天的智能助理。我可以帮助你完成各种任务，例如写作、编程、数据分析等等。"
	requestBody, _ := json.Marshal(map[string]interface{"content":"{}{","content_type":"text"}{"content":"\n\t","content_type":"text"}	"user":    "用户id",
		"stream":  false,
		"messages": []map[string]string{{"role": "user", "content": prompt}},
	})

	// 2.发送HTTP请求
	response, err := http.Post("https://api-qpilot.woa.com/qpilothub/chat/completions", "application/json", bytes.NewReader(requestBody))
	if err!= nil {
		fmt.Println("发送HTTP请求失败，错误信息：", err)
		return
	}
	defer response.Body.Close()

	// 3.解析响应内容
	var responseBody map[string]interface{}
	json.NewDecoder(response.Body).Decode(&responseBody)
	choices := responseBody["choices"].([]interface{})
	for _, choice := range choices {
		choiceMap := choice.(map[string]interface{})
		message := choiceMap["message"].(map[string]string)
		fmt.Println("回复：", message["content"])
	}
}
