package runner

import (
	"fmt"
	"log"
	"math"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/HuMoLix/airport_collector/common"
	"github.com/HuMoLix/airport_collector/common/utils"
	"github.com/HuMoLix/airport_collector/platform/fofa"
	"github.com/oschwald/geoip2-golang"
	"github.com/schollz/progressbar/v3"
)

var datas [][]map[string]interface{}

func Runner(keys []string, Conf *common.Conf, Options *common.ApCOptions) {
	var GeoIP common.Geoip
	Init(keys, Conf)
	SignalHandler()
	if Options.GeoIP {
		GeoIP.City, _ = geoip2.Open(common.GeoIP + "/" + common.GeoIPPrefix + "City.mmdb")
		GeoIP.Country, _ = geoip2.Open(common.GeoIP + "/" + common.GeoIPPrefix + "Country.mmdb")
		GeoIP.ASN, _ = geoip2.Open(common.GeoIP + "/" + common.GeoIPPrefix + "ASN.mmdb")
	}
	bar := progressbar.NewOptions(Options.QueryPageCount*100, progressbar.OptionShowCount())
	ch := make(chan []map[string]interface{})
	wg := &sync.WaitGroup{}
	for page := 1; page < Options.QueryPageCount+1; page++ {
		_, results := fofa.GetAddressList(keys, page)
		for _, host := range results.Results {
			wg.Add(1)
			go func(host []string, group *sync.WaitGroup) {
				defer group.Done()
				err, data := utils.FormatData(host[1], host[2], Options.GeoIP, &GeoIP)
				if err && data != nil {
					ch <- data
				}
				bar.Add(1)
			}(host, wg)
		}
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	for data := range ch {
		datas = append(datas, data)
	}
	if !utils.GenXml(datas) {
		log.Fatal("[-]导出yaml文件失败。请提交issue。")
	}
	fmt.Println("\n[+]yaml文件导出成功")
}

func Init(keys []string, Conf *common.Conf) {
	Query := &Conf.Fofa.Query
	err, results := fofa.GetAddressList(keys, Query.Last_Query_Page)
	if !err {
		fmt.Println("[-]错误，查询失败，请检查与Fofa的网络连接或参数。", results.ErrMsg)
		os.Exit(0)
	}
	requiredToQueryPage := int(math.Round(float64(results.Size)/100) - float64(Query.Last_Query_Page) + 1)
	fmt.Printf("[*]根据记录，您上次查询总得结果%d条，查询至%d页，并在您查询后Fofa更新了%d条数据，您还剩%d次需要查询。\n", Query.Total_Size, Query.Last_Query_Page, (results.Size - Query.Total_Size), requiredToQueryPage)
	Query.Total_Size = results.Size
	Query.Last_Query_Page = results.Page
	utils.UpdateConfig(Conf)
}

func SignalHandler() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\n[*]检测到程序中断，正在导出yaml文件。")
		utils.GenXml(datas)
		os.Exit(0)
	}()
}
