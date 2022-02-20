/*
Name: plume
File: utils_http.go
Author: Landers
*/

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// 基于http的rest请求

func Get(url string) ([]byte, error) {
	res, e := http.Get(url)
	if e != nil {
		return nil, e
	}
	data, e := io.ReadAll(res.Body)
	defer res.Body.Close()
	if e != nil {
		return nil, e
	}
	return data, nil
}

func Post(url string, body interface{}) ([]byte, error) {
	client := &http.Client{Timeout: 30 * time.Second}
	jsonStr, e := json.Marshal(body)
	if e != nil {
		return nil, e
	}

	res, e := client.Post(url, "application/json", bytes.NewBuffer(jsonStr))
	if e != nil {
		return nil, e
	}
	data, e := io.ReadAll(res.Body)
	defer res.Body.Close()
	if e != nil {
		return nil, e
	}
	return data, nil
}

func PostStatus(url string, body io.ReadCloser) string {
	client := &http.Client{Timeout: 30 * time.Second}
	res, e := client.Post(url, "application/json", body)
	if e != nil {
		fmt.Println(e)
		return ""
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("create container: %v\n", err)
	} else {
		fmt.Printf("create container: %s\n", data)
	}

	return fmt.Sprintf("%d", res.StatusCode)
}

func DeleteStatus(url string) string {
	client := &http.Client{Timeout: 30 * time.Second}
	req, e := http.NewRequest("DELETE", url, nil)
	req.Header.Set("Content-Type", "application/json")
	if e != nil {
		return ""
	}
	res, e := client.Do(req)
	if e != nil {
		return ""
	}

	return fmt.Sprintf("%d", res.StatusCode)
}
