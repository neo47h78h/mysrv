package ai

import (
	"context"
	"fmt"

	"github.com/yankeguo/zhipu"
)

func GetAIResponse(content string) (retString string, err error) {
	client, err := zhipu.NewClient()
	// or you can specify the API key
	client, err = zhipu.NewClient(zhipu.WithAPIKey("ae47f3ed9bfaf53f2c085ea334061800.IA1IyKt4Lm1xoLgF"))
	service := client.ChatCompletion("glm-4-flash").
		AddMessage(zhipu.ChatCompletionMessage{
			Role:    "user",
			Content: content,
		})

	res, err := service.Do(context.Background())
	if err != nil {
		err = fmt.Errorf("err:%s", zhipu.GetAPIErrorCode(err))
		return
	}

	retString = res.Choices[0].Message.Content
	return
}
