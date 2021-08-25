/*
 * @Author: manguanghui
 * @Date: 2021-08-25 10:58:11
 * @Desc: file content
 */
package ceshi

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func Ceshi(url string) {

	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer response.Body.Close()
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	path := make(map[string]int, 0)
	// b := []byte(url)
	doc.Find("script,link").Each(func(i int, s *goquery.Selection) {
		if s.Text() != "" {
			fmt.Printf("%d: %s\n", i, s.Text())
		} else {
			value1, _ := s.Attr("href")
			value2, _ := s.Attr("src")
			if value1 != "" {
				if strings.HasSuffix(value1, ".js") {
					fmt.Printf("%d: %s\n", i, value1)
					// 拼接URL
					path[DealJs(url, value1)]++
				}
			}
			if value2 != "" {
				if strings.HasSuffix(value2, ".js") {
					fmt.Printf("%d: %s\n", i, value2)
					// 拼接URL
					path[DealJs(url, value2)]++
				}
			}
		}
	})
	fmt.Println(path)
}

// 生成JS绝对路径
func DealJs(str, js_path string) string {
	res, _ := url.Parse(str)
	var baseUrl string
	tmpPath := make([]string, 0)
	if res.Path == "" {
		baseUrl = res.Scheme + "://" + res.Host + "/"
	} else {
		baseUrl = res.Scheme + "://" + res.Host + res.Path
		if !strings.HasSuffix(res.Path, "/") {
			baseUrl = baseUrl + "/"
		}
	}
	if !strings.HasSuffix(str, "/") {
		tmpPath = strings.Split(res.Path, "/")
		tmpPath = tmpPath[:]

		tmpPath = tmpPath[:len(tmpPath)-1]
		baseUrl = res.Scheme + "://" + res.Host + strings.Join(tmpPath, "/") + "/"
	}
	if strings.HasPrefix(js_path, "http") {
		return js_path
	} else if strings.HasPrefix(js_path, "../") {
		dirCount := strings.Count(js_path, "../")
		tmpCount := 1
		js_path = strings.Replace(js_path, "../", "", -1)
		new_tmpPath := tmpPath[:]
		for tmpCount <= dirCount {
			new_tmpPath = new_tmpPath[:len(new_tmpPath)-1]
			tmpCount = tmpCount + 1
		}
		baseUrl = res.Scheme + "://" + res.Host + strings.Join(new_tmpPath, "/") + "/"
		return baseUrl + js_path
	} else if strings.HasPrefix(js_path, "//") {
		return res.Scheme + ":" + js_path
	} else if strings.HasPrefix(js_path, "./") {
		return baseUrl + strings.Replace(js_path, "./", "", -1)
	} else if strings.HasPrefix(js_path, "/") {
		return res.Scheme + "://" + res.Host + js_path
	} else {
		return baseUrl + js_path
	}
}
