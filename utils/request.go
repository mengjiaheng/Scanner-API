/*
 * @Author: manguanghui
 * @Date: 2021-08-25 16:21:59
 * @Desc: file content
 */
package utils

import (
	"bytes"
	"net/http"
)

type Request struct {
	header map[string]string
}

func NewRequest() *Request {
	return &Request{
		header: make(map[string]string),
	}
}

// 添加请求头
func (r *Request) AddHeader(key, value string) {
	r.header[key] = value
}

func (r *Request) Request(method, url string, data []byte) (*http.Response, error) {

	body := bytes.NewBuffer(data)
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	for k, v := range r.header {
		request.Header.Add(k, v)
	}

	client := &http.Client{}
	return client.Do(request)
}
