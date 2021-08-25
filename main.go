package main

import (
	"github.com/mengjiaheng/scanapi/config"
	"github.com/mengjiaheng/scanapi/service"
)

func main() {
	cf := config.Default()
	service.RequestUrl(cf.TargetUrl)
}
