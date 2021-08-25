package service

import (
	"fmt"
	"net/url"
)

func dealJs(targeturl string, jsPath []string) {
	// 处理url多余部分
	res, _ := url.Parse(targeturl)

	fmt.Println(res)
}
