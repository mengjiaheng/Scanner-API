package service

import (
	"fmt"
	"net/url"
)

func RequestUrl(targeturl string) error {

	// request, err := http.NewRequest("GET", targeturl, nil)
	// if err != nil {
	// 	return err
	// }

	// request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:78.0) Gecko/20100101 Firefox/78.0")

	// client := &http.Client{}
	// response, err := client.Do(request)
	// if err != nil {
	// 	return err
	// }
	// bytes, err := ioutil.ReadAll(response.Body)
	// if err != nil {
	// 	return err
	// }
	// defer response.Body.Close()
	// fmt.Println(string(bytes))

	jsPath := make([]string, 0)
	jsPath = append(jsPath, "./static/js/8.50e09590.chunk.js")
	jsPath = append(jsPath, "./static/js/main.66ccc09c.chunk.js")

	dealJs(targeturl, jsPath)
	return nil
}

func dealJs(targeturl string, jsPath []string) {
	// 处理url多余部分
	res, _ := url.Parse(targeturl)

	fmt.Println(res)
}
