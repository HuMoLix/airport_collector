package utils

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/HuMoLix/airport_collector/common"
	"github.com/go-resty/resty/v2"
)

type XUI struct {
	Status bool   `json:"success"` // 判断返回状态码 ture false
	Msg    string `json:"msg"`     // 返回信息
	Obj    []struct {
		Id             int    `json:"id"`             // Id
		Enable         bool   `json:"enable"`         // 是否开启
		Port           int    `json:"port"`           // 端口信息
		Protocol       string `json:"protocol"`       // 协议信息
		Settings       string `json:"settings"`       // 设置
		StreamSettings string `json:"streamSettings"` // 流设置
	} `json:"obj"` // 返回功能点信息
}

type Xray struct {
	Status bool   `json:"success"` // 判断返回状态码 ture false
	Msg    string `json:"msg"`     // 返回信息
	Obj    struct {
		Xray struct {
			State string `json:"state"` // 判断是否正常运行
		} `json:"xray"`
	} `json:"obj"`
}

// 尝试登录x-ui并获取Cookie

func getCookie(ip string) (bool, string) {
	client := resty.New().SetTimeout(common.Timeout)
	result := &XUI{}
	reqBody := map[string]interface{}{
		"username": "admin",
		"password": "admin",
	}
	resp, _ := client.R().SetResult(result).SetBody(reqBody).Post("http://" + ip + "/login")
	if !result.Status {
		return false, ""
	}
	return true, resp.Cookies()[0].Value
}

// 判断功能是否正常运行

func checkXray(ip string, cookie string) bool {
	client := resty.New().SetTimeout(common.Timeout)
	Model := &Xray{}
	client.R().SetResult(Model).SetHeader("Cookie", "session="+cookie).Post("http://" + ip + "/server/status")
	return Model.Obj.Xray.State == "error"
}

// 使用 Cookie 获取 vpn 设置

func getBound(ip string, cookie string) []map[string]interface{} {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err, ip)
		}
	}()
	client := resty.New().SetTimeout(common.Timeout)
	bound := &XUI{}
	client.R().SetResult(bound).SetHeader("Cookie", "session="+cookie).Post("http://" + ip + "/xui/inbound/list")
	var results []map[string]interface{}
	for i := 0; i < len(bound.Obj); i++ {
		Protocol := bound.Obj[i].Protocol
		Settings := strings.Replace(strings.Replace(bound.Obj[i].Settings, "\n", "", -1), " ", "", -1)
		StreamSettings := strings.Replace(strings.Replace(bound.Obj[i].StreamSettings, "\n", "", -1), " ", "", -1)
		var result map[string]interface{}
		switch Protocol {
		case "socks":
			result = Socks5(Settings)
			bound.Obj[i].Protocol = "socks5"
		case "shadowsocks":
			result = ShadowSocks(Settings)
			bound.Obj[i].Protocol = "ss"
		case "vmess", "vless":
			result = Vlmess(Settings, StreamSettings, Protocol, ip)
		case "http":
			result = Http(Settings)
		case "dokodemo-door":
			continue
		case "trojan":
			// result = Trojan(Settings)
			continue
		}
		result["port"] = strconv.Itoa(bound.Obj[i].Port)
		result["type"] = bound.Obj[i].Protocol
		results = append(results, result)
	}
	return results
}

func Socks5(Settings string) map[string]interface{} {
	settings := make(map[string]interface{})
	json.Unmarshal([]byte(Settings), &settings)
	if settings["auth"].(string) == "noauth" {
		return map[string]interface{}{}
	}
	result := map[string]interface{}{
		"username": settings["accounts"].([]interface{})[0].(map[string]interface{})["user"].(string),
		"password": settings["accounts"].([]interface{})[0].(map[string]interface{})["pass"].(string),
	}
	return result
}

func ShadowSocks(Settings string) map[string]interface{} {
	settings := make(map[string]string)
	json.Unmarshal([]byte(Settings), &settings)
	result := map[string]interface{}{
		"cipher":   settings["method"],
		"password": settings["password"],
		"network":  settings["network"],
	}
	return result
}

func Vlmess(Settings string, StreamSettings string, Protocol string, ip string) map[string]interface{} {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err, ip)
			fmt.Println(Settings, StreamSettings)
		}
	}()
	settings := make(map[string]interface{})
	streamsettings := make(map[string]interface{})
	json.Unmarshal([]byte(Settings), &settings)
	json.Unmarshal([]byte(StreamSettings), &streamsettings)
	result := map[string]interface{}{
		"uuid":    settings["clients"].([]interface{})[0].(map[string]interface{})["id"].(string),
		"tls":     isTLSenable(streamsettings["security"].(string)),
		"cipher":  "auto",
		"network": streamsettings["network"].(string),
	}
	switch Protocol {
	case "vmess":
		result["alterId"] = settings["clients"].([]interface{})[0].(map[string]interface{})["alterId"].(float64)
	}
	if isTLSenable(streamsettings["security"].(string)) {
		result["servername"] = streamsettings["tlsSettings"].(map[string]interface{})["serverName"]
		result["skip-cert-verify"] = true
	}
	switch result["network"] {
	case "tcp":
		result["network"] = streamsettings["tcpSettings"].(map[string]interface{})["header"].(map[string]interface{})["type"]
		if result["network"] == "none" {
			result["network"] = streamsettings["network"].(string)
			break
		}
		result[fmt.Sprintf("%v-opts", result["network"])] = map[string]interface{}{
			"method": streamsettings["tcpSettings"].(map[string]interface{})["header"].(map[string]interface{})["request"].(map[string]interface{})["method"].(string),
			"path":   streamsettings["tcpSettings"].(map[string]interface{})["header"].(map[string]interface{})["request"].(map[string]interface{})["path"].([]interface{}),
		}
	case "ws":
		result["ws-opts"] = map[string]string{
			"path": streamsettings["wsSettings"].(map[string]interface{})["path"].(string),
		}
		if _, ok := streamsettings["wsSettings"].(map[string]interface{})["headers"].(map[string]interface{})["Host"].(string); ok {
			result["ws-opts"].(map[string]string)["host"] = streamsettings["wsSettings"].(map[string]interface{})["headers"].(map[string]interface{})["Host"].(string)
		}
	}
	return result
}

func Http(Settings string) map[string]interface{} {
	settings := make(map[string]interface{})
	json.Unmarshal([]byte(Settings), &settings)
	result := map[string]interface{}{
		"username": settings["accounts"].([]interface{})[0].(map[string]interface{})["user"].(string),
		"password": settings["accounts"].([]interface{})[0].(map[string]interface{})["pass"].(string),
	}
	return result
}

// func Trojan(Settings string) map[string]interface{} {

// }

func isTLSenable(Security string) bool {
	if Security == "tls" {
		return true
	} else {
		return false
	}
}
