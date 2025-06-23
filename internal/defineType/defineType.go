package defineType

import "time"

// 微信公众号配置
type WxConfig struct {
	Appid     string `json:"appid"`
	Secret    string `json:"secret"`
	GrantType string `json:"grant_type"`
}

type AutoPhone struct {
	RequestUrl string `json:"requestUrl"`
	Code       string `json:"code"`
}

type AccessToken struct {
	RequestUrl string `json:"requestUrl"`
	GrantType  string `json:"grant_type"`
	Appid      string `json:"appid"`
	Secret     string `json:"secret"`
}

// 小程序
type MinConfig struct {
	RequestUrl string `json:"requestUrl"`
	Appid      string `json:"appid"`
	Secret     string `json:"secret"`
	GrantType  string `json:"grant_type"`
}

// 运行日志
type OutputsForLogger struct {
	Time       time.Time `json:"time"`
	Host       string    `json:"host"`
	RequestURI string    `json:"requestURI"`
	Params     string    `json:"params"`
	RunTime    float64   `json:"runTime"`
	File       string    `json:"file"`
	FileRule   string    `json:"fileRule"`
	RotateSize string    `json:"rotateSize"`
	Stdout     bool      `json:"stdout"`
	Path       string    `json:"path"`
	Throughput float64   `json:"throughput"`
	MemUsed    uint64    `json:"memUsed"`
}
