// Author: HuMoLix
package main

import (
	"os"

	"github.com/HuMoLix/airport_collector/common"
	"github.com/HuMoLix/airport_collector/common/utils"
	"github.com/HuMoLix/airport_collector/platform/fofa"
	"github.com/HuMoLix/airport_collector/runner"
)

// 主程序入口

func main() {
	var Config common.Conf
	var Options common.ApCOptions
	common.Flag(&Options)           // 解析命令行参数
	utils.IsInit(&Config)           // 检查是否初始化
	utils.IsGeoIP(&Options)         // 检查是否使用GeoIP库
	utils.LoadConfig(&Config)       // 读取 config.yaml 配置
	ok := fofa.ConnectFofa(&Config) // 检查与Fofa平台的连接状态
	if !ok {                        // 判断运行方式 license or fofakey
		os.Exit(0) // 退出程序
	} else {
		runner.Runner([]string{Config.Fofa.Email, Config.Fofa.Apikey}, &Config, &Options)
	}
}
