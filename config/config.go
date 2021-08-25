package config

import (
	_ "embed"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	ScanTarget
	Server
}

type ScanTarget struct {
	TargetUrl string `yaml:"target_url"`
}

type Server struct {
	Port string `yaml:"port"`
	Addr string `yaml:"addr"`
}

//go:embed config.yaml
var configContent []byte

var cf *Config

func init() {
	// 读出config.yaml 绑定基本配置参数
	err := yaml.Unmarshal(configContent, &cf)
	if err != nil {
		log.Fatal(err)
	}
}

func Default() *Config {
	return cf
}
