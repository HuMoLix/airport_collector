package utils

import (
	crand "crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/HuMoLix/airport_collector/common"
	"github.com/go-resty/resty/v2"
	"gopkg.in/yaml.v2"
)

func IsFileExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func IsFolderExists(folderpath string) bool {
	_, err := os.Stat(folderpath)
	return !os.IsNotExist(err)
}

func InputString(prefix string, Default string) string {
	fmt.Print(prefix)
	var result string
	fmt.Scanln(&result)
	if result == "" {
		return Default
	}
	return result
}

func ParseBool(str string) (bool, error) {
	switch str {
	case "1", "t", "T", "true", "TRUE", "True", "y":
		return true, nil
	case "0", "f", "F", "false", "FALSE", "False", "n":
		return false, nil
	}
	return false, &strconv.NumError{}
}

func String2Bool(Bool bool) string {
	if Bool {
		return "y"
	} else {
		return "n"
	}
}

func BoolChoose(prefix string, Default bool) bool {
	fmt.Print(prefix + ",默认选项:" + String2Bool(Default) + ",[y/n] ")
	var result string
	fmt.Scanln(&result)
	res, err := ParseBool(result)
	if err != nil {
		return Default
	}
	return res
}

func GetRandomBase64String(length int) string {
	b := make([]byte, length)
	_, err := crand.Read(b)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(b)
}

func GetRandomNumber() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprint(r.Uint32())
}

func DecodeFofa(data []string) []string {
	email := strings.Split(data[0], "@")
	key := []byte(data[1])
	for i := 0; i < len(data[1]); i++ {
		key[i] = data[1][len(data[1])-1-i]
	}
	email1 := strings.Replace(email[1], "google.com", "qq.com", -1)
	email0 := []byte(email[0])
	for i := 0; i < len(email[0]); i++ {
		email0[i] = email0[i] + 2
	}
	return []string{string(email0) + "@" + email1, string(key)}
}

func IsInit(Conf *common.Conf) bool {
	if err, _ := IsFileExist("config.yaml"); !err {
		fmt.Println("[-]错误，未识别config.yaml，现进行初始化操作。请按照指引完成初始化")
		Init(Conf)
	}
	return true
}

func Init(Conf *common.Conf) bool {
	if err, _ := IsFileExist(common.ConfigYaml); !err {
		conf := &common.Conf{
			// Version: config.Version,
			Settings: common.Settings{
				IsBlackList: BoolChoose("[*]是否记录无法登录的IP地址", true),
			},
			Fofa: common.FofaConf{
				Email:  InputString("[*]请输入Fofa邮箱: ", ""),
				Apikey: InputString("[*]请输入Fofa账号ApiKey: ", ""),
				Query: common.Query{
					Total_Size:      0,
					Last_Query_Page: 1,
				},
			},
		}
		data, _ := yaml.Marshal(conf)
		err := os.WriteFile(common.ConfigYaml, data, 0777)
		if err != nil {
			fmt.Printf("[-]配置文件创建失败 %v", err)
			os.Exit(0)
		}
		fmt.Println("[+]配置文件创建成功")
	}
	if !IsFolderExists(common.GeoIP) {
		err := os.Mkdir(common.GeoIP, 0777)
		if err != nil {
			fmt.Println("[-]数据库文件夹创建失败")
			os.Exit(0)
		}
		fmt.Println("[*]为了更好的体验，请前往 https://www.maxmind.com/ 下载 City Country ASN 数据库，并放置于" + common.GeoIP + "文件夹下。")
	}
	return true
}

func IsGeoIP(Options *common.ApCOptions) bool {
	if err, _ := IsFileExist(common.GeoIP + "/" + common.GeoIPPrefix + "ASN.mmdb"); !err {
		fmt.Println("[-]未检测到GeoIP数据库，将使用在线IP地址查询。")
		Options.GeoIP = false
	} else {
		Options.GeoIP = true
	}
	return true
}

func LoadConfig(Conf *common.Conf) bool {
	yamlFile, err := os.ReadFile("config.yaml")
	if err != nil {
		fmt.Println("[-]错误，无法找到config.yaml文件。")
		os.Exit(0)
	}
	if err := yaml.Unmarshal(yamlFile, Conf); err != nil {
		fmt.Println("[-]错误，配置文件加载失败")
		os.Exit(0)
	}
	return true
}

func UpdateConfig(Conf *common.Conf) bool {
	data, _ := yaml.Marshal(Conf)
	err := os.WriteFile("config.yaml", data, 0777)
	if err != nil {
		fmt.Println("[-]更新配置文件失败")
		os.Exit(0)
	}
	return true
}

func FormatData(ip string, port string, geoip bool, GeoIP *common.Geoip) (bool, []map[string]interface{}) {
	var Location Location
	IP := fmt.Sprintf("%v:%v", ip, port)
	err, cookie := getCookie(IP)
	if !err {
		return false, nil
	}
	if checkXray(IP, cookie) {
		return true, nil
	}
	data := getBound(IP, cookie)
	if geoip {
		Location = *GetIpLocationGeoIP(ip, GeoIP)
	} else {
		Location = *GetIpLocationOnline(ip)
	}
	var results []map[string]interface{}
	for _, host := range data {
		// host["name"] = Location + utils.GetRandomBase64String(5)
		host["name"] = fmt.Sprintf("%v|%v|%v|%v|%v", host["type"], Location.Country+Location.City, ip, host["port"], Location.ASN)
		host["server"] = ip
		results = append(results, host)
	}
	return true, results
}

