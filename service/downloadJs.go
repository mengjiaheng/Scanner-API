package service

import (
	"fmt"
	"io/ioutil"

	"github.com/mengjiaheng/scanapi/utils/request"
)

func downloadJs(jsPath map[string]int) {
	req := request.NewRequest()
	req.AddHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:78.0) Gecko/20100101 Firefox/78.0")
	req.AddHeader("Content-Type", "application/x-www-form-urlencoded")
	req.AddHeader("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")

	for k, _ := range jsPath {
		rep, _ := req.Request("GET", k, nil)
		bytes, _ := ioutil.ReadAll(rep.Body)
		fmt.Println("测试：", string(bytes))
	}

}
