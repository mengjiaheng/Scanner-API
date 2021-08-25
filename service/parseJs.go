/*
 * @Author: manguanghui
 * @Date: 2021-08-25 15:43:17
 * @Desc: file content
 */
package service

import (
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/mengjiaheng/scanapi/config"
	"github.com/mengjiaheng/scanapi/utils/file"
	"github.com/mengjiaheng/scanapi/utils/request"
)

func RequestUrl(url string) error {

	req := request.NewRequest()
	response, err := req.Request("GET", url, nil)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return err
	}
	path := make(map[string]int)
	doc.Find("script,link").Each(func(i int, s *goquery.Selection) {
		if s.Text() != "" {
			if err := file.IsNotExistMkDir("./jsFile"); err != nil {
				fmt.Println(err.Error())
				return
			}
			// 将标签里面的js代码读取出来放到指定文件
			f, err := file.Open("./jsFile/ss", os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModeAppend)
			if err != nil {
				fmt.Println("文件创建错误", err.Error())
				return
			}
			_, err = f.Write([]byte(s.Text()))
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			defer f.Close()
			// fmt.Printf("%d: %s\n", i, s.Text())
		} else {
			value1, _ := s.Attr("href")
			value2, _ := s.Attr("src")
			if value1 != "" {
				if strings.HasSuffix(value1, ".js") {
					// 拼接URL
					b, err := BlackListDomain(DealJs(url, value1))
					if err != nil {
						return
					}
					if !b {
						path[DealJs(url, value1)]++
					}
				}
			}
			if value2 != "" {
				if strings.HasSuffix(value2, ".js") {
					// 拼接URL
					b, err := BlackListDomain(DealJs(url, value2))
					if err != nil {
						return
					}
					if !b {
						path[DealJs(url, value2)]++
					}
				}
			}
		}
	})
	downloadJs(path)
	return nil
}

// 获取js文件名
func GetFileName(url string) string {
	array := strings.Split(url, "/")
	filename := array[len(array)-1]
	return strings.Split(filename, "?")[0]
}

// 过滤黑名单js和域名
func BlackListDomain(path string) (bool, error) {
	res, err := url.Parse(path)
	if err != nil {
		return false, err
	}
	jsRealPathDomain := strings.ToLower(res.Host)
	jsRealPathFilename := strings.ToLower(GetFileName(path))

	// 获取黑名单域名
	conf := config.Default()
	var flag int
	for _, v := range strings.Split(conf.Blacklist.Domain, ",") {
		if strings.Count(jsRealPathDomain, v) > 0 { // 是否存在黑名单域名
			flag = 1
			break
		} else {
			flag = 0
		}
	}
	if flag > 0 {
		return true, nil
	}
	for _, v := range strings.Split(conf.Blacklist.Filename, ",") {
		if strings.Count(jsRealPathFilename, v) > 0 {
			flag = 1
			break
		} else {
			flag = 0
		}
	}
	if flag > 0 {
		return true, nil
	}
	return false, nil
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