func GenXml(proxies [][]map[string]interface{}) bool {
	xml := &common.XML{
		Port:                7890,
		Socks_Port:          7891,
		Redir_Port:          7892,
		Allow_LAN:           false,
		Mode:                "rule",
		Log_Level:           "silent",
		External_Controller: "0.0.0.0:9090",
		Secret:              "",
		CFW_Latency_Timeout: 5000,
		CFW_Bypass: []string{
			"qq.com",
			"music.163.com",
			"localhost",
			"127.*",
			"10.*",
			"172.16.*",
			"172.17.*",
			"172.18.*",
			"172.19.*",
			"172.20.*",
			"172.21.*",
			"172.22.*",
			"172.23.*",
			"172.24.*",
			"172.25.*",
			"172.26.*",
			"172.27.*",
			"172.28.*",
			"172.29.*",
			"172.30.*",
			"172.31.*",
			"192.168.*",
			"<local>",
		},
		Rule: []string{
			"DOMAIN-SUFFIX,edgedatg.com,国外视频网站",
			"DOMAIN-SUFFIX,go.com,国外视频网站",
			"DOMAIN-SUFFIX,abema.io,国外视频网站",
			"DOMAIN,linear-abematv.akamaized.net,国外视频网站",
			"DOMAIN-SUFFIX,abema.tv,国外视频网站",
			"DOMAIN-SUFFIX,akamaized.net,国外视频网站",
			"DOMAIN-SUFFIX,ameba.jp,国外视频网站",
			"DOMAIN-SUFFIX,hayabusa.io,国外视频网站",
			"DOMAIN-SUFFIX,aiv-cdn.net,国外视频网站",
			"DOMAIN-SUFFIX,amazonaws.com,国外视频网站",
			"DOMAIN-SUFFIX,amazonvideo.com,国外视频网站",
			"DOMAIN-SUFFIX,llnwd.net,国外视频网站",
		},
	}
	for i := 0; i < len(proxies); i++ {
		xml.Proxies = append(xml.Proxies, proxies[i]...)
	}
	var proxiesNames []string
	for i := 0; i < len(proxies); i++ {
		for j := 0; j < len(proxies[i]); j++ {
			proxiesNames = append(proxiesNames, proxies[i][j]["name"].(string))
		}
	}
	xml.Proxy_Groups = append(xml.Proxy_Groups, common.Proxy_Groups{
		Name:    "代理池",
		Type:    "select",
		Proxies: proxiesNames,
	})
	xml.Proxy_Groups = append(xml.Proxy_Groups, common.Proxy_Groups{
		Name:    "国内视频网站",
		Type:    "select",
		Proxies: []string{"DIRECT", "代理池"},
	})
	xml.Proxy_Groups = append(xml.Proxy_Groups, common.Proxy_Groups{
		Name:    "国外视频网站",
		Type:    "select",
		Proxies: []string{"代理池"},
	})
	Xml, _ := yaml.Marshal(xml)
	if !IsFolderExists(common.Results) {
		err := os.Mkdir(common.Results, 0777)
		if err != nil {
			return false
		}
	}
	err := os.WriteFile(fmt.Sprintf("%v/%v.yaml", common.Results, time.Now().Format("2006-01-02_15-04-05")), Xml, 0777)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

// 获取IP的物理地址

func GetIpLocationOnline(ip string) *Location {
	var asn string
	client := resty.New()
	resp, _ := client.R().
		SetHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.87 Safari/537.36").
		Get("https://m.ip138.com/iplookup.php?ip=" + ip)
	if resp.StatusCode() != 200 {
		return &Location{}
	}
	re := regexp.MustCompile(`<tbody>(.*?)</tbody>`).FindStringSubmatch(strings.Replace(string(resp.Body()), "\n", "", -1))[0]
	location := regexp.MustCompile(`<span>(.*?)</span>`).FindStringSubmatch(re)[1]
	if regexp.MustCompile(`<tr>(.*?)</tr>`).FindStringSubmatch(re) != nil {
		asn = regexp.MustCompile(`<td>(.*?)</td>`).FindStringSubmatch(regexp.MustCompile(`<tr>(.*?)</tr>`).FindStringSubmatch(re)[0])[1]
	}
	Location := &Location{
		Country: location,
		City:    "",
		ASN:     asn,
	}
	return Location
}

type Location struct {
	Country string // 国家
	City    string // 城市
	ASN     string // 运营商
}

func GetIpLocationGeoIP(ip string, GeoIP *common.Geoip) *Location {
	IP := net.ParseIP(ip)
	city, _ := GeoIP.City.City(IP)
	country, _ := GeoIP.Country.Country(IP)
	asn, _ := GeoIP.ASN.ASN(IP)
	Location := &Location{
		Country: country.Country.Names["zh-CN"],
		City:    city.City.Names["zh-CH"],
		ASN:     asn.AutonomousSystemOrganization,
	}
	return Location
}
