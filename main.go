/*
 * @Author: manguanghui
 * @Date: 2021-08-25 10:43:51
 * @Desc: file content
 */
package main

import (
	"github.com/mengjiaheng/scanapi/config"
	"github.com/mengjiaheng/scanapi/service"
)

func main() {
	cf := config.Default()
	service.RequestUrl(cf.TargetUrl)
}
