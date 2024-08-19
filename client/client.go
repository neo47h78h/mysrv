package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// 服务器地址
	url := "http://localhost:8080"

	// 发送 HTTP GET 请求
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error making GET request: %v\n", err)
		return
	}
	defer resp.Body.Close() // 确保请求完成后关闭响应体

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return
	}

	// 打印响应状态和内容
	fmt.Printf("Response Status: %s\n", resp.Status)
	fmt.Printf("Response Body: %s\n", body)
}
