package update

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

func ValidVersion() {
	github_url := "http://github.com/HuMoLix/airport_collector/update/update_info.json"
	client := resty.New()
	resp, _ := client.R().
		SetHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.87 Safari/537.36").
		Get(github_url)
	fmt.Println(resp)
}
