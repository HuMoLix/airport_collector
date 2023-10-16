package common

import (
	"time"

	"github.com/oschwald/geoip2-golang"
)

const (
	Version = "0.2_release"
	Author  = "HuMoLix"
)
const Banner = `
█████╗ ██╗██████╗ ██████╗  ██████╗ ██████╗ ████████╗
██╔══██╗██║██╔══██╗██╔══██╗██╔═══██╗██╔══██╗╚══██╔══╝
███████║██║██████╔╝██████╔╝██║   ██║██████╔╝   ██║   
██╔══██║██║██╔══██╗██╔═══╝ ██║   ██║██╔══██╗   ██║   
██║  ██║██║██║  ██║██║     ╚██████╔╝██║  ██║   ██║   By: ` + Author + `
╚═╝  ╚═╝╚═╝╚═╝  ╚═╝╚═╝      ╚═════╝ ╚═╝  ╚═╝   ╚═╝   Version: ` + Version + `
                                                     
██████╗ ██████╗ ██╗     ██╗     ███████╗ ██████╗████████╗███████╗██████╗ 
██╔════╝██╔═══██╗██║     ██║     ██╔════╝██╔════╝╚══██╔══╝██╔════╝██╔══██╗
██║     ██║   ██║██║     ██║     █████╗  ██║        ██║   █████╗  ██████╔╝
██║     ██║   ██║██║     ██║     ██╔══╝  ██║        ██║   ██╔══╝  ██╔══██╗
╚██████╗╚██████╔╝███████╗███████╗███████╗╚██████╗   ██║   ███████╗██║  ██║
 ╚═════╝ ╚═════╝ ╚══════╝╚══════╝╚══════╝ ╚═════╝   ╚═╝   ╚══════╝╚═╝  ╚═╝ 

 工具仅用于信息收集，请勿用于非法用途
 开发人员不承担任何责任，也不对任何滥用或损坏负责.                                                                         

`
const (
	ConfigYaml  = "config.yaml"
	License     = ".license"
	GeoIP       = "geoip"
	GeoIPPrefix = "GeoLite2-"
	Results     = "results"
)

const (
	routineCountTotal = 40
	Timeout           = 5 * time.Second
)

type ApCOptions struct {
	QueryPageCount int
	Runmode        string
	GeoIP          bool
}

type Conf struct {
	// Version  string   `yaml:"version"`  // 程序版本，用于检查更新
	Settings Settings `yaml:"settings"` // 全局变量
	Fofa     FofaConf `yaml:"fofa"`     // Fofa
}

type Settings struct {
	IsBlackList bool   `yaml:"isBlackList"` // 是否记录无法登录的IP地址
	License     string // License
}

type FofaConf struct {
	Email  string `yaml:"email"`  // 邮箱
	Apikey string `yaml:"apikey"` // ApiKey
	Query  Query  `yaml:"query"`  // 查询参数
}

type Query struct {
	Total_Size      int `yaml:"total_Size"`      // 总数量
	Last_Query_Page int `yaml:"last_Query_page"` // 程序结束时查询的页数
}

type Geoip struct {
	City    *geoip2.Reader
	Country *geoip2.Reader
	ASN     *geoip2.Reader
}

type XML struct {
	Port                int                      `yaml:"port"`                // HTTP 代理端口
	Socks_Port          int                      `yaml:"socks-port"`          // SOCKS5 代理端口
	Redir_Port          int                      `yaml:"redir-port"`          // Linux 和 macOS 的 redir 代理端口
	Allow_LAN           bool                     `yaml:"allow-lan"`           // 允许局域网的连接
	Mode                string                   `yaml:"mode"`                // 规则模式：Rule（规则） / Global（全局代理）/ Direct（全局直连）
	Log_Level           string                   `yaml:"log-level"`           // 设置日志输出级别 (默认级别：silent，即不输出任何内容，以避免因日志内容过大而导致程序内存溢出）。
	External_Controller string                   `yaml:"external-controller"` // Clash 的 RESTful API
	Secret              string                   `yaml:"secret"`              // RESTful API 的口令
	CFW_Latency_Timeout int                      `yaml:"cfw-latency-timeout"`
	CFW_Bypass          []string                 `yaml:"cfw-bypass"`
	Proxies             []map[string]interface{} `yaml:"proxies"`
	Proxy_Groups        []Proxy_Groups           `yaml:"proxy-groups"`
	Rule                []string                 `yaml:"Rule"`
}

type Proxy_Groups struct {
	Name    string   `yaml:"name"`
	Type    string   `yaml:"type"`
	Proxies []string `yaml:"proxies"`
}
