package fofa

import (
	"fmt"
	"strconv"

	"github.com/HuMoLix/airport_collector/common"
	"github.com/go-resty/resty/v2"
)

type FofaInfo struct {
	Error            bool   `json:"error"`            // 是否出现错误
	Email            string `json:"email"`            // 邮箱地址
	Username         string `json:"username"`         // 用户名
	Category         string `json:"category"`         // 用户种类
	Remain_Api_Query int    `json:"remain_api_query"` // API月度剩余查询次数
	Remain_Api_Data  int    `json:"remain_api_data"`  // API阅读剩余返回数量
	IsVip            bool   `json:"is_vip"`           // 是否是会员
	Message          string `json:"message"`          // 返回信息
}

type FofaHostInfo struct {
	Error            bool       `json:"error"`            // 是否出现错误
	ErrMsg           string     `json:"errmsg"`           // 报错信息
	Page             int        `json:"page"`             // 当前页码
	Size             int        `json:"size"`             // 查询总数量
	Required_fPoints int        `json:"required_fpoints"` // 扣除F点
	Results          [][]string `json:"results"`          // 查询结果
}

func ConnectFofa(Conf *common.Conf) bool {
	client := resty.New()
	fofainfo := &FofaInfo{}
	client.R().SetResult(fofainfo).Get("https://fofa.info/api/v1/info/my?email=" + Conf.Fofa.Email + "&key=" + Conf.Fofa.Apikey)
	if fofainfo.Error {
		fmt.Println("[-]错误，请检查用户名或apiKey是否错误。")
		return false
	}
	fmt.Printf("[+]用户%v登录成功，本月还剩余%d次查询次数。\n", fofainfo.Username, fofainfo.Remain_Api_Query)
	return true
}

func GetAddressList(keys []string, Page int) (bool, *FofaHostInfo) {
	qBase64 := "ZmlkPSJ6MkVOWXZoUi82US9hZ0VYR0ZWbWRBPT0i" // 查询语句
	client := resty.New()
	fofainfo := &FofaHostInfo{}
	Pages := strconv.Itoa(Page)
	client.R().SetResult(&fofainfo).Get("https://fofa.info/api/v1/search/all?email=" + keys[0] + "&key=" + keys[1] + "&qbase64=" + qBase64 + "&page=" + Pages + "&full=true")
	if fofainfo.Error {
		return false, fofainfo
	}
	return true, fofainfo
}
