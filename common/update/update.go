package update

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/HuMoLix/airport_collector/common"
	"github.com/go-resty/resty/v2"
)

func ValidVersion() {
	var version common.Version
	currentVersion := common.GetCurrentVersion()
	github_url := "https://raw.githubusercontent.com/HuMoLix/airport_collector/main/common/update/update_info.json"
	client := resty.New().SetTimeout(common.Timeout)
	resp, _ := client.R().
		SetHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.87 Safari/537.36").
		Get(github_url)
	if resp.StatusCode() != 200 {
		fmt.Println("[-]与版本更新服务器连接超时，请检查网络配置后重试。")
		os.Exit(0)
	}
	json.Unmarshal(resp.Body(), &version)
	if currentVersion != version.Version {
		fmt.Println("[*]有新版本更新，请前往 " + common.Github + " 下载新版本。")
		os.Exit(0)
	}
}
