package main

type Config struct {
	Server server `yaml:"server"`
}

type ScanTarget struct {
	TargetUrl string `yaml:"target_url"`
}

type server struct {
	Port string `yaml:"port"`
	Addr string `yaml:"addr"`
}
